from selenium import webdriver
from selenium.webdriver.common.by import By
import time

def getPage(browser):
    
    text=''
    print('爬取页面')
    while(True):
        
        browser.execute_script("window.scrollTo(0, document.body.scrollHeight);")
        time.sleep(3)
        try:

            element=browser.find_element(By.CLASS_NAME,'loading-state')
            text=element.text

        except:
            pass

        if(text=='没有更多评论'):
            break
        
    content=browser.page_source   
    browser.close()
    print("页面爬取完成")

    return content