from selenium import webdriver

def getBrowser(url):
    
    options=webdriver.ChromeOptions()
    options.add_experimental_option('excludeSwitches', ['enable-automation'])
    options.add_experimental_option('useAutomationExtension', False)
    options.add_argument('--headless')

    browser = webdriver.Chrome(options=options)
    
    browser.execute_cdp_cmd(
            'Page.addScriptToEvaluateOnNewDocument',
            {'source':'Object.defineProperty(navigator,"webdriver",{get:()=>undefined})'}
            )
    
    browser.get(url)
    print("浏览器中打开页面")
    return browser

if __name__=='__main__':
    browser=getBrowser('https://www.bilibili.com/')
    browser.close()
    