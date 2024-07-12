import scrapy
from VNExpressProject.items import NewsItem
from urllib.parse import urlencode
import re
import unicodedata
import json


API_KEY = '07635ea2-929a-43f2-875b-c1e43ff4b72d'

def get_proxy_url(url):
    payload = {
        'api_key': API_KEY, 
        'url': url, 
        # 'render_js': 'true', 
        # 'residential': 'true', 
        # 'country': 'us'
    }
    proxy_url = 'https://proxy.scrapeops.io/v1/?' + urlencode(payload)
    return proxy_url
    # return url


def remove_diacritics(text):
    # Normalize the text to its decomposed form
    normalized = unicodedata.normalize('NFD', text)
    # Filter out non-spacing marks (diacritics)
    without_diacritics = ''.join([c for c in normalized if unicodedata.category(c) != 'Mn'])
    return without_diacritics


class NewsspiderSpider(scrapy.Spider):
    name = 'newsspider'
    allowed_domains = ['vnexpress.net', 'proxy.scrapeops.io']
    start_urls = ['https://vnexpress.net']
    
    custom_settings = {
        'FEEDS': {
            'newsdata.json': {'format': 'json', 'overwrite': True},
        }
    }
    def __init__(self, num_page_to_crawl = None, mains = None, subs = None):
        super().__init__()
        
        self.num_page_to_crawl = int(num_page_to_crawl) if num_page_to_crawl is not None else None

        if mains is not None:
            try:
                mains = json.loads(mains)
            except json.decoder.JSONDecodeError:
                assert isinstance(mains, (list, tuple)),  "mains should be a list or tuple:  subs='[\"A\", \"B\"]', notice: not video and podcast"  
            mains = [remove_diacritics(category).lower() for category in mains]
        self.mains = mains
        
        if subs is not None:
            try:
                subs = json.loads(subs)
            except json.decoder.JSONDecodeError:
                assert isinstance(subs, (list, tuple)),  "subs should be a list or tuple:  subs='[\"A\", \"B\"]', notice: if mains, in defined mains else arbitrary"  
            subs = [remove_diacritics(category).lower() for category in subs]
        self.subs = subs
        
        
    def start_requests(self):
        yield scrapy.Request(url=get_proxy_url(self.start_urls[0]), callback=self.parse)
        
    def parse(self, response):
        main_categories = response.xpath("//section[@class='section wrap-main-nav']/nav[@class='main-nav']/ul[@class='parent']/li/a")[2:-1]
        main_categories.pop(3)     
        main_categories.pop(3)
            
        for main_category in main_categories:
            if self.mains is not None:
                main_category_title = main_category.css("a ::attr(title)").get()
                main_category_title = remove_diacritics(main_category_title).lower()
                
                print(main_category_title)
                print(self.mains)
                if main_category_title not in self.mains:
                    continue
            relative_url = main_category.css("a ::attr(href)").get()
            main_category_url = self.start_urls[0] + relative_url
            yield response.follow(get_proxy_url(main_category_url), callback=self.parse_main_category_page)
        
    def parse_main_category_page(self, response):
        sub_categories = response.xpath("//ul[@class='ul-nav-folder']/li/a")
        for sub_category in sub_categories:
            if self.subs is not None:
                sub_category_title = sub_category.css("a::attr(title)").get()
                sub_category_title = remove_diacritics(sub_category_title).lower()
                if sub_category_title not in self.subs:
                    continue
            relative_url = sub_category.css("a ::attr(href)").get()
            sub_category_url = self.start_urls[0] + relative_url
            yield response.follow(get_proxy_url(sub_category_url), callback=self.parse_sub_category_page)
    
    def parse_sub_category_page(self, response):
        def extract_page_number(url):
            match = re.search(r'p(\d+)$', url)
            if match:
                return int(match.group(1))
            return None
        
        news=response.xpath("//section[@class='section section_container mt15']//div[@id='automation_TV0']//article[contains(@class, 'item-news')]/div/a")
        for single_news in news:
            news_page_url = single_news.css("a::attr(href)").get()
            yield response.follow(get_proxy_url(news_page_url), callback=self.parse_news_page)
        
        next_page = response.css("a.btn-page.next-page::attr(href)").get()
        if next_page is not None:
            if self.num_page_to_crawl is not None:
                page_number = extract_page_number(next_page)
                if page_number is not None and page_number <= self.num_page_to_crawl:
                    next_page_url = self.start_urls[0] + next_page
                    yield response.follow(get_proxy_url(next_page_url), callback=self.parse_sub_category_page)
            else:
                next_page_url = self.start_urls[0] + next_page
                yield response.follow(get_proxy_url(next_page_url), callback=self.parse_sub_category_page)
    
    def parse_news_page(self, response):
        news_item = NewsItem()
        
        header_section = "//div[contains(@class,'header-content')]"
        body_section = "//article[contains(@class,'fck_detail')]"    
            
        # ---- url -----
        news_item['url'] = response.url
        
        # --- category ----
        news_item['main_category'] = response.xpath(header_section + "//li[1]/a/@title").get()
        news_item['sub_category'] = response.xpath(header_section + "//li[2]/a/@title").get()
        
        # ------ title ------
        news_item['title'] = response.xpath(header_section + "/following::h1/text()").get()

        # ----- author ------
        author_section = body_section + "//p[(@style='text-align:right;') or (contains(@class,'author')) or (@align='right')]"
        author = response.xpath(author_section + "//strong/text()").get()
        news_item['author'] = author
            
        # ----- release_date ------
        news_item['release_date'] = response.xpath(header_section + "/span[@class='date']/text()").get()
        
        # ----- description ------
        news_item['description'] = response.xpath(header_section + "/following::p[@class='description']/text()").get()
            
        # ----- content ------
        content_section = author_section + "/preceding::p[@class='Normal']"
        news_item['news_content'] = response.xpath(content_section + "/text()").getall()
            
        # ----- relating image ------
        relating_image_section = author_section + "/preceding::figure[@itemprop='associatedMedia image']//img[@itemprop='contentUrl']"
        news_item['relating_image'] = response.xpath(relating_image_section + "/@data-src").getall()
            
        yield news_item
        
        
