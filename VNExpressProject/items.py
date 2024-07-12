# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html

import scrapy


class VnexpressprojectItem(scrapy.Item):
    # define the fields for your item here like:
    # name = scrapy.Field()
    pass

class NewsItem(scrapy.Item):
    url = scrapy.Field()
    main_category = scrapy.Field()
    sub_category = scrapy.Field()
    title = scrapy.Field()
    author = scrapy.Field()
    release_date = scrapy.Field()
    description = scrapy.Field()
    news_content = scrapy.Field()
    relating_image = scrapy.Field()
    
class PodcastItem(scrapy.Item):
    url = scrapy.Field()
    main_category = scrapy.Field()
    sub_category = scrapy.Field()
    title = scrapy.Field()
    author = scrapy.Field()
    release_date = scrapy.Field()
    description = scrapy.Field()
    relating_podcast = scrapy.Field()