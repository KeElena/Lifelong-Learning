# -*- coding: utf-8 -*-
import mysql.connector
from datetime import datetime
import emoji
#获取数据库连接
    #参数：host：主机IP,user：用户名,passwd：密码,database：数据库名
def getConnect(host,user,passwd,database):
    con=mysql.connector.connect(host=host,user=user,passwd=passwd,database=database)
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
def store(uri,data_all,num_main,num_second,con,game=''):
    
    print('准备入库')
    try:
        dt=datetime.now().isoformat(sep=' ')
        dt=dt[0:16]
        sql='INSERT INTO '+game+'_datasource(uri,operate_time,num_main,num_second) VALUES (%s,%s,%s,%s)'
        insert(con,sql,[tuple([uri,dt,num_main,num_second])])
    except:
        pass
    print('准备完毕开始入库')
    sql='INSERT INTO '+game+'_comment(uid,`name`,`level`,datetime,`comment`,`like`,uri) VALUES (%s,%s,%s,%s,%s,%s,%s)'
    insert(con,sql,data_all)
    print('入库完毕')
    