import scrapy
from VNExpressProject.items import PodcastItem
from urllib.parse import urlencode
import unicodedata
import json
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from scrapy.http import HtmlResponse


API_KEY = '3cecba9a-d9c2-4aac-b2db-77d8c087cb29'

def get_proxy_url(url):
    # payload = {
    #     'api_key': API_KEY, 
    #     'url': url, 
    #     # 'render_js': 'true', 
    #     # 'residential': 'true', 
    #     # 'country': 'us'
    # }
    # proxy_url = 'https://proxy.scrapeops.io/v1/?' + urlencode(payload)
    # return proxy_url
    return url

def remove_diacritics(text):
    # Normalize the text to its decomposed form
    normalized = unicodedata.normalize('NFD', text)
    # Filter out non-spacing marks (diacritics)
    without_diacritics = ''.join([c for c in normalized if unicodedata.category(c) != 'Mn'])
    return without_diacritics


class PodcastspiderSpider(scrapy.Spider):
    name = "podcastspider"
    allowed_domains = ['vnexpress.net', 'proxy.scrapeops.io']
    start_urls = ["https://vnexpress.net/podcast"]
    
    custom_settings = {
        'FEEDS': {
            'podcastdata.json': {'format': 'json', 'overwrite': True},
        }
    }
    
    def __init__(self, showmore_times=0, subs=None, browser='Safari'):
        super().__init__()
        self.showmore_times = int(showmore_times) 
        
        if subs is not None:
            try:
                subs = json.loads(subs)
            except json.decoder.JSONDecodeError:
                assert isinstance(subs, (list, tuple)),  "subs should be a list or tuple:  subs='[\"A\", \"B\"]'"  
            subs = [remove_diacritics(category).lower() for category in subs]
        self.subs = subs
        
        if self.showmore_times != 0:
            if browser == 'Safari':
                self.driver = webdriver.Safari()
            else:
                self.driver = webdriver.Chrome()
        else:
            self.driver = None
    
    def close_spider(self, spider):
        if self.driver is not None:
            self.driver.quit()
        
    def start_requests(self):
        yield scrapy.Request(url=get_proxy_url(self.start_urls[0]), callback=self.parse)
    
    def parse(self, response):
        sub_categories = response.xpath("//ul[@class='ul-nav-folder']/li/a")
        for sub_category in sub_categories:
            if self.subs is not None:
                sub_category_title = sub_category.css("a::attr(title)").get()
                sub_category_title = remove_diacritics(sub_category_title).lower()
                if sub_category_title not in self.subs:
                    continue
            sub_category_url = sub_category.css("a ::attr(href)").get()
            yield response.follow(get_proxy_url(sub_category_url), callback=self.parse_sub_category_page)

    def parse_sub_category_page(self, response):
        if self.showmore_times != 0:
            self.driver.get(response.url)
            for _ in range(self.showmore_times):
                try:
                    wait = WebDriverWait(self.driver, 10)
                    view_more_button = self.driver.find_element(By.ID, "btn_view_more")
                    self.driver.execute_script("arguments[0].click();", view_more_button)
                    wait.until(
                        EC.presence_of_element_located((By.XPATH, "//*[@id='list-podcast-lazy']//article[contains(@class, 'item_loading')]"))  # Adjust as needed
                    )
                    wait.until(
                        EC.invisibility_of_element_located((By.XPATH, "//*[@id='list-podcast-lazy']//article[contains(@class, 'item_loading')]"))
                    )
                except Exception as e:
                    print(f"There are some errors, maybe because of internet")
                    break
                
            response = HtmlResponse(url=self.driver.current_url, body=self.driver.page_source, encoding='utf-8')
            
        podcast_pages = response.xpath("//*[@id='btn_view_more']/preceding::*[contains(@class,'title-news')]/a")
        for podcast_page in podcast_pages:
            podcast_page_url = podcast_page.css("a::attr(href)").get()
            yield response.follow(get_proxy_url(podcast_page_url), callback=self.parse_podcast_page)
            
    def parse_podcast_page(self, response):
        
        podcast_item = PodcastItem()
        
        # ---- url -----
        podcast_item['url'] = response.url
        
        # --- category ----
        podcast_item['main_category'] = "Podcasts"
        podcast_item['sub_category'] = response.xpath("//ul[@class='ul-nav-folder']/li[@class='active']/a/text()").get()
        
        # ------ title ------
        podcast_item['title'] = response.xpath("//h1[contains(@class,'title-news')]/text()").get()
        
        # ----- author ------
        podcast_item['author'] = response.xpath("//span[@class='author-in-player']/text()").get()
        
        # ----- release_date ------
        podcast_item['release_date'] = response.xpath("//span[@class='date']/text()").get()
        
        # ----- description ------
        podcast_item['description'] = response.xpath("//h1[contains(@class,'title-news')]/following-sibling::p[@class='description']/text()").get()
        
        # ----- relating_podcast ------
        podcast_item['relating_podcast'] = response.xpath("//audio/@src").get()

        yield podcast_item
        

    
