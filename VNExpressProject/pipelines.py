# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


# useful for handling different item types with a single interface
from itemadapter import ItemAdapter
from VNExpressProject.items import NewsItem, PodcastItem
import mysql.connector
from datetime import datetime
import re


class ProcessPipeline:
    def process_item(self, item, spider):
        adapter = ItemAdapter(item)

        if isinstance(item, NewsItem):
            self.process_news_item(adapter)
        elif isinstance(item, PodcastItem):
            self.process_podcast_item(adapter)

        return item
    
    def process_news_item(self, adapter):
        adapter['title'] = adapter['title'].strip()
        
        adapter['author'] = [re.sub(r'^\W+|\W+$', '', name.strip()) for name in adapter.get('author').split('-')]

        
        cleaned_datetime_string = adapter['release_date'].split(',', 1)[1].strip()
        date_time_components = [part.strip() for part in cleaned_datetime_string.replace('(', ',').replace(')', '').split(',')]
        date_object = datetime.strptime(date_time_components[0], '%d/%m/%Y').date()
        date_time_components[0] = date_object.strftime('%Y-%m-%d')
        adapter['release_date'] = date_time_components
        
        adapter['news_content'] = " ".join(adapter['news_content'])
        
    def process_podcast_item(self, adapter):
        adapter['author'] = [re.sub(r'^\W+|\W+$', '', name.strip()) for name in adapter.get('author').split('-')]
        
        cleaned_datetime_string = adapter['release_date'].split(',', 1)[1].strip()
        date_time_components = [part.strip() for part in cleaned_datetime_string.replace('(', ',').replace(')', '').split(',')]
        date_object = datetime.strptime(date_time_components[0], '%d/%m/%Y').date()
        date_time_components[0] = date_object.strftime('%Y-%m-%d')
        adapter['release_date'] = date_time_components
        

class SaveToMySQLPipeline:
    def __init__(self):
        self.conn = mysql.connector.connect(
            host = 'localhost',
            user = 'HongQuan',
            password = '18122003', #add your password here if you have one set 
            database = 'VnExpressDatabase'
        )

        # Create cursor, used to execute commands
        self.cur = self.conn.cursor()
        # Enable auto-commit
        # self.conn.autocommit = True

        self.cur.execute("""
        CREATE TABLE IF NOT EXISTS news (
            id INT NOT NULL AUTO_INCREMENT,
            url VARCHAR(550) NOT NULL,
            main_category VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
            sub_category VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
            title TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
            day DATE,
            time TIME,
            time_zone VARCHAR(50),
            description TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
            news_content TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
            
            PRIMARY KEY (id),
            CONSTRAINT unique_news_url UNIQUE (url)
        )
        """)
        
        self.cur.execute("""
        CREATE TABLE IF NOT EXISTS podcasts (
            id INT NOT NULL AUTO_INCREMENT,
            url VARCHAR(550) NOT NULL,
            main_category VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
            sub_category VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
            title TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
            day DATE,
            time TIME,
            time_zone VARCHAR(50),
            description TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
            relating_podcast VARCHAR(255),
            
            PRIMARY KEY (id),
            CONSTRAINT unique_podcast_url UNIQUE (url)
        )
        """)
        
        self.cur.execute("""
        CREATE TABLE IF NOT EXISTS images_of_news (
            id_news INT NOT NULL,
            url VARCHAR(550) NOT NULL,
            
            FOREIGN KEY (id_news) REFERENCES news(id),
            CONSTRAINT unique_pair_of_news_image UNIQUE (id_news, url)
        )
        """)
        
        self.cur.execute("""
        CREATE TABLE IF NOT EXISTS authors (
            id INT NOT NULL AUTO_INCREMENT,
            name VARCHAR(50) NOT NULL CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
            
            PRIMARY KEY (id),
            CONSTRAINT unique_author_name UNIQUE (name)
        )
        """)
        
        self.cur.execute("""
        CREATE TABLE IF NOT EXISTS news_authors (
            id_news INT NOT NULL,
            id_author INT NOT NULL,
            
            FOREIGN KEY (id_news) REFERENCES news(id),
            FOREIGN KEY (id_author) REFERENCES authors(id),
            CONSTRAINT unique_pair_of_news_author UNIQUE (id_news, id_author)
        )
        """)
        
        self.cur.execute("""
        CREATE TABLE IF NOT EXISTS podcasts_authors (
            id_podcast INT NOT NULL,
            id_author INT NOT NULL,
            
            FOREIGN KEY (id_podcast) REFERENCES podcasts(id),
            FOREIGN KEY (id_author) REFERENCES authors(id),
            CONSTRAINT unique_pair_of_podcast_author UNIQUE (id_podcast, id_author)
        )
        """)
        


    def process_item(self, item, spider):
		    ## Define insert statement
      
        if isinstance(item, NewsItem):
            self.process_news_item(item)
        elif isinstance(item, PodcastItem):
            self.process_podcast_item(item)

        return item
    
    def process_news_item (self, item):
        
        self.cur.execute("SELECT id FROM news WHERE url = %s", (item["url"], ))
        row = self.cur.fetchone()
        
        if row is None:
            insert_to_news = """ 
                    INSERT INTO news (url, main_category, sub_category, title, day,
                                    time, time_zone, description, news_content) 
                    values (%s, %s, %s, %s, %s, %s, %s, %s, %s)            
                """
            data_to_news = (item["url"], item["main_category"], item["sub_category"], item["title"], item["release_date"][0],
                            item["release_date"][1], item["release_date"][2], item["description"], item["news_content"])   
            
            self.cur.execute(insert_to_news, data_to_news)
            self.conn.commit()
        
            id_news = self.cur.lastrowid
        else:
            id_news = row[0]
        
        for image_url in item['relating_image']:
            insert_to_images_of_news = """ 
                    INSERT INTO images_of_news (id_news, url) 
                    values (%s, %s)
                """
            data_to_images_of_news = (id_news, image_url)   
            self.cur.execute(insert_to_images_of_news, data_to_images_of_news)
            self.conn.commit()

        
        for author_name in item['author']:
            
            self.cur.execute("SELECT id FROM authors WHERE name = %s", (author_name, ))
            row = self.cur.fetchone()
            
            if row is None:
                insert_to_authors = """ 
                        INSERT INTO authors (name) 
                        values (%s)
                    """
                data_to_authors = (author_name, )
                
                self.cur.execute(insert_to_authors, data_to_authors)
                self.conn.commit()
                
                id_author = self.cur.lastrowid
            else:
                id_author = row[0]
            
            insert_to_news_authors =  """ 
                    INSERT INTO news_authors (id_news, id_author) 
                    values (%s, %s)
                """
            data_to_news_authors = (id_news, id_author)   
            
            self.cur.execute(insert_to_news_authors, data_to_news_authors)
            self.conn.commit()

        
    
    def process_podcast_item (self, item):
        
        self.cur.execute("SELECT id FROM podcasts WHERE url = %s", (item["url"], ))
        row = self.cur.fetchone()
        
        if row is None:
            insert_to_podcasts = """ 
                    INSERT INTO podcasts (url, main_category, sub_category, title, day,
                                        time, time_zone, description, relating_podcast) 
                    values (%s, %s, %s, %s, %s, %s, %s, %s, %s)
                """
            data_to_podcasts = (item["url"], item["main_category"], item["sub_category"], item["title"], item["release_date"][0],
                                item["release_date"][1], item["release_date"][2], item["description"], item["relating_podcast"])   
            self.cur.execute(insert_to_podcasts, data_to_podcasts)
            self.conn.commit()
            
            id_podcast = self.cur.lastrowid
        else:
            id_podcast = row[0]


        
        for author_name in item['author']:
            
            self.cur.execute("SELECT id FROM authors WHERE name = %s", (author_name, ))
            row = self.cur.fetchone()
            
            if row is None:
                insert_to_authors = """ 
                        INSERT INTO authors (name) 
                        values (%s)
                    """
                data_to_authors = (author_name, )
                
                self.cur.execute(insert_to_authors, data_to_authors)
                self.conn.commit()
                
                id_author = self.cur.lastrowid
            else:
                id_author = row[0]
                
            insert_to_podcasts_authors =  """ 
                    INSERT INTO podcasts_authors (id_podcast, id_author) 
                    values (%s, %s)
                """
            data_to_podcasts_authors = (id_podcast, id_author)   
            self.cur.execute(insert_to_podcasts_authors, data_to_podcasts_authors)
            self.conn.commit()
            
                

    
    def close_spider(self, spider):
        ## Close cursor & connection to database 
        self.cur.close()
        self.conn.close()

    
    