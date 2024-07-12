package main

import (
	"context"
	"flag"
	"io"
	"log"
	"strings"
	"time"

	pb "github.com/HongQuan1812/VNExpressProject/VNExpress_selector"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost:8089", "The server address in the format of host:port")
	types      = flag.String("types", "news", "Do you want to select news or podcasts?")

	mainCategories = flag.String("mainCategories", "", "Comma-separated list of main categories, Example: \"Thể thao, Thời sự\" ")
	subCategories  = flag.String("subCategories", "", "Comma-separated list of sub categories, Example: \"Chính trị, Dân sinh\" ")
	author         = flag.String("author", "", "Comma-separated list of authors, Example: \"Hồng Quân, Nam Án\" ")
	wholeDay       = flag.Bool("wholeDay", true, "Whole day flag")
	dayComparisor  = flag.String("dayComparisor", "", "Comma-separated list of day comparisors, Example: if wholeDay is true, \"=\", else \"=, >, any\" ")
	release_day    = flag.String("release_day", "", "Comma-separated list of days, Example: \"yyyy-mm-dd, any-any-any\" ")
	timeComparisor = flag.String("timeComparisor", "", "Time comparisor")
	release_time   = flag.String("release_time", "", "Comma-separated list of times, Example: \"hh:mm, hh:mm\" ")
	limit          = flag.String("limit", "5", "Limit")
)

func check_empty(attribute string) []string {
	if attribute != "" {
		return strings.Split(attribute, ",")
	}
	return []string{}
}

func main() {
	// ------------- Set up a connection to the server and Create a client. ----------------
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewVNExpressSelectorClient(conn)

	// ---------------------- Create a new CreateRequest ---------------------

	req := &pb.Range{
		MainCategories: check_empty(*mainCategories),
		SubCategories:  check_empty(*subCategories),
		Author:         check_empty(*author),
		WholeDay:       *wholeDay,
		DayComparisor:  check_empty(*dayComparisor),
		Day:            check_empty(*release_day),
		TimeComparisor: *timeComparisor,
		Time:           check_empty(*release_time),
		Limit:          *limit,
	}

	// -------------- Contact the server and print out its response. -------------
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if *types == "news" {
		stream, err := client.SelectNews(ctx, req)
		if err != nil {
			log.Fatalf("could not create: %v", err)
		}

		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving response: %v", err)
			}
			log.Printf(
				`News:\n
					URL=%q,\n
					MainCategory=%q,\n
					SubCategory=%q,\n
					Title=%q,\n
					Day=%q,\n
					Time=%q,\n
					TimeZone=%q,\n
					Description=%q,\n
					NewsContent=%q,\n
					RelatingImage=%v,\n
					Author=%q,\n
			 	`,
				res.Url, res.MainCategory, res.SubCategory, res.Title, res.Day, res.Time, res.TimeZone,
				res.Description, res.NewsContent, res.RelatingImage, res.Author,
			)
		}

	} else {
		stream, err := client.SelectPodcast(ctx, req)
		if err != nil {
			log.Fatalf("could not create: %v", err)
		}

		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while receiving response: %v", err)
			}
			log.Printf(
				`Podcast:\n 
					URL=%q,\n 
					MainCategory=%q,\n
					SubCategory=%q,\n
					Title=%q,\n
					Day=%q,\n
					Time=%q,\n
					TimeZone=%q,\n
					Description=%q,\n
					RelatingPodcast=%q,\n
					Author=%q,\n
				`,
				res.Url, res.MainCategory, res.SubCategory, res.Title, res.Day, res.Time, res.TimeZone,
				res.Description, res.RelatingPodcast, res.Author,
			)
		}
	}

}
