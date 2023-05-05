# JavaScript简介

**一、JavaScript**

* JavaScript是浏览器脚本语言
* 一种具有函数优先的轻量级，解释型或即时编译型的编程语言
* JavaScript严格区分大小写
* 在浏览器的控制台进行调试

**二、JavaScript常见框架**

* `Vue`
* `React`

**三、UI框架**

* `Ant-Design`：阿里巴巴开发，基于React的UI框架
* `ElementUI.iview,ice`：饿了么开发，基于Vue的UI框架
* `WeUI`：微信小程序

**四、后端框架**

* `node.js`

# JavaScript的导入

**一、内部导入**

* 使用`script`标签导入`JavaScript`代码

```html
<html>
    <head>
        <script>
        	alert("hello world");
        </script>
    </head>
</html>
```

**二、外部导入**

* 新建一个`JavaScript`文件，使用`script`导入内容
* 使用`src`属性指定内容
* 要求`script`标签要成对

```html
<html>
    <head>
        <script src="./demo.js"></script>
    </head>
</html>
```

```javascript
alert("hello world");
```

# 浏览器控制台

| 标签页                  | 常见作用           |
| ----------------------- | ------------------ |
| Elements（元素）        | 爬取网站内容       |
| Console（控制台）       | 调试JavaScript     |
| Sources（源代码）       | 断点调试JavaScript |
| Network（网络）         | 抓包               |
| Application（应用程序） | 查看网站Cookie     |

# 变量与类型

## 定义变量

**一、使用let定义局部变量**

* 使用`let`定义局部变量
* JavaScript严格区分大小写
* 定义在函数、for循环、if等内部时，只作用于内部

```javascript
let a=1;
```

**二、使用const定义常量**

* 使用`const`定义常量
* JavaScript严格区分大小写
* 使用`const`定义全局对象时，对象的值可以修改

```javascript
const a=1;
const car = {type:"Fiat", model:"500", color:"white"};
car.color="red";
```

**三、使用var定义全局**

* `ES6`以后的版本不推荐在局部区块使用，会引起变量提升等错误

```javascript
var a=1;
```

## 常见的数据类型

**一、number**

* 表示整数、浮点数、科学计数法、负数、Nan（not a number）、Infinity（无限大）
* 使用浮点数时存在精度的损失，尽量避免使用浮点数进行运算

**二、字符串**

* 使用单引号或双引号

**三、布尔值**

* true、false

**四、逻辑运算**

* &&：与
* ||：或
* !：非

**五、比较运算符**

* =：赋值
* ==：等于（类型不一样值一样）
* 绝对等于（类型一样，值一样）
* NaN与所有的数值都不相等，包括NaN自己，只能通过`isNaN()`函数判断值是否是NaN

**六、null和undefined**

* null：空
* undefined：未定义

**七、数组**

* JS中不需要数组的元素是一系列相同类型的元素
* 开发过程中建议数组是一系列相同类型的元素
* 取元素时，如果超过数组索引则会返回undefined

```javascript
let arr=[1,2,3,"hello",null,true];
```

**八、对象**

```javascript
var person={
    name:"demo",
    age:3,
    tags:['C','java','Python','Go'],
}
```

## 严格检查模式

**一、使用严格检查模式**

* 必须写在脚本的第一行

```javascript
'use strict';
```

## 打印与弹窗

**一、打印**

* 打印内容只在浏览器控制台输出
* 使用`console.log()`方法实现打印

```javascript
let score=80;
console.log(score);
```

**二、弹窗**

* 使用`alert()`方法实现弹窗

```javascript
alert("hello world");
```

# 数据类型

## 字符串类型

**一、字符串相关规则**

* 使用单引号和双引号包裹
* 使用特殊字符时使用`\`转义
* 不能改变字符串的元素，只能通过新的字符串代替原来的字符串

**二、特殊字符的表示**

```javascript
\'			//特殊字符
\n			//换行
\t			//表格符
\u4e2d		//utf-8
\x41		//ASCII
```

**三、多行字符串**

* 使用反引号编写多行字符串

```javascript
let str=
`
hello
world
`;
```

**四、形式字符串**

* 构建形式化字符串时使用反引号
* 在字符串中使用`${变量}`实现字符串的形式化（类似shell）

```javascript
let name="demo";
let age=13;
var msg=`你好${name}，今年${age}岁`;
console.log(msg);
```

**五、获取字符串长度**

* 使用`str.length`获取字符串长度

```javascript
let str="hello world";
let n=str.length;
console.log(n);
```

**六、大小写转换**

* 使用`str.toUpperCase()`方法将字符串变成大写
* 使用`str.toLowerCase()`方法将字符串变成小写

```javascript
let str="hello world";
str=str.toUpperCase();
console.log(str);
str=str.toLowerCase();
console.log(str);
```

**七、获取子串的索引位置**

* 使用`str.indexOf(子串)`获取子串的索引（从0开始）

```javascript
let str="hello world";
let idx=str.indexOf("wor");
console.log(idx);
```

**八、截取字符串**

* 使用`str.substring(开始,结束)`方法截取字符串
* 开始使用和结束索引满足左闭右开原则

```javascript
let str="hello world";
let newstr=str.substring(1,5);
console.log(newstr);
```

**九、拼接字符串**

* 使用`+`拼接字符串

```javascript
let str="abc"+"def";
```

## 数组类型

**一、数组的定义**

```javascript
//定义数组
let arr=[1,2,3,4,5];
//取值
let num=arr[1];
//赋值
arr[1]=10;
//多维数组的定义
let arr=[[1,2,3],[4,5,6],[7,8,9]];
//多维数组取值
let num=arr[0][2];
```

**二、数组长度与伸缩**

* `JS`可以修改数组的长度
* 扩展时多出来的元素为`undefined`
* 收缩时去除多余的元素

```javascript
//获取长度
arr.length;
//扩展长度
arr.length=10;
//收缩长度
arr.length=2;
```

**三、通过元素的值获取索引**

* 使用`arr.indexOf()`方法通过元素的值获取索引

```javascript
let idx=arr.indexOf(2)
```

**四、截取数组**

* 使用`arr.slice()`方法截取数组
* 使用索引进行截取，满足左闭右开原则

```javascript
let newarr=arr.slice(1,3);
```

**五、入栈与出栈**

* 使用`arr.push()`方法实现入栈操作，可一次添加多个值（尾部追加）
* 使用`arr.pop()`方法实现出栈操作，只能弹出一个元素（尾部弹出）
* 入栈时，如果没有undefined元素则会自动改变数组长度，出栈时会自动改变数组长度

```javascript
//入栈
arr.push(7,8,9);
//出栈
arr.pop();
```

**六、头插**

* 使用`arr.unshift()`方法在头部插入多个值

```javascript
arr.unshift(1,2,3);
```

**七、头部弹出**

* 使用`arr.shift()`将元素从头部弹出一个元素

```javascript
arr.shift();
```

**八、排序**

* 使用`arr.sort()`方法排序
* 数字在前，字母在后（字符串的数字与数字归类在数字里）

```javascript
let arr=[2,1,5,3,6,8,9,7];
arr.sort();
console.log(arr);
```

**九、元素反转**

* 使用`arr.reverse()`方法实现元素反转

```javascript
arr.reverse();
```

**十、拼接数组**

* 使用`arr.concat()`方法拼接数组
* 不会改变原来的数组，需要一个数组接收拼接后的数组

```javascript
let arr1=[1,2,3];
let arr2=["a","b","c"];
let newarr=arr1.concat(arr2);
console.log(newarr);
```

**十一、拼接数组所有元素**

* 使用`arr.join()`方法拼接数组
* 可以输入一个拼接字符串
* 返回字符串类型的数据

```javascript
let arr=[1,2,3,4];
let str=arr.join("-");
console.log(str);
```

## 对象类型

**一、对象的定义**

* 定义时用`:`赋值
* 多属性`,`隔开，最后一个属性不加`,`
* 键都是字符串，值是任意类型

```javascript
let 对象名={
    key:value,
    key:value
}
```

**二、对象赋值**

* 使用`=`赋值

```javascript
let person={
    name:"demo",
    age:3
}
person.name="hello";
person.age=10;
```

**三、动态删除对象里的属性**

* 使用`delete`删除对象里某个属性

```javascript
delete person.name;
```

**四、动态添加属性**

* 定义一个对象不存在的属性并赋值

```javascript
let person={
    name:"demo",
    age:3
}
person.sex="boy";
```

**五、判断对象是否存在某属性**

* 键用字符串表示

```javascript
let person={
    name:"demo",
    age:3
}
"name" in person;
```

**六、判断属性是否是对象新增的而表示继承得到的**

* 使用`hasOwnProperty()`判断

```javascript
person.hasOwnProperty("name");
```

# 基础的语句结构

## If else语句

**一、语法结构**

```javascript
if (比较式){
    
}else if(比较式){
    
}else{
    
}
```

**二、使用示例**

```html
<html>
    <head>
        <script>
        	let score=78;
            if(score<60){
                alert("不合格");
            }else if(score>=60 && score<80){
                alert("优良");
            }else{
                alert("优秀");
            }
        </script>
    </head>
</html>
```

## for循环

```javascript
//语法结构
for(let i=0;i<次数;i++){
    语句块
}
```

```javascript
for(let i=0;i<arr.length;i++){
    console.log(arr[i])
}
```

## for-in循环

```javascript
//语法结构
for(let idx in 对象数组){
    语句块
}
```

## while循环

```javascript
//语法结构
while(比较式){
    语句块
}
```

# 基础数据结构

## map

**一、定义map**

* 使用`new Map()`方法定义map数据

* 使用二维数组形式定义map数据

```javascript
let m=new Map([["name","hello world"],["age",4]]);
```

**二、获取map数据**

* 使用`get()`方法根据key获取值

```javascript
let m=new Map([["name","hello world"],["age",4]]);
let val=m.get("name");
```

**三、修改键值对的值或添加新的键值对**

* 使用`set()`方法设置新的键值对或修改键值对的值

```javascript
let m=new Map([["name","hello world"],["age",4]]);
m.set("sex","boy");
```

**四、删除键值对**

* 使用`delete()`方法根据key删除键值对

```javascript
let m=new Map([["name","hello world"],["age",4]]);
m.delete("name");
```

## set

**一、定义set集合**

* set是无序不重复集合
* 通过数组输入值
* 多余的元素直接去重

```javascript
let s = new Set([1,2,3,4,5]);
```

**二、添加元素**

* 使用`add()`添加元素

```javascript
let s = new Set([1,2,3,4,5]);
s.add(6);
```

**三、删除某个元素**

* 使用`delete()`方法删除

```javascript
let s = new Set([1,2,3,4,5]);
s.delete(1);
```

**四、判断是否有某值**

* 使用`has()`判断是否有某值

```javascript
let s = new Set([1,2,3,4,5]);
s.has(1);
```

# 函数

## 函数的定义

**一、定义函数**

* 使用`function()`定义函数和方法
* 如果没有执行return，函数执行完也会返回结果（`NaN`或`undefined`）
* 函数内定义变量时，会自动在执行函数前声明该变量（相当于先声明，后赋值）

```javascript
//function定义的格式
function 函数名(变量){
    代码块;
    return 变量或值;
}
```

```javascript
function demo(x,y){
    if(x>y){
        return x;
    }else{
        return y;
    }
}
```

**二、定义匿名函数**

* 使用`function()`定义匿名函数
* 匿名函数可被当作值进行接收

```javascript
//定义格式
var func =function(变量){
    代码块;
}
```

```javascript
var demo=function(x,y){
    if(x>y){
        return x;
    }else{
        return y;
    }
}
```

## `arguments`

**一、`arguments`作用**

* 将函数传进的所有参数集中到`arguments`中，`arguments`相当于数组
* 使用`arguments.length`获取函数传进参数的个数
* 可通过索引获取`arguments`的值（`arguments[i]`）

```javascript
function demo(x){
    for(let i=0;i<arguments.length;i++){
        console.log(arguments[i]);
    }
}
```

## `rest`

**一、`rest`**

* 如果函数的需要a个参数，而传入多余的参数只能通过for循环得到，必须要跳过前面a个参数
* 使用rest获取除了定义的参数之外的所有参数
* 使用`...rest`表示其余参数作为一个数组

```javascript
function demo(x,...rest){
	console.log(rest)
}
```

## 变量的作用域

**一、var的作用域判断**

* 在函数内用var定义时，函数外不能取到值（局部）

**二、同名变量的处理**

* 函数内部定义了一个同名变量，如果与外部某变量同名，则屏蔽外部变量
* 如果函数内定义了一个与前面某个变量同名的变量，则屏蔽前面的变量

**三、函数内部变量的定义**

* 函数内定义变量时，会自动在执行函数前声明该变量（相当于先声明，后赋值）

**四、window变量**

* `javascript`将所有的全局变量都放在`window`对象里
* 需要使用`var`才能定义全局变量
* `javascript`的全局函数是window的方法

```javascript
var num=1;
console.log(window.num);
```

## 闭包

**一、闭包**

* 概念和go的闭包一样
* 如果函数内部没有某个变量，则在函数所处的环境查找是否有该变量
* 若有多层函数时，最里层的函数获取通过闭包取到最外层特有的变量

```javascript
function demo(){
    console.log(val)
}
let val=1;
demo()
```

# 方法

**一、定义方法**

* 使用`this`调用对象内的属性（this表示当前对象）
* <font color=red>调用对象的方法时要`()`否则无用</font>（`对象.方法()`）

```javascript
var 对象名={
    变量:值,
    方法名:function(输入参数){
        方法代码块
    }
}
```

```javascript
var person={
    name:"demo",
    getName:function(){
        console.log(this.name);
    }
}
person.getName();
```

**二、`apply`方法修改this指向的对象**

* 将某对象的方法牛到另一个对象上使用，要求另一个对象要满足能运行该方法的所需的条件（必要变量）
* 语法：`方法.apply(对象,[方法参数]);`
* 要求方法在外部定义对象内部绑定，否则外部不能调用`apply`方法

```javascript
//外部定义方法
let getName=function(){
        console.log(this.name);
}
//定义person对象
var person={
    name:"demo",
    getName:getName
}
//定义student对象
var student={
    name:"student"
}
//将person对象的getName方法应用到student中
getName.apply(student,[]);
```

# 内部对象

## 对象的类型

**一、类型种类**

* number：数字类型
* string：字符串类型
* boolean：布尔类型
* object：对象类型，数组等数据结构是object类型
* function：方法类型

## Date对象

**一、获取日期对象**

* 使用`new Date()`获取Date对象

```javascript
let now =new Date();	//获取日期对象
now.getFullYear();		//获取年
now.getMonth();			//获取月，需要加1
now.getDate();			//获取日
now.getHours();			//获取小时
now.getMinutes();		//获取小时
now.getSeconds();		//获取小时
now.getDay();			//获取星期
now.getTimr();			//获取时间戳
now.toLocaleString()	//获取时间字符串
```

**二、通过时间戳获取时间**

* 使用`new Date()`时传入时间戳，返回时间对象

```javascript
let last=new Date(1674116760942);
```

## JSON对象

**一、JSON格式**

* 对象用`{}`隔开
* 数组用`[]`隔开
* 键值对用`key:value`表示，同类对象的键值对用`,`隔开

**二、对象转为JSON数据**

* 使用`JSON.stringify()`方法将对象转为json对象

```javascript
let person={
    name:"demo",
    age:18,
    sex:"男"
};
let str=JSON.stringify(person);
console.log(str);
```

**三、JSON数据转为对象**

* 使用`JSON.parse(数据)`方法解析json数据

```javascript
str=`{"name":"demo","age":18,"sex":"男"}`
let demo=JSON.parse(str);
console.log(demo);
```

# 面向对象编程

## 原型对象继承的实现

**一、继承原型的属性和方法**

* 定义一个新的对象，使用`__proto__`继承某个已存在对象的属性和方法
* 如果新对象的属性与旧对象有冲突，则优先新对象的值
* 在`javascript`继承的对象可以随意改变，而其它后端语言不允许（过于随意）

```javascript
var student={
    name:"student",
    age:18,
    run:function(){
        console.log(this.name+" running",this.age)
    }
};
var demo={
    name:"demo"
};
demo.__proto__=student;
demo.run();
```

## class类的定义与对象的实例化

**一、类的定义**

* 使用`class`定义一个类
* `constructor()`为`javascript`统一的构造方法
* `javascript`里实例变量的定义和赋值统一在`constructor()`方法里完成
* 可以直接使用方法名定义新的方法

```javascript
class person{
    constructor(name){
        this.name=name;
    }
    hello(){
        console.log(this.name+" hello world");
    }
}
```

**二、通过类定义对象并使用方法**

* 使用`new`实例化对象

```javascript
let p =new person("demo");
p.hello();
```

## extends继承

**一、通过继承获得新的类**

* 使用`extends`实现继承（类似`javascript`）
* 使用`super()`给父类的属性赋值

```javascript
class person{
    constructor(name,age){
        this.name=name;
        this.age=age;
    }
}
class student extends person{
    constructor(name,age,grade){
        super(name,age);
        this.grade=grade;
    }
    getMsg(){
        console.log(this.name+" "+this.age+" in "+this.grade);
    }
}
let s=new student("demo",18,101);
s.getMsg();
```

# 操作BOM对象

## BOM相关概念

**一、BOM**

* 浏览器相关对象

## window对象

**一、弹窗**

* 使用`window.alert()`实现弹窗

```javascript
window.alert();
```

**二、获取页面宽度**

```javascript
window.innerWidth;
```

**三、获取页面高度**

```javascript
window.innerHeight;
```

**四、获取浏览器宽度**

```javascript
window.outerWidth;
```

**五、获取浏览器高度**

```javascript
window.outerHeight;
```

## navigator对象

* 不建议用于判断，可以被人为修改

**一、获取浏览器版本**

```javascript
navigator.appName;
```

**二、获取用户的UserAgent**

```javascript
navigator.userAgent;
```

**三、获取系统版本**

```javascript
navigator.platform;
```

## screen对象

**一、获取主机屏幕宽度**

```javascript
screen.width;
```

**二、获取主机屏幕高度**

```javascript
screen.height;
```

## location对象

**一、获取主机域名**

```javascript
location.host;
```

**二、获取链接**

```javascript
location.href;
```

**三、获取协议**

```javascript
location.protocol;
```

**四、刷新页面**

* 使用`location.reload()`方法实现

```javascript
location.reload();
```

**五、劫持网页**

* 使用`location.assign()`方法实现网页的跳转或劫持

```javascript
location.assign("https://www.baidu.com");
```

## document页面对象

**一、获取页面标题**

```javascript
document.title;
```

**二、通过Id选择器获取文档标签对象**

* 使用`document.getElentById()`获取文档标签对象

```javascript
let elem=document.getElementById('app');
```

**三、通过Class选择器获取文档标签对象**

* 使用`document.getElentByClassName()`获取文档标签对象

```javascript
let elem=document.getElementByClassName('app');
```

**四、通过标签选择器获取文档标签对象**

* 使用`document.getElentByTagName()`获取文档标签对象

```javascript
let elem=document.getElementByTagName('h1');
```

**五、获取Cookie**

* 使用`document.cookie`

```javascript
document.cookie
```

## history历史对象

* 浏览器的历史记录
* 不建议使用

**一、返回**

```javascript
history.back();
```

**二、前进**

```javascript
history.foward();
```

# DOM对象

* 文档对象模型

## 获取DOM对象

**一、获取DOM对象**

* 使用`document.getElentById()`通过id选择器获取文档标签对象

* 使用`document.getElentByClassName()`通过类选择器获取文档标签对象

* 使用`document.getElentByTagName()`通过标签选择器获取文档标签对象

```javascript
let elem=document.getElementById('app');
let elem=document.getElementByClassName('app');
let elem=document.getElementByTagName('h1');
```

**二、获取父节点下的所有子节点**

* 使用`children`获取父节点下的所有子标签

```javascript
let tags=father.children;
```

**三、获取第一个子标签**

* 使用`firstChild`获取第一个子标签

```javascript
let elem=father.firstChild;
```

**四、获取最后一个子标签**

* 使用`lastChild`获取最后一个子标签

```javascript
let elem=father.lastChild;
```

**五、获取同层相邻的标签对象**

* 使用`nextElementSibling`获取同层相邻的标签对象

```javascript
elem.nextSibling;
```

## 更新DOM对象

**一、改变标签文本**

* 使用`dom.innerText`改变文本

```javascript
elem.innerText="C/C++";
```

**二、改变标签文本并解析HTML内容**

* 使用`dom.innerHTML`改变dom文本并解析HTML内容

```javascript
elem.innerHTML="<strong>C/C++</strong>";
```

**三、改变文本颜色**

* 使用`dom.style.color`改变文本颜色

```javascript
elem.style.color="red";
```

**四、修改文本大小**

* 使用`dom.style.fontSize`修改文本大小

```javascript
elem.style.fontSize("20px");
```

## 删除标签

**一、删除标签**

* 先获取标签的父节点，再获取子标签，只能通过父节点删除子节点
* 使用`removeChild()`删除标签
* 可以通过`dom.parentElement`获取父标签
* 通过`dom.children`获得子标签时，删除前面的标签后面的自动会补上前面的空位

```javascript
let father = document.getElementById("father");
let arrs=father.children;
\\arrs内容发生改变，前面的元素被删除后面的会补上前面的空位
father.removeChild(arrs[0]);
```

**二、清空所有子标签**

* 先获取父类标签，通过`remove()`清空所有子类标签数据

```javascript
let father =document.getElementById("father");
father.remove();
```

## 插入标签

**一、选择标签插入**

* 使用`appendChild()`选择标签对象并插入到该父类标签纸

```javascript
let father = document.getElementById("father");
let elem = document.getElementById("other");
father.appendChild(elem);
```

```html
<html>
    <body>
        <p id="other">other</p>
        <ul id=father>
            <li>C</li>
            <li>Java</li>
            <li>Python</li>
            <li>Go</li>
        </ul>
    </body>
</html>
```

**二、创建新的标签并追加到后面**

* 使用`document.createElement("标签")`;
* 使用`dom.id`设置id属性，可以使用`setAttribute()`以键值对形式设置属性
* 使用`dom.innerText`设置文本内容
* 使用`appendChild()`方法将新创建的标签追加到父类元素里

```javascript
//获取父类文档对象
let father =document.getElementById("father");
//创建新的li标签对象
let newElem=document.createElement("li");
//给新对象设置id
newElem.id="HTML";
//设置文本内容
newElem.innerText="HTML";
//追加到父类里
father.appendChild(newElem);
```

**三、创建新的标签并插入到某个标签最前面**

* 使用`insertBefore()`方法将新创建的标签，并插入到某个标签最前面

```javascript
//获取父类文档对象
let oldElem =document.getElementById("old");
//创建新的li标签对象
let newElem=document.createElement("p");
//给新对象设置id
newElem.id="HTML";
//设置文本内容
newElem.innerText="HTML";
//追加到父类里
father.insertBefore(newElem,oldElem);
```

## 修改标签属性

**一、修改标签属性键值对**

* 使用`dom.setAttribute()`修改属性键值对（万能）

```javascript
//切换图片来源
let p=document.getElementById("photo");
p.setAttribute("src","./hoyo_emoji/刻晴-晚安.png");
```

```html
<html>
    <body>
		<img id="photo" src="./hoyo_emoji/阿波尼亚-飞扑.png">
    </body>
</html>
```

## 操作表单

### 文本框

 **一、获取文本框内容**

* 使用`dom.value`获取文本内容，也可以修改文本框内容

```javascript
let content=document.getElementById("username");
content.value;
content.value="hello world";
```

```html
<html>
    <body>
        <p>
            用户名：<input type="text" id="username">
        </p>
    </body>
</html>
```

### 单选框与多选框

**一、查看选项是否被选中**

* 使用`dom.checked`查看是否被选中，返回`true`和`false`
* 多选与单选通用

```javascript
//获取DOM对象
let man=document.getElementById("one");
let wonman=document.getElementById("two");
//查看是否被选中
wonman.checked;
man.checked
```

```html
<html>
    <body>
        <p>性别：
        <input type="radio" name="sex" value="man" id="one">男
        <input type="radio" name="sex" value="woman" id="two">女
        </p>
    </body>
</html>
```

### 表单的提交、MD5验证和钩子函数的设置

**一、点击事件**

* 使用`onClick`设置点击钩子函数（点击后触发的`JavaScript`函数）
* 需要再提交标签里设置`onClick`函数，在`JavaScript`里构建函数

```html
<html>
    <head>
        <script>
        function handler(){
            console.log("click once");
        }
        </script>
    </head>
    <body>
        <form action="#" method="post">
            <p>账号：<input type="text" id="username" name="user"></p>
            <p>密码：<input type="password" id="password" name="pwd"</p>
            <p>
                <input type="submit" onclick="handler()">
                <input type="reset">
            </p>
        </form>
    </body>
</html>
```

**二、表单提交事件**

* 在`form`标签内使用`onsubmit`属性设置提交事件的钩子函数

```html
<html>
   	<body>
        <form action="#" method="post" onsubmit="handler()">
            
        </form>
    </body>
</html>
```



**三、MD5加密**

* MD5链接：`<script src="https://cdn.bootcss.com/blueimp-md5/2.10.0/js/md5.min.js"></script>`

```html
<html>
    <head>
        <script src="https://cdn.bootcss.com/blueimp-md5/2.10.0/js/md5.min.js"></script>
        <script>
        function handler(){
            let pwd=document.getElementById("password");
            pwd.value=md5(pwd.value);
        }
        </script>
    </head>
    <body>
        <form action="#" method="post">
            <p>账号：<input type="text" id="username" name="user"></p>
            <p>密码：<input type="password" id="password" name="pwd"</p>
            <p>
                <input type="submit" onclick="handler()">
                <input type="reset">
            </p>
        </form>
    </body>
</html>
```

**四、表单的拦截与放行**

* 在`form`标签里用`onsubmit`设置钩子函数（`onsubimt="return 函数()"`）

* 在提交数据的钩子函数里设置`return true;`时，放行数据
* 在提交数据的钩子函数里设置`return false;`时，拦截数据

```html
<html>
    <head>
        <script src="https://cdn.bootcss.com/blueimp-md5/2.10.0/js/md5.min.js"></script>
        <script>
        function handler(){
            let passwd=document.getElementById("password");
            let uname=document.getElementById("username");
            if (passwd.value.length==0 || uname.value.length==0){
                return false;
            }
            passwd.value=md5(passwd.value);
            return true;
        }
        </script>
    </head>
    <body>
        <form action="#" method="post" onsubmit="return handler()">
            <p>账号：<input type="text" id="username" name="uname"/></p>
            <p>密码：<input type="password" id="password" name="passwd"/></p>
            <p>
                <input type="submit">
                <input type="reset">
            </p>
        </form>
    </body>
</html>
```

# JQuery

## JQuery相关概念

**一、JQuery**

* JQuery是JavaScript的封装库
* [JQuery库下载](https://jquery.com/download/)（production：生产环境版本，development：开发版）
* `CDN`导入`<script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>`
* JQuery公式：`$(selector).action()`（选择器是CSS的选择器）
* [jQuery API 在线手册](https://jquery.cuishifeng.cn/jquery.html)

**二、选择器的使用**

* 选择器的语法和CSS选择控件一样
* `JavaScript`要写在`body`后面，否则不能绑定标签元素

```html
<html>
    <head>
        <script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
        <script>var count=0;</script>
    </head>
    <body>
        <a href="#">点我</a>
        <a href="#",id="test">点我</a>
        <a href="#" class="demo">点我</a>
    </body>
    <script>
        //标签选择器
        $("a").click(function(){
            count++;
            alert(count);
        });
        //类选择器
        $(".demo").click(function(){
            count++;
            alert(count);
        });
        //id选择器
        $("#test").click(function(){
            count++;
            alert(count);
        });
    </script>
</html>
```

## 事件

### 鼠标事件

**一、常见鼠标事件**

| 事件         | 描述                     |
| ------------ | ------------------------ |
| mousedown()  | 鼠标按下或按住时         |
| mouseenter() | 鼠标进入某区域           |
| mouseleave() | 鼠标离开选择的某区域     |
| mouseout()   | 鼠标离开父标签内的子标签 |
| mousemove()  | 鼠标移动                 |
| mouseover()  | 位于标签元素上方时       |
| mouseup()    | 鼠标结束点击             |

**二、鼠标事件的运用**

* 设置响应函数时，可以用`this`获取选择器里的标签对象
* 可以使用`css()`函数设置css参数
* 使用`mouse.move()`方法时，需要构造含e参数的匿名函数，参数e含有鼠标的位置
* 使用`e.pageX`和`e.pageY`获取鼠标的位置

```html
<html>
    <head>
        <meta charset="UTF-8">
        <title>Title</title>
        <script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
        <style>
            .area{
                width:300px;
                height:200px;
                border:3px solid red;
            }
        </style>
    </head>
    <body>
        <div class="area" id="one">
            <ul>
                <li>C</li>
                <li>Java</li>
                <li>Python</li>
            </ul>
        </div>
        <h2>
            鼠标位置：<span id="text"></span>
        </h2>
        <div class="area" id="two">

        </div>
    </body>
    <script>
        //离开子元素时
        $("#one").mouseout(function () {
            $(this).css("background-color","skyblue");
        });
        //点击时
        $("#one").mouseup(function () {
            $(this).css("background-color","red");
        });
    </script>
    <script>
        //区域内移动时
        $("#two").mousemove(function(e){
            $("#text").text("x:"+e.pageX+" y:"+e.pageY);
        });
        //离开区域时
        $("#two").mouseleave(function () {
            $(this).css("background-color","skyblue");
        });
    </script>
</html>
```

### 网页事件

**一、网页加载完成时**

* 使用`$(document).ready()`设置网页加载完成时的响应事件
* 可以简写为`$()`，直接在括号里添加匿名函数

```html
<html>
    <head>
        <style>
            p{
                width:200px;
                height:200px;
                border:2px solid blue;
            }
        </style>
        <script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
    </head>
    <body>
        <p></p>
    </body>
    <script>
        $(document).ready(function(){
            $("p").css("background-color","skyblue");
        });
    </script>
</html>
```

### 操作DOM

**一、获取文本**

* 使用`text()`获取文本

```html
<html>
    <head>
        <script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
    </head>
    <body>
        <h1>
            hello world
        </h1>
    </body>
    <script>
    	let v=$("h1").text();
        alert(v);
    </script>
</html>
```

**二、设置文本**

* 使用`text()`修改文本

```html
<html>
    <head>
        <script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
    </head>
    <body>
        <h1>
            hello world
        </h1>
    </body>
    <script>
    	$("h1").text("demo");
    </script>
</html>
```

**三、设置或修改含HTML标签的文本**

* 使用`html()`设置或修改含html的文本

```html
<html>
    <head>
        <script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
    </head>
    <body>
    </body>
    <script>
    	$("body").html("<h1>hello world</h1>");
    </script>
</html>
```

**四、设置或修改css参数**

* 使用`css()`设置或修改css参数

```javascript
$(document).ready(function(){
    $("p").css("background-color","skyblue");
});
```

**五、标签元素的显示与隐藏**

* 使用`show()`显示标签元素
* 使用`hide()`隐藏标签元素
* 本质：将display设置为none

```javascript
$("p").hide();
$("p")show();
```

### 焦点事件

**一、失去焦点事件**

* 在标签内使用`onblur`属性设置失去焦点事件

```html
<html>
    <body>
        <input type="text" onblur="handler()">
    </body>
    <script>
    function handler(){
        alert("hello world");
    }
    </script>
</html>
```

### 输入框

**一、获取输入框输入的内容**

* 使用`val()`方法获取输入内容

```html
<html>
    <head>
        <script src="http://libs.baidu.com/jquery/2.0.0/jquery.min.js"></script>
        <script>
            function handler(){
                alert($("#input").val());
            }
        </script>        
    </head>
    <body>
        <input type="text" id="input" onblur="handler()">
    </body>
</html>
```

### 键盘事件

**一、按下**

* 使用`keydown(function(event){ })`设置键盘按下的响应事件
* 使用`event.Code`获取键盘代码
* 按下后只触发一次，按住时不触发

```javascript
$(window).keydown(function(event){
  switch(event.keyCode) {
    // ...
    // 不同的按键可以做不同的事情
    // 不同的浏览器的keycode不同
    // 更多详细信息:     https://unixpapa.com/js/key.html
    // 常用keyCode： 空格 32   Enter 13   ESC 27
  }
});
```

**二、按住**

* 使用`keydown(function(event){ }`设置按住时的响应事件
* 按住时，每输入一个字符都会响应一次

```javascript
$("input").keydown(function(){
  $("span").text(i+=1);
});
```

**三、松开**

* 使用`keyup(function(){ }`设置松开时的响应事件

```javascript
$("input").keyup(function(){
  $("input").css("background-color","#D6D6FF");
});
```

## AJAX

## 回调函数

**一、回调函数的使用**

* 如果后台返回`null`的值，则前端会调用`error`的回调函数

**二、回调函数种类**

| 回调函数       | 作用                   |
| -------------- | ---------------------- |
| ajaxComplete() | 请求完成时执行函数     |
| ajaxError()    | 请求发生错误时执行函数 |
| ajaxSend()     | 请求发送前执行函数     |
| ajaxStart()    | 请求开始时执行函数     |
| ajaxStop()     | 请求结束时执行函数     |
| ajaxSuccess()  | 请求成功时执行函数     |

## ajax请求

**一、设置ajax请求**

* 使用`$.ajax({})`或`$.get()`设置get请求
* 使用`$.post({})`设置post请求
* 需要设置的参数
  * `url`：使用`./`获取当前页面的域名，后面写路径（使用Go后端）
  * `data`：设置发送的`json`数据，需要设置键（`data:{"键":值}`）
  * `type`：设置请求类型，默认`get`（xml, html, script, json, text）
* 设置的回调函数，处理返回的数据：`success:function(data,status){ }`（可以设置其它回调函数）
* 使用`data.键`获取返回的`json`的值，返回多个`json`时使用`for`循环遍历`data`获取值

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>test</title>
        <!--后端使用Go-gin框架-->
        <script type="text/javascript" src="/static/jquery-3.6.3.js"></script>
        <script>
            function handler(){
                $.ajax({
                    url: "./getValue",
                    data: {"msg": $("#input").val()},
                    success: function (data, status) {
                        alert(data.msg+ status);
                    }
                });
            }
        </script>
    </head>
    <body>
        <p>
            <input type="text" id="input" onblur="handler()">
        </p>
    </body>
</html>
```

## Get请求

**一、`get`请求**

* 使用`$.get()`设置get请求
* 请求参数：
  * url：请求地址
  * data：发送的数据
  * success：请求成功处理的回调函数
  * type：返回内容的格式，默认`json`

**二、发送json并触发回调函数**

* 需要设置key
* 直接在url后面添加json
* 在json数据后面追加含`data`参数的匿名回调方法

```html
<html>
    <head>
        <title>test</title>
        <script type="text/javascript" src="/static/jquery-3.6.3.js"></script>
        <script>
            function handler(){
                $.get("/getValue",{"msg":$("#msg").val()},function (data) {
                    alert(data.msg);
                });
            }
        </script>
    </head>
    <body>
        <div>
            <input type="text" onblur="handler()" id="msg">
        </div>
    </body>
</html>
```

## post请求

**一、post请求**

* 输入参数：
  * url：发送请求的地址
  * data：发送的键值对
  * success：请求处理成功后的回调函数（返回非null值）
  * type：返回内容格式

```javascript
function handler(){
    $.post("/getValue",{"msg":$("#msg").val()},function (data) {
        alert(data.msg);
    });
}
```

## json请求

**一、json请求**

* 本质是数据类型设为`json`的`$.ajax()`请求，默认为get请求

* 输入参数：
  * url：请求地址
  * data：发送的数据
  * success：请求处理成功后的回调函数（返回非null值）

**二、发送json请求**

* 使用`$.getJSON()`发送`json`请求
* 可以处理josnp请求，用？代替url的参数

```javascript
$.getJSON("test.js", function(json){
    alert("JSON Data: " + json.users[3].name);
});
```

## js脚本请求

**一、js脚本请求**

* 输入参数：
  * url：`js`文件来源
  * success：请求处理成功后的回调函数（返回非null值）

* 获取js文件后会自动加载并执行js文件

```javascript
$.getScript("test.js", function(){
  alert("Script loaded and executed.");
});
```









