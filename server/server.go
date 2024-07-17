package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	// "net/http"
	"net"
	"strings"

	pb "github.com/HongQuan1812/VNExpressProject/VNExpress_selector"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
)

type vNExpressSelectorServer struct {
	pb.UnimplementedVNExpressSelectorServer
}

func (*vNExpressSelectorServer) ConnectDatabase() (*sql.DB, error) {
	dsn := "HongQuan:18122003@tcp(localhost:3306)/VnExpressDatabase"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Check if the connection is alive
	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database!")
	return db, nil
}

func Specify_Relative_Duration(comparisor string, time1 interface{}, time2 interface{}) string {
	var relative_duration string

	time1_ := fmt.Sprintf("%#v", time1)
	time2_ := fmt.Sprintf("%#v", time2)
	fmt.Println((time1_))

	if comparisor == "BETWEEN" {
		relative_duration = fmt.Sprintf("BETWEEN %v AND %v", time1_, time2_)
	} else {
		relative_duration = fmt.Sprintf("%s %v", comparisor, time1_)
	}
	return relative_duration
}

func Check_Valid(attribute sql.NullString) string {
	if attribute.Valid {
		return attribute.String
	}
	return ""
}

func (s *vNExpressSelectorServer) SelectNews(range_news *pb.Range, stream pb.VNExpressSelector_SelectNewsServer) error {
	db, err := s.ConnectDatabase()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Build the query
	contentQuery := `SELECT * FROM news, news_authors, authors WHERE 
						news_authors.id_news = news.id AND 
						news_authors.id_author = authors.id AND 
						1=1`

	var conditions []string
	var args []interface{}

	if len(range_news.MainCategories) > 0 {
		placeholders := make([]string, len(range_news.MainCategories))
		for i := range placeholders {
			placeholders[i] = "?"
			args = append(args, range_news.MainCategories[i])
		}

		condition := fmt.Sprintf("main_category IN (%s)", strings.Join(placeholders, ", "))
		conditions = append(conditions, condition)
	}

	if len(range_news.SubCategories) > 0 {
		placeholders := make([]string, len(range_news.SubCategories))
		for i := range placeholders {
			placeholders[i] = "?"
			args = append(args, range_news.SubCategories[i])
		}

		condition := fmt.Sprintf("sub_category IN (%s)", strings.Join(placeholders, ", "))
		conditions = append(conditions, condition)
	}

	if len(range_news.Author) > 0 {
		placeholders := make([]string, len(range_news.Author))
		for i := range placeholders {
			placeholders[i] = "?"
			args = append(args, range_news.Author[i])
		}
		condition := fmt.Sprintf("name IN (%s)", strings.Join(placeholders, ", "))
		conditions = append(conditions, condition)

	}

	if len(range_news.Day) > 0 {

		if range_news.WholeDay {

			relative_duration := Specify_Relative_Duration(range_news.DayComparisor[0], range_news.Day[0], range_news.Day[1])

			condition := strings.Join([]string{"day", relative_duration}, " ")
			conditions = append(conditions, condition)

		} else {
			components1 := strings.Split(range_news.Day[0], "-")
			components2 := strings.Split(range_news.Day[1], "-")
			component := []string{"DATE", "MONTH", "YEAR"}

			for i := range components1 {
				if components1[i] != "any" {
					value1, _ := strconv.Atoi(components1[i])
					value2, _ := strconv.Atoi(components2[i])
					relative_duration := Specify_Relative_Duration(range_news.DayComparisor[i], value1, value2)
					temp := fmt.Sprintf("%s(day)", component[i])
					condition := strings.Join([]string{temp, relative_duration}, " ")
					conditions = append(conditions, condition)
				}
			}
		}
	}

	if len(range_news.Time) > 0 {
		relative_duration := Specify_Relative_Duration(range_news.TimeComparisor, range_news.Time[0], range_news.Time[1])

		condition := strings.Join([]string{"time", relative_duration}, " ")
		conditions = append(conditions, condition)
	}

	if len(conditions) > 0 {
		contentQuery = fmt.Sprintf("%s AND %s", contentQuery, strings.Join(conditions, " AND "))
	}

	if range_news.Limit != "" {
		limit := fmt.Sprintf("LIMIT %s", range_news.Limit)
		contentQuery = strings.Join([]string{contentQuery, limit}, " ")
	}

	fmt.Print("--------------------\n")
	fmt.Printf("contentQuery: %s \n", contentQuery)
	fmt.Printf("args: %v \n", args)
	fmt.Print("--------------------\n")

	rows, err := db.Query(contentQuery, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	// Define variables to store the data
	var id_news, id_author int
	var url, mainCategory, subCategory, title, timeZone, description, newsContent, day, time, author_name sql.NullString

	for rows.Next() {
		// Your logic to handle each row
		err := rows.Scan(&id_news, &url, &mainCategory, &subCategory, &title, &day, &time, &timeZone, &description, &newsContent, &id_news, &id_author, &id_author, &author_name)
		if err != nil {
			log.Fatal(err)
		}

		news := pb.News{
			Url:           Check_Valid(url),
			MainCategory:  Check_Valid(mainCategory),
			SubCategory:   Check_Valid(subCategory),
			Title:         Check_Valid(title),
			Day:           Check_Valid(day),
			Time:          Check_Valid(time),
			TimeZone:      Check_Valid(timeZone),
			Description:   Check_Valid(description),
			NewsContent:   Check_Valid(newsContent),
			RelatingImage: []string{},
			Author:        Check_Valid(author_name),
		}

		imagesQuery := "SELECT url FROM images_of_news WHERE id_news = ?"
		imgRows, err := db.Query(imagesQuery, id_news)
		if err != nil {
			return fmt.Errorf("failed to execute query: %v", err)
		}
		defer imgRows.Close()

		var imageURL string
		for imgRows.Next() {
			err := imgRows.Scan(&imageURL)
			if err != nil {
				log.Fatal(err)
			}
			news.RelatingImage = append(news.RelatingImage, imageURL)
		}

		if err := imgRows.Err(); err != nil {
			log.Fatal(err)
		}

		if err := stream.Send(&news); err != nil {
			return err
		}
	}
	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *vNExpressSelectorServer) SelectPodcast(range_podcast *pb.Range, stream pb.VNExpressSelector_SelectPodcastServer) error {
	db, err := s.ConnectDatabase()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Build the query
	contentQuery := `SELECT * FROM podcasts, podcasts_authors, authors WHERE 
						podcasts_authors.id_podcast = podcasts.id AND 
						podcasts_authors.id_author = authors.id AND
						1=1`

	var conditions []string
	var args []interface{}

	if len(range_podcast.MainCategories) > 0 {
		placeholders := make([]string, len(range_podcast.MainCategories))
		for i := range placeholders {
			placeholders[i] = "?"
		}

		condition := fmt.Sprintf("main_category IN (%s)", strings.Join(placeholders, ", "))
		conditions = append(conditions, condition)
		for _, main_category := range range_podcast.MainCategories {
			args = append(args, main_category)
		}
	}

	if len(range_podcast.SubCategories) > 0 {
		placeholders := make([]string, len(range_podcast.SubCategories))
		for i := range placeholders {
			placeholders[i] = "?"
		}

		condition := fmt.Sprintf("sub_category IN (%s)", strings.Join(placeholders, ", "))
		conditions = append(conditions, condition)
		for _, sub_category := range range_podcast.MainCategories {
			args = append(args, sub_category)
		}
	}

	if len(range_podcast.Author) > 0 {
		placeholders := make([]string, len(range_podcast.Author))
		for i := range placeholders {
			placeholders[i] = "?"
		}
		condition := fmt.Sprintf("name IN (%s)", strings.Join(placeholders, ", "))
		conditions = append(conditions, condition)
		for _, author_name := range range_podcast.Author {
			args = append(args, author_name)
		}
	}

	if len(range_podcast.Day) > 0 {

		if range_podcast.WholeDay {

			relative_duration := Specify_Relative_Duration(range_podcast.DayComparisor[0], range_podcast.Day[0], range_podcast.Day[1])

			condition := strings.Join([]string{"day", relative_duration}, " ")
			conditions = append(conditions, condition)

		} else {
			components1 := strings.Split(range_podcast.Day[0], "-")
			components2 := strings.Split(range_podcast.Day[1], "-")
			component := []string{"DATE", "MONTH", "YEAR"}

			for i := range components1 {
				if components1[i] != "any" {
					relative_duration := Specify_Relative_Duration(range_podcast.DayComparisor[i], components1[i], components2[i])
					temp := fmt.Sprintf("%s(day)", component[i])
					condition := strings.Join([]string{temp, relative_duration}, " ")
					conditions = append(conditions, condition)
				}
			}
		}
	}

	if len(range_podcast.Time) > 0 {
		relative_duration := Specify_Relative_Duration(range_podcast.TimeComparisor, range_podcast.Time[0], range_podcast.Time[1])

		condition := strings.Join([]string{"time", relative_duration}, " ")
		conditions = append(conditions, condition)
	}

	if len(conditions) > 0 {
		contentQuery = fmt.Sprintf("%s AND %s", contentQuery, strings.Join(conditions, " AND "))
	}

	if range_podcast.Limit != "" {
		limit := fmt.Sprintf("LIMIT %s", range_podcast.Limit)
		contentQuery = strings.Join([]string{contentQuery, limit}, " ")
	}

	fmt.Print("--------------------\n")
	fmt.Printf("contentQuery: %s \n", contentQuery)
	fmt.Printf("args: %v \n", args)
	fmt.Print("--------------------\n")

	rows, err := db.Query(contentQuery, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	// Define variables to store the data
	var id_podcast, id_author int
	var url, mainCategory, subCategory, title, day, time, timeZone, description, relatingPodcast, author string

	// Handle the results
	for rows.Next() {
		err := rows.Scan(&id_podcast, &url, &mainCategory, &subCategory, &title, &day, &time, &timeZone, &description, &relatingPodcast, &id_podcast, &id_author, &id_author, &author)
		if err != nil {
			log.Fatal(err)
		}

		podcast := pb.Podcast{
			Url:             url,
			MainCategory:    mainCategory,
			SubCategory:     subCategory,
			Title:           title,
			Day:             day,
			Time:            time,
			TimeZone:        timeZone,
			Description:     description,
			RelatingPodcast: relatingPodcast,
			Author:          author,
		}

		if err := stream.Send(&podcast); err != nil {
			return err
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &vNExpressSelectorServer{}
	pb.RegisterVNExpressSelectorServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
