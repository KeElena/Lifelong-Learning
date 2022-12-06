from selenium import webdriver
from selenium.webdriver.common.by import By
from time import sleep
from bs4 import  BeautifulSoup
import os
import urllib

def getBrowser(url):
    
    options=webdriver.ChromeOptions()
    options.add_experimental_option('excludeSwitches', ['enable-automation'])
    options.add_experimental_option('useAutomationExtension', False)


    browser = webdriver.Chrome(options=options)
    
    browser.execute_cdp_cmd(
            'Page.addScriptToEvaluateOnNewDocument',
            {'source':'Object.defineProperty(navigator,"webdriver",{get:()=>undefined})'}
            )
    
    browser.get(url)
    sleep(3)
    browser.execute_script("var q=document.documentElement.scrollTop=500")    
    sleep(2)
    print("浏览器中打开页面")
    return browser

def getEmoji(Xpth,browser):
    browser.find_element(By.XPATH,Xpth).click()
    sleep(2)
    page=browser.page_source
    soup=BeautifulSoup(page,'lxml')
    content=soup.find('div',{'class':'mhy-emoticon__list'})
    emoji_all=[]
    for i in content:
        emoji=[]
        name=i.get('title')
        url=i.get('style')
        x=url.find('https')
        y=url.find('png')
        url=url[x:y+3]
        emoji.append(name)
        emoji.append(url)
        emoji_all.append(tuple(emoji))
    return emoji_all

def download(total_emoji=[],addr=r'C:\Users\21927\Desktop\hoyo_emoji'):

    for i in total_emoji:
        img=urllib.request.urlopen(i[1]).read()
        f=open(addr+'//'+str(i[0])+'.png','wb')
        f.write(img)
        f.close()
        
def main():
    
    emoji_lists=['//*[@id="reply"]/div[1]/div[2]/div[2]/div[3]/div/div[2]/div[1]/div[1]/div',
                 '//*[@id="reply"]/div[1]/div[2]/div[2]/div[3]/div/div[2]/div[1]/div[2]/div',
                 '//*[@id="reply"]/div[1]/div[2]/div[2]/div[3]/div/div[2]/div[1]/div[3]/div',
                 '//*[@id="reply"]/div[1]/div[2]/div[2]/div[3]/div/div[2]/div[1]/div[4]/div',
                 '//*[@id="reply"]/div[1]/div[2]/div[2]/div[3]/div/div[2]/div[1]/div[5]/div']
    
    browser=getBrowser('https://bbs.mihoyo.com/bh3/article/9522955')
    addr=r'C:\Users\21927\Desktop\hoyo_emoji'

    browser.find_element(By.XPATH,'//*[@id="reply"]/div[1]/div[2]/div[2]/div[1]/i[1]').click()
    
    try:
        if addr not in os.listdir():
            os.mkdir(addr)    # 创建文件夹
    except:
        pass
    os.chdir(addr)        #打开文件夹
    
    for Xpth in emoji_lists:
        download(getEmoji(Xpth,browser),addr)
        
    """换页"""
    browser.find_element(By.XPATH,'//*[@id="reply"]/div[1]/div[2]/div[2]/div[3]/div/div[2]/div[2]/div[2]').click()
    
    emoji_lists=['//*[@id="reply"]/div[1]/div[2]/div[2]/div[3]/div/div[2]/div[1]/div[1]/div',
                 '//*[@id="reply"]/div[1]/div[2]/div[2]/div[3]/div/div[2]/div[1]/div[2]/div',
                 '//*[@id="reply"]/div[1]/div[2]/div[2]/div[3]/div/div[2]/div[1]/div[3]/div']
    
    for Xpth in emoji_lists:
        download(getEmoji(Xpth,browser),addr)
    
    browser.close()
    print('完成')
        
if __name__=='__main__':
    main()