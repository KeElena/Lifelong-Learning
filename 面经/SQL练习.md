# SELECT

**零、基本句式结构**

```sql
SELECT 字段 FROM 表;
```

**一、获取 Customers 表中的所有列**

```sql 
__ * FROM Customers;
```

**二、从 Customers 表中选择 City 列**

```sql
__ __ __ Customers;
```

**三、从 Customers 表中的 Country列中选择所有不同值**

```sql
__ __ Country FROM Customers;
```

**四、答案**

```sql
1、SELECT
2、SELECT City FROM 
3、SELECT DISTINCT
```

# WHERE

**零、基本句式结构**

```sql
SELECT 字段 FROM 表 WHERE 条件;
```

**一、选择City列的值为Berlin的所有记录**

```sql
SELECT * FROM Customers __='__';
```

**二、使用 NOT 关键字选择 City不是Berlin的所有记录**

```sql
SELECT * FROM Customers ____ ='__';
```

**三、选择 CustomerID 列的值为 32 的所有记录**

```sql
SELECT * FROM Customers __ CustomerID ____;
```

**四、选择 City 列的值为 Berlin 的所有记录 并且PostalCode列的值为 12209**

```sql
__ * FROM Customers __ City = 'Berlin' __ ___ =12209;
```

**五、选择 City列的值为Berlin或London的所有记录**

```sql
__ * FROM Customers __ City = 'Berlin' __ __ ='___';
```

**六、答案**

```sql
1、WHERE City ='Berlin'
2、WHERE NOT City='Berlin'
3、WHERE CustomerID =32
4、SELECT * FROM Customers WHERE City= 'Berlin' AND PostalCode=12209;
5、SELECT * FROM Customers WHERE City='Berlin' OR City='London';
```

# ORDER BY

**零、基本句式结构**

```sql
SELECT * FROM 表 ORDER BY 字段 排序顺序;
```

**一、从 Customers 表中选择所有记录，按 City列的字母顺序对结果进行排序**

```sql
SELECT * FROM Customers ____ __;
```

**二、从 Customers 表中选择所有记录，按 City进行倒序排序**

```sql
SELECT * FROM Customers _____ __ __;
```

**三、从 Customers 表中选择所有记录， 对结果进行排序，首先按列 Country， 然后，按列 City**

```sql
SELECT * FROM Customers ORDER BY Country ASC,City ASC;
```

**四、答案**

```sql
1、 ORDER BY City
2、 ORDER BY City DESC
3、ORDER BY Country,City
```

# INSERT

**零、基本句式结构**

```sql
INSERT INTO 表(字段) VALUES(值);
```

**一、在Customers表中插入一条新记录**

```sql
____ Customers _ CustomerName,Address,City,PostalCode,Country _ __ _ 'Hekkan Burger', 'Gateveien 15', 'Sandnes', '4306', 'Norway' _;
```

**二、答案**

```sql
1、INSERT INTO Customers (...) VALUES(...);
```

# NULL

**零、基本句式结构**

```SQL
//查询空
SELECT * FROM 表 WHERE 字段 IS NULL;
//查询非空
SELECT * FROM 表 WHERE 字段 IS NOT NULL;
```

**一、从Customers中选择PostalCode为空的所有记录**

```**sql
SELECT * FROM Customers WHERE __ __ __;
```

**二、从Customers中选择所有记录，其中PostalCode列不为空**

```sql
SELECT * FROM Customers WHERE __ __ __ __;
```

**三、答案**

```sql
1、 SELECT * FROM Customers WHERE PostallCode IS NULL;
2、 SELECT * FROM Customers WHERE WHERE PostalCode IS NOT NULL;
```

# UPDATE

**零、基本句式结构**

```sql
//单字段修改
UPDATE 表 SET 字段=值 WHERE 条件;
//多字段修改
UPDATE 表 SET 字段=值,字段=值 WHERE 条件;
```

**一、更新 Customers 表中所有记录的 City列**

```sql
__ Customers __ City ='Oslo';
```

**二、修改Country为挪威的记录的City为  Oslo**

```sql
__ Customers __ City = 'Oslo' WHERE Country ='Norway';
```

**三、更新 City 值和Country值**

```sql
UPDATE Customer SET CITY='Oslo',Couuntry='Norway' WHERE CustomerID=32;
```

**答案**

```sql
1、 UPDATE Customers SET City ='Oslo';
2、 UPDATE Customers SET City ='Oslo' WHERE Country ='Norway';
3、 UPDATE Customer SET CITY='Oslo',Couuntry='Norway' WHERE CustomerID=32;
```

# DELETE

**零、基本句式结构**

```sql
//删除记录
DELETE FROM 表 WHERE 条件;
//清空表
DELETE FROM 表;
```

**一、从Customers表中删除Country值为Norway的所有记录**

```sql
DELETE FROM Customers WHERE Country='Norway';
```

**二、清空表**

```sql
DELETE FROM Customers;
```

# SQL函数

**零、基本句式结构**

```sql
1、 获取最小值
SELECT MIN(字段) FROM 表;
2、 获取最大值
SELECT MAX(字段) FROM 表;
3、 获取记录数
SELECT count(*) FROM 表;
4、 获取平均价格
SELECT AVG(字段) FROM 表;
5、 获取总和
SELECT SUM(字段) FROM 表;
```

**一、选择 Price列值最小的记录**

```sql
SELECT MIN(Price) FROM Products;
```

**二、 Price 列中值最高的记录**

```sql
SELECT MAX(Price) FROM Products;
```

**三、返回 Price 值设置为18的记录数**

```sql
SELECT COUNT(*) FROM Products WHERE Price=18;
```

**四、计算所有产品的平均价格**

```sql
SELECT AVG(Price) FROM Products;
```

**五、Products表中所有Price列值的总和**

```sql
SELECT SUM(Price) FROM Products;
```

# LIKE

**零、基本句式结构**

```sql
#多字母模糊匹配
SELECT * FROM 表 WHERE 字段 LIKE 'a%';
#单字母模糊匹配
SELECT * FROM 表 WHERE 字段 LIKE 'a_';
#模糊取反
SELECT * FROM 表 WHERE 字段 NOT LIKE 'a%';
#多选择匹配
SELECT * FROM 表 WHERE 字段 LIKE '[abc]%';
SELECT * FROM 表 WHERE 字段 LIKE '[a-f]%'; #a到f
#多选择匹配取反
SELECT * FROM 表 WHERE 字段 LIKE '[^abc]%';
```

**一、选择City列的值以字母a开头的所有记录**

```sql
SELECT * FROM Customers WHERE ___________;
```

**二、选择City列的值以字母a结束的所有记录**

```sql
SELECT * FROM Customers WHERE ___________;
```

**三、选择City列的值包含字母a的所有记录**

```sql
SELECT * FROM Customers WHERE ___________;
```

**四、选择City列的值以字母a开头并以字母b结尾的所有记录**

```sql
SELECT * FROM Customers WHERE ___________;
```

**五、选择City列的值不以字母a开头的所有记录**

```sql
SELECT * FROM Customers WHERE ___________;
```

**六、选择City的第二个字母为a的所有记录**

```sql
SELECT * FROM Customers WHERE ___________;
```

**七、选择City的第一个字母为a或c或s的所有记录**

```sql
SELECT * FROM Customers WHERE ___________;
```

**八、选择City的第一个字母以从a到f的任意字符开头的所有记录**

```sql
SELECT * FROM Customers WHERE ___________;
```

**九、选择City不以acf开头的所有记录**

```sql
SELECT * FROM Customer WHERE ___________;
```

**答案**

```sql
1、 SELECT * FROM Customers WHERE City LIKE 'a%';
2、 SELECT * FROM Customers WHERE City LIKE '%a';
3、 SELECT * FROM Customers WHERE City LIKE '%a%';
4、 SELECT * FROM Customers WHERE City LIKE 'a%b';
5、 SELECT * FROM Customers WHERE City NOT LIKE 'a%';
6、 SELECT * FROM Customers WHERE City LIKE '_a%';
7、 SELECT * FROM Customers WHERE City LIKE '[acs]%';
8、 SELECT * FROM Customers WHERE City Like '[a-f]%';
9、 SELECT * FROM Customer WHERE City Like '[^acf]%';
```

# IN

**零、基本句式结构**

```sql
SELECT * FROM 表 WHERE 字段 IN(值1 值2);
```

**一、获取Country值为Norway或France的所有记录**

```sql
SELECT * FROM Customers WHERE __________;
```

**二、、获取Country值不是Norway或France的所有记录**

```sql
SELECT * FROM Customers WHERE __________;
```

**三、答案**

```sql
1、 SELECT * FROM Customers WHERE Country IN('Norway','France');
2、 SELECT * FROM Customers WHERE Country NOT IN('Norway','France');
```

# BETWEEN

**零、基本句式结构**

```sql
#闭区间
SELECT * FROM Products WHERE 字段 BETWEEN 值1 AND 值2;
```

**一、列出`Price`值在10到20之间的所有记录**

```sql
#[10,20]
SELECT * FROM Products WHERE _________________;
```

**二、列出`Price`值不在10到20之间的所有记录**

```sql
#![10,20]
SELECT * FROM Products WHERE _________________;
```

**三、选择在“Geitost”和“Pavlova”之间的所有记录**

```sql
SELECT * FROM Products WHERE _________________;
```

**答案**

```sql
1、 SELECT * FROM Products WHERE Price BETWEEN 10 AND 20;
2、 SELECT * FROM Products WHERE Price NOT BETWEEN 10 AND 20;
3、 SELECT * FROM Products WHERE ProductName BETWEEN 'Geitost' AND 'Pavlova';
```

# ALIAS

**零、基本句式结构**

```sql
SELECT 字段 AS 别名 FROM 表名 AS 表别名;
```

**一、字段PostalCode使用Pno别名输出**

```sql
SELECT PostalCode _____ FROM Customers;
```

**二、给表起别名为Consumers**

```sql
SELECT * FROM Customers _____;
```

**答案**

```sql
SELECT PostalCode AS Pno FROM Customers;
SELECT * FROM Customers AS Consumers;
```

# JOIN

**零、基本句式结构**

```sql
#外连接，允许左表出现不匹配的记录
SELECT * FROM 表1 LEFT JOIN 表2 ON 表1.ID=表2.ID;
#内连接，只筛选匹配的记录
SELECT * FROM 表1 INNER JOIN 表2 ON 表1.ID=表2.ID;
```

**一、使用外连接连接Orders和Customers并获取所有记录**

```sql
SELECT * FROM Orders LEFT JOIN Customers ON Orders.CustomerId=Customers.CustomerId;
```

**二、所有内连接连接Orders和Customers并获取所有记录**

```sql
SELECT * FROM Orders INNER JOIN Customers ON Orders.CustomerId=Customers.CustomerId;
```

# GROUP BY

**零、基本句式结构**

```sql
SELECT COUNT(*) FROM 表 GROUP BY 字段;
```

**一、列出每个国家/地区的客户数量**

```sql
SELECT COUNT(CustomerID),Country FROM Customers GROUP BY Country;
```

**二、列出每个国家/地区的客户数量，按客户最多的国家/地区优先排序**

```sql
#方法一
SELECT COUNT(CustomerID),Country FROM Customers GROUP BY CustomerID ORDER BY COUNT(CustomerID) DESC;
#方法二
SELECT COUNT(CustomerID) AS num,Country FROM Customers GROUP BY CustomerID ORDER BY num DESC;
```

# 引用

[sql 代码练习测试](https://www.w3schools.cn/sql/exercise.asp)

