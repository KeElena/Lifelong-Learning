from bs4 import  BeautifulSoup
from datetime import datetime 
import emoji
def getValue(content,uri):
    
    soup=BeautifulSoup(content,'lxml')
    
    replys=[]               #统计楼中楼
    data_all=[]                #汇总数据
    
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
            print(temp)
            if temp !=[]: 
                for z in temp:
                    img=z.get('alt')
                    img=str(img)
                    imgs=imgs+img           
        comment=comment+imgs
        comment=emoji.demojize(comment)
        data.append(comment)

        like=i.find('span',{'class':'like'}).get_text()                 #点赞
        if like=='':
            like='0'
        data.append(like)
        
        data.append(uri)                                                #插入uri
        
        reply = i.find('div', {'class': 'view-more'})                   #楼中楼情况

        try:
            num=reply.text
            num= num.replace('条回复, 点击查看', '')
            num = num.replace('共', '')
            replys.append(int(num))

        except:
            pass
        
        data_all.append(tuple(data))                                        #列表转化为元组再给data列表

    #汇总
    num_main=len(data_all)                                                  #计算主评论数量
    print('主评论数量：',num_main)
    
    num_second=0                                                             #计算楼中楼
    for i in replys:
        num_second=num_second+i
    
    print('子评论数量：',num_second)
    print('所有评论数量：',num_main+num_second)
    
    return data_all,num_main,num_second