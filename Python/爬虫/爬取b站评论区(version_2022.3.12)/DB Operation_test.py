import mysql.connector
import emoji
#获取数据库连接
    #参数：host：主机IP,user：用户名,passwd：密码,database：数据库名
def getConnect(host,user,passwd,database):
    con=mysql.connector.connect(host=host,user=user,passwd=passwd,database=database,charset='utf8')
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
   
    