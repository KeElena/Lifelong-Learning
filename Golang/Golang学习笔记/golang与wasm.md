# hello world

**一、hello world代码**

```go
import "syscall/js"

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("Hello World!")
}
```

**二、IDE飘红问题**

* 修改编辑器约束
  * `OS`:`js`
  * `Arch`:`wasm`
* Go版本1.21.4及以上

**三、构建wasm**

* 创建`static`目录并构建wasm文件

```shell
GOOS=js GOARCH=wasm go build -o static/main.wasm
```

* 复制JavaScript运行wasm所需的文件到static文件夹

```shell
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" static
```

**四、运行**

* 新建html文件，并在浏览器运行
* 导入`wasm_exec.js`文件
* 获取go对象
* 使用`WebAssembly.instantiateStreaming()`加载golang编译的wasm二进制文件，使用`fetch()`获取本地文件

```html
<html>
<script src="static/wasm_exec.js"></script>
<script>
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("static/main.wasm"), go.importObject)
        .then((result) => go.run(result.instance));
</script>

</html>
```

# window下MinGW-W64的安装

[镜像下载](https://files.1f0.de/mingw/)（需要配置bin目录到环境）

[在线下载](https://github.com/Vuniverse0/mingwInstaller/releases)

# Vue2/3使用wasm的环境配置

* 在`public`目录创建`util`目录
* 将`.wasm`文件和`wasm_exec.js`文件放入`util`目录
* 找到`index.html`，在`<head>`标签导入`wasm_exec.js`

```html
<head>
    ......
    <script type="module" src="/util/wasm_exec.js"></script>
</head>
```

# Go

## 设置调用函数

```go
func Print(this js.Value, inputs []js.Value) interface{} {
	......
	return ......
}

func main() {
	c := make(chan struct{})
    //设置函数
	js.Global().Set("functionName", js.FuncOf(Print))
	<-c
}
```

## 调用Javascript函数

* 使用`Get()`获取函数或变量，使用`Invoke()`设置参数并调用Javascript函数

```go
js.Global().Get("alert").Invoke("/temp/hello.txt")
```

## 读取Javascript传递的字节数组



## 无效操作

* wasm写文件操作

# Javascript

## 传递字节数组到wasm

```typescript
//reader对象
let reader = new FileReader();
//字节读取blob对象
reader.readAsArrayBuffer(file)
reader.onload = ()=>{
    //读取字节数据
    let bytes=new Uint8Array(reader.result)
	...
}
```

## 使用wasm内的函数

```javascript
const go =new window.Go()
//加载wasm文件，wasm文件要求放在public目录
WebAssembly.instantiateStreaming(fetch("util/main.wasm"), go.importObject).then((result) =>{
    go.run(result.instance)
    //调用wasm内部函数
    console.log(goprint("hhhh"))
})
```

## Vue3+ts加载函数

```typescript
//setup
//定义函数类型
type wasmFunc={
	print?:Function
}
//获取对象
let wasmFc:wasmFunc={}
//加载wasm文件
let go =new window.Go()
WebAssembly.instantiateStreaming(fetch("util/main.wasm"), go.importObject).then((result) =>{
    go.run(result.instance)
    //加载函数到javascript运行环境
    wasmFc.print=goprint
})
```

# Javascript文件操作

## 创建文件对象

```javascript
//文件内容
let fileContent = "Hello world";
//文件名
let fileName = "hello.txt";
//创建文件对象
let file = new File([fileContent], fileName);
```

## 读取文件内容

* `reader.readAsDataURL()`：输入`Blob地址`或文件对象

```javascript
let reader = new FileReader();
reader.readAsDataURL(file)
//加载完时的回调函数
reader.onload = ()=>{
    console.log(reader.result?.toString())  // reader.result为获取结果
}
```

## blob地址的创建与销毁

**一、创建blob地址**

* 输入文件对象或MediaSource，返回一个blob地址

```javascript
URL.createObjectURL(file)
```

**二、销毁blob地址**

* 输入blob地址并销毁，最后释放内存空间

```javascript
URL.revokeObjectURL(url)
```

## 下载文件

* 输入文件名和blob地址

```javascript
function downloadFile(name:string,url:string){
    let a = document.createElement('a');
    let event = new MouseEvent('click');
    a.download = name;
    a.href = url;
    a.dispatchEvent(event);
}
```









