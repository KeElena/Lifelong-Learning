# -*- coding: utf-8 -*-
from selenium import webdriver
from selenium.webdriver.common.by import By
import time
from bs4 import  BeautifulSoup
from datetime import datetime
import pymysql
import gc

#合取浏览器对象
def getBrowser(url):
    
    options=webdriver.ChromeOptions()
    options.add_experimental_option('excludeSwitches', ['enable-automation'])
    options.add_experimental_option('useAutomationExtension', False)
    options.add_argument('--disable-gpu')
    options.add_argument('--headless')

    browser = webdriver.Chrome(options=options)
    
    browser.execute_cdp_cmd(
            'Page.addScriptToEvaluateOnNewDocument',
            {'source':'Object.defineProperty(navigator,"webdriver",{get:()=>undefined})'}
            )
    
    browser.get(url)
    print("浏览器中打开页面")
    return browser
#获取页面
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

#获取有效数据
def getValue(content,uri):
    
    soup=BeautifulSoup(content,'lxml')
    
    replys=[]               #统计楼中楼
    data_all=[]                #汇总数据
    compositions=[]
    
    contents = soup.find('div', {'class':'comment-list'})
    #收集每条评论的信息
    for i in contents:

        data=[]
        
        uid= i.find('a').get('data-usercard-mid')                       #uid
        data.append(uid)
        
        name = i.find('a', {'class': {'name'}}).get_text()              #昵称
        data.append(name)
        
        level = i.find('img',{'class':'level'}).get('src')              #等级
        num=level.find("level_")
        data.append(level[num+6:num+7])
        
        dt=datetime.now()
        date_time=i.find('span',{'class':'time'}).get_text()            #日期时间
        if '小时' in date_time:
            dt=dt.replace(hour=int(date_time[0:1]))
            date_time=dt.isoformat(sep=' ')
            date_time=date_time[0:16]

        if '分钟' in date_time:
            dt=dt.replace(minute=int(date_time[0:1]))
            date_time=dt.isoformat(sep=' ')
            date_time=date_time[0:16]
       
        if '秒' in date_time:
            dt=dt.replace(second=int(date_time[0:1]))
            date_time=dt.isoformat(sep=' ')   
            date_time=date_time[0:16]
        data.append(date_time)       

        comments=i.find_all('p', {'class': 'text'})                     #评论
        comment=comments[0].get_text()
        imgs=''        
        for j in comments:
            temp=j.find_all('img')
            if temp !=[]: 
                for z in temp:
                    img=z.get('alt')
                    img=str(img)
                    imgs=imgs+img   
        comment=comment+imgs
        data.append(comment)

        like=i.find('span',{'class':'like'}).get_text()                 #点赞
        if like=='':
            like='0'
        data.append(like)
        
        data.append(uri)                                                #插入uri
        
        reply = i.find('div', {'class': 'view-more'})                   #楼中楼情况
        amount=0
        try:
            num=reply.text
            num= num.replace('条回复, 点击查看', '')
            num = num.replace('共', '')
            amount=int(num)
        except:
            pass
        
        reply=i.find_all('div',{'class','reply-item reply-wrap'})
        if len(reply)>amount:
            amount=len(reply)
        replys.append(amount)        

        if(len(data[4]) > 100):                                         #列表转化为元组再给data列表
            compositions.append(tuple(data))
        else:
            data_all.append(tuple(data))                                    

    #汇总
    num_composition=len(compositions)
    num_main=len(data_all)+num_composition                                 #计算主评论数量
    print('主评论数量：',num_main)
    
    num_second=0                                                             #计算楼中楼
    for i in replys:
        num_second=num_second+i
    
    print('小作文数量：',num_composition)
    print('子评论数量：',num_second)
    print('所有评论数量：',num_main+num_second)
    
    return data_all,num_main,num_second,compositions,num_composition

#获取数据库连接
    #参数：host：主机IP,user：用户名,passwd：密码,database：数据库名
def getConnect(host,user,passwd,database):
    con=pymysql.connect(host=host,user=user,passwd=passwd,database=database)
    return con
#数据库查询
    #con:=：数据库连接，sql：sql语句，val列表：预编译数据
def select(con,sql,val=[]):
    mycursor=con.cursor()
    mycursor.execute(sql,val)
    result=mycursor.fetchall()
    return result
#数据库插入
def insert(con,sql,val=[]):
    #con:=：数据库连接，sql：sql语句，val列表：预编译数据
    mycursor=con.cursor()
    mycursor.executemany(sql,val)
    con.commit()
    return True
#数据库更新
    #con:=：数据库连接，sql：sql语句，val列表：预编译数据
def update(con,sql,val=[]):
    mycursor=con.cursor()
    mycursor.execute(sql,val)
    con.commit()
    return True

def delete(con,sql,val=[]):
    #con:=：数据库连接，sql：sql语句，val列表：预编译数据
    mycursor=con.cursor()
    mycursor.execute(sql,val)
    con.commit()
    return True
#数据入库
def befor(uri,num_main,num_second,con,num_composition,game='',official='',event=''):
    
    print('准备入库')
    try:
        dt=datetime.now().isoformat(sep=' ')
        dt=dt[0:16]
        sql='INSERT INTO '+game+'_datasource(uri,operate_time,num_main,num_second,`num_composition`,`official`,`event`) VALUES (%s,%s,%s,%s,%s,%s,%s)'
        insert(con,sql,[tuple([uri,dt,num_main,num_second,num_composition,official,event])])
    except:
        pass
    print('开始入库')    

def store(data_all,con,game='',type=0,uri=''):
    if type==0:
        sql='INSERT INTO '+game+'_comment(uid,`name`,`level`,datetime,`comment`,`like`,uri) VALUES (%s,%s,%s,%s,%s,%s,%s)'
        insert(con,sql,data_all)
        print('入库完毕')
    else:
        sql='INSERT INTO composition(uid,`name`,`level`,datetime,`comment`,`like`,uri) VALUES (%s,%s,%s,%s,%s,%s,%s)'
        insert(con,sql,data_all)
        try:
            sql='UPDATE composition SET `game`=%s WHERE uri=%s'
            update(con,sql,val=[tuple([game,uri])])
        except:
            pass
        print('小作文入库完毕')   

if __name__=='__main__':
    
    uris=[]
    
    game='' 
    official=''
    event=''    
    
    for i in uris:
        #输入参数
        uri=i
        #爬虫执行
        browser=getBrowser(uri)
        content=getPage(browser)
        del browser
        #本地存储
        f=open(r'./demo.txt','w',encoding='utf-8')
        f.write(content)
        f.close()        
        #解析存储
        data_all,num_main,num_second,compositions,num_composition=getValue(content,uri)
        del content
        
        con=getConnect('localhost','root','password','game')

        befor(uri,num_main,num_second,con,num_composition,game,official,event)
        store(data_all,con,game)
        store(compositions,con,game,1,uri)
        
        del data_all,num_main,num_second,compositions,num_composition
        
        con.close()
        #清理内存
        gc.collect() 
