# TypeScript介绍

**一、什么是TypeScript**

* TypeScript是JavaScript的超集
* 在JS的基础上添加了类型支持

**二、JS的缺陷与TS的改进**

* JS的绝大数错误是类型错误，增加了找bug的时间，严重影响开发效率，TS里增加了类型支持
* 代码执行时才能发现错误，TS在编译时就能发现错误

**三、TS的优势**

* **所有合法的JS代码都是TS代码**
* 更早发现bug，减少找bug时间，提高了开发效率
* 代码提示
* 类型系统提升了代码可维护性
* 支持ECMAScript语法
* TS类型推断机制，不需要每个地方标注类型
* 目前主流前端语言首选

# 安装TS

**一、安装nodejs**

* [node.js官方安装地址](https://nodejs.cn/download/)

**二、设置国内镜像**

```shell
npm config set registry https://registry.npmmirror.com
```

**三、安装TS**

```shell
npm install -g typescript
```

**四、验证安装**

```shell
tsc -v
```

# TS代码的hello world

**一、运行简单的TS代码**

* 编写TS代码

```typescript
console.log("hello world")
```

* 编译TS代码得到JS代码

```shell
tsc hello.ts
```

* 运行JS代码

```shell
node hello.js
```

**二、简化运行TS**

* 安装`ts-node`包

```shell
npm i -g ts-node
```

* 运行ts代码

```typescript
ts-node hello.ts
```

# TypeScript数据类型

## 基础数据类型

**一、JS的数据类型**

* JS的数据类型能在TS使用

| 类型      | 描述                           |
| --------- | ------------------------------ |
| number    | 数字类型                       |
| string    | 字符串类型                     |
| boolean   | 布尔类型                       |
| null      | 空类型                         |
| undefined | 未定义类型                     |
| symbol    | 构造类型                       |
| object    | 对象（包含数组、对象和函数等） |

**二、TS新增基础类型**

* 联合类型、自定义类型（类型别名）、接口、元组、字面量类型、枚举、void、any

## 定义变量

**一、定义局部变量**

* 格式：`let 变量名:类型=值`

```typescript
let age : number=18
let name:string="demo"
let ok:boolean=false
let s:symbol=Symbol()
```

**二、定义常量**

* 格式：`const 变量名:类型=值`
* 常量值不可修改

```typescript
const str:string="hello"
```

## 数组类型

**一、数组类型**

* 定义数字数组语法：`let 数组名: 值类型[]=[v1,v2,v3]`

```typescript
let arr :number[]=[1,2,3]
```

**二、联合类型数组**

* 使用联合类型实现数组有多种类型数据
* 格式：`let 数组名:(类型1 | 类型2)[]=[v1,v2,v3]`

```typescript
let arr:(number | string)[]=[1,"a",2,"b"]
```

## 类型别名

**一、类型别名**

* 用于自定义类型
* 使用场景：类型复杂，且被多次使用时，可以通过类型别名简化该类型的使用

* 使用`type`定义类型别名
* 格式：`type 类型别名=类型`

```typescript
type myType=(number | string)[]
```

## 函数类型

**一、定义函数**

* 格式一

```typescript
function 函数名(参数1:类型1,参数2:类型2) :返回值类型{
    return v
}
```

* 格式二

```typescript
let add=(参数1:类型1,参数2:类型2):返回值类型=>{
    return v
}
```

**二、定义函数类型**

* 函数类型格式：`函数名:(输入参数)=>输出参数`

```typescript
let add:(n1:number,n2:number)=>number = (n1,n2)=>{
    return n1+n2
}
```

## void类型

**一、void类型**

* 函数没有返回值时，则函数返回类型为void

```typescript
function f(name: string):void{
    console.log("hello ",name)
}
```

## 函数类型

**一、可选参数**

* 函数实现某个功能时参数可以传也可以不传，在给函数参数指定类型时可以用到可选参数
* 可选参数只能在必选参数的后面
* 例子：数组的`slice()`可以只输入起始值，也可以输入起始值和结束值

```typescript
function f(start :number,end?:number):void{
    console.log("start:",start,"end:",end)
}
```

## 对象类型

**一、对象类型**

* 使用`{}`描述对象结果，属性采用`属性名:类型`的形式定义
* 定义类型时使用`;`分隔不同的参数，换行时不用符号作结尾
* 对象里可以定义方法类型
* 格式：`let 对象名:类型集={实例数据}`

```typescript
//定义对象
let person:{
    name:string
    age:number
    //函数写法1
    sayHello(name:string):void
    //函数写法二
    sayHi(name:string)=>void
}={
    name:"jack",
    age:18,
    sayHello(){
        console.log("hello")
    }
}
```

**二、对象的可选属性**

* 必选放前，可选放后

```typescript
function myAxios(config:{url:string;method?:string})
```

**三、对象的结构化**

* 使用`type`定义结构体

```typescript
type Animal={
    name:string
    num:number
}
let dog:Animal={
    name:"dog",
    num:1
}
```

## 接口

**一、接口**

* 当对象类型多次使用时，一般使用接口来描述对象类型
* 只定义类型，不定义具体实例
* 使用`interface`定义接口

```typescript
//接口的定义
interface IPerson{
    name: string
    age:number
    sayHi():void
}
//接口的使用
let person:IPerson={
    name:"jack",
    age:19,
    sayHi(){}
}
```

**二、接口的继承**

* 使用`extends`实现接口的继承
* 格式：`interfave 新接口 extends 父接口 {新属性}`

```typescript
interface Point2D {x:number;y:number}
//继承
interface Point3D extends Point2D{z:number}
```

**三、接口和类型别名的区别**

* 相同点：都可以给对象指定类型
* 不同点：
  * 接口，只能为对象指定类型，能继承，代码重用率高
  * 类型别名，可以为任何类型指定别名

```typescript
//接口
interface IPerson{
    name:string
    age:number
    sayHi():void
}
//类型别名
type IPerson={
    name:string
    age:number
    sayHi():void
}
type myArr=number|string
```

## 元组

**一、元组**

* 元组在数组的基础上限制了数组的长度和每个元素的类型

* 元组定义格式：`let 变量名:[类型1,类型2,...]=[...]`

```typescript
let position:[number,number,string]=[39,39,"demo"]
```

## 类型推断

**一、类型推断**

* TS可以通过变量的初始值判断变量的类型

```typescript
let age :number=18
//类型推断
let age=18

function add(n1:number,n2 :number):number{
    return n1+n2
}
//类型推断
function(n1:number,n2:number){
    return n1+n2
}
```

## 类型断言

**一、类型断言**

* `any`：表示任意类型，与golang的`interface{}`类型相同
* 对于`any`类型需要使用类型断言才能使用到其对象的方法
* 使用`as`实现类型断言
* 格式：`any对象 as 具体类型`

```typescript
let aLink=document.getElementById("link") as HTMLAnchorElement
console.log(aLink.href)
```

## 枚举变量

**一、枚举变量**

* 枚举变量用于表示一组可选常量
* 枚举变量的值默认为数字索引，可以自定义赋值字符串
* 定义枚举变量格式：`enum 枚举变量名`
* 枚举变量的使用：`枚举变量名.枚举变量里的常量`
* 使用`.`可以访问枚举变量内部的常量

```typescript
//定义枚举变量
enum Direction{
    Up,
    Down,
    Left,
    Right
}
//枚举变量的使用
function changeDirection(dir:Direction){
    console.log(dir,dir==Direction.Down)
}

changeDirection(Direction.Down)
```

**二、自定义枚举变量索引**

* 枚举变量的常量索引默认为自增（即索引为前面变量的自增+1）
* 在常量后面使用`=`赋值索引值
* 前面常量使用自定义索引，后面的常量会依据该自定义索引自增

```typescript
//定义枚举变量
enum Direction{
    Up,
    Down=1024,
    Left,
    Right
}
//枚举变量的使用
function changeDirection(){
    //输出0 1024 1025
    console.log(Direction.Up,Direction.Down,Direction.Left)
}

changeDirection()
```

**三、字符串枚举**

* 使用`=`可以为常量赋值字符串

```typescript
//定义字符串枚举变量
enum Direction{
    Up="Up",
    Down="Down",
    Left="Left",
    Right="Right"
}

function changeDirection(){
    console.log(Direction.Up,Direction.Down,Direction.Left)
}
changeDirection()
```

## Any类型

**一、any类型**

* 当值的类型为any时，可以对该值进行任意操作，不会由有代码提示
* 使用any类型会导致类型失去TS类型的保护优势

## typeof()函数和typeof操作符

**一、typeof()函数**

* `typeof()`用于获取数据类型

```typescript
console.log(typeof("abc"))
```

**二、typeof操作符**

* TS提供的typeof操作符可以在类型上下文中引用变量或类型的属性
* 为了简化类型书写，可以使用typeof获取对象

```typescript
//上文
let p={x:1,y:2}
//下文定义函数
function formatPoint(point: typeog p){}
//使用函数
formatPoint(p)
```

# TypeScript高级类型

## class类

**一、class类**

* 用于实现面向对象编程
* 使用`class`定义类
* 定义属性有两种方式：
  * 无默认值的属性定义
  * 有默认值的属性定义

```typescript
class Person{
    //属性定义
    age:number
    //带默认值的属性定义
    name="man"
}
```

**二、类的构造方法**

* 用于给实例赋值初始值
* **TS里无默认值的属性必须拥有一个初始值，如果没有则会报错**
* 在`class`区块里使用`constructor()`定义构造方法
* 使用`this`访问类内部的属性

```typescript
class Person{
    name:string
    age:number
    //构造方法
    constructor(name:string,age:number){
        this.name=name
        this.age=age
    }
}

let p=new Person("hello",18)
console.log(p.name)
```

**三、类的方法**

* 在`class`区块内定义方法
* 使用`this`访问类内部的属性

```typescript
class Point{
    x:number
    y:number
    constructor(x:number,y:number){
        this.x=x
        this.y=y
    }
    //定义方法
    computeDistance(p:number,q:number):number{
        return Math.sqrt(Math.pow((this.x-p),2)+Math.pow((this.y-q),2))
    }
}

let p=new Point(3,3)
console.log(p.computeDistance(9,9))
```

## class的继承

**一、继承类**

* 使用`class`关键字继承父类属性和方法
* 使用`super()`调用父类的构造方法，并初始化父类属性
* 与JAVA一样

```typescript
class Animal{
    name:string
    move(){ console.log(this.name,"Moving") }
    constructor(name:string){
        this.name=name
    }
}
class Dog extends Animal{
    age:number
    eat(){ console.log(this.age,this.name,"Eating") }
    constructor(name:string,age:number){
        super(name)
        this.age=age
    }
}

let pet=new Dog("dog",2)
pet.move()
pet.eat()
```

## interface接口

**一、interface**

* 定义类的抽象，要求类必须包含某些方法并实现
* 使用`implements`实现接口
* 必须要明确标出返回类型
* 与JAVA一样

```typescript
//定义接口
interface Animal{
    Eat():void		//无返回必须设置:void
}
//实现接口得到类
class Dog implements Animal{
    food:string
    constructor(food:string){
        this.food=food
    }
    //实现方法
    Eat(){
        console.log("Dog eat",this.food)
    }
}
let dog=new Dog("meat")
dog.Eat()
```

## 类的可见性和属性可读性

**一、class的可见性**

* 可见性：
  * `public`：公有的，所有成员可访问，默认的可见性（可以在main里使用）
  * `protected`：受保护的，对所在的类和子类可见（只能在类和子类里使用）
  * `private`：私有的，只对所在类可见（只能在类里使用）

**二、属性的只读性**

* 使用`readonly`设置属性只读补可写
* 只读的属性除了直接设置初始值外，只能在构造方法设置值

```typescript
class Animal{
    readonly name:string
    constructor(name:string){
        this.name=name
    }
}

let dog=new Animal("dog
```

## 交叉类型

**一、交叉类型接口**

* 使用`&`将多个接口或对象组合起来得到新的接口

```typescript
//接口
interface Phone{
    phoneNum:number
}
interface Address{
    address:string
}
type Person=Address&Phone
class p implements Person{
    phoneNum: number;
    address: string;
    constructor(phoneNum:number,address:string){
        this.phoneNum=phoneNum
        this.address=address
    }
}

let Mary=new p(123,"abc")
console.log(Mary.phoneNum,Mary.address)
```

```typescript
//对象
type Animal={
    name:string
}
type Action={
    action:string
}
type Dog=Animal&Action
let dog:Dog={
    name:"aa",
    action:"123"
}
console.log(dog.name,dog.action)
```

## 泛型

**一、泛型**

* 保证类型安全的前提下，让函数与多种类型一起工作实现复用
* 泛型函数的定义：`function 函数名<Type>(value:Type):Type{return value}`
* `Type`：类型占用符，具体类型由用户调用该函数时决定

```typescript
function id<Type>(value:Type):Type{
    return value
}
//调用时限定类型
let num=id<number>(10)
let str=id<string>("a")
//简化调用
let num=id(10)
```

## map类型

**一、定义map类型**

* 格式：`[key:类型]:类型`
* `key`为占位符，用于指定key

```typescript
type myMAP ={
    [key:string]:number
}
let map:myMAP={
    "a":1,
    "b":2
}
console.log(map["a"])
```

**二、map的遍历**

* 使用`for-in`循环遍历

```typescript
type myMAP ={
    [key:string]:number
}
let map:myMAP={
    "a":1,
    "b":2
}
for (let key in map) {
    //判断指定的key是否有值，如果没有则返回undefined
    if (Object.prototype.hasOwnProperty.call(map, key)) {
        console.log(key,map[key])
    }
}
//不判断是否有值
for(let key in map){
    console.log(key,map[key])
}
```











































