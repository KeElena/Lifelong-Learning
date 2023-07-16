# web3模块下载

```bash
npm install web3
```

# 合约的部署与获取接口

**一、合约的bin文件**

* 是合约编译后的字节码，用于部署到以太坊区块链上

**二、合约的abi文件**

* 合约的web3应用接口，由web3.js调用

# web3js的基本使用

## 创建web3对象

**一、创建web3对象**

* 使用`new Web3()`创建web3对象
* 如果页面存在`web3`则使用当前页面的`web3`对象，否则新创建一个`web3`对象
* 如果浏览器有`metamask`插件，浏览器内就会有一个`web3`对象

```javascript
//node环境导入库
var Web3=require("web3");
//获取web3对象
var web3 = new Web3(new Web3.providers.HttpProvider("http://127.0.0.1:8545"));
```

```javascript
//如果存在web3
if(typeof web3!='undefined'){
    //则使用当前web3对象，注意http服务器地址
    web3=new Web3(web3.currentProvider);
}else{
    //新创建web3
    web3=new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
}
```

## 大数处理

**一、大数的处理**

* `JavaScript`的数字精度小，需要使用`web3.js`的依赖库`BigNumber`进行处理
* 内部处理时转换为`wei`，显示给用户时转换为`ether`

```javascript
var BigNumber=require("bignumber.js");
var num=new BigNumber('123456789123456789123456789123456789123456789');
```

## 查询以太坊网络状态

```javascript
web3.eth.net.isListening().then(console.log)
```

## Provider

**一、查看当前设置的web3 provider**

```javascript
web3.currentProvider
```

**二、、查看浏览器环境设置的web3 provider**

```javascript
web3.givenProvider
```

**三、设置provider**

```javascript
web3.setProvider(new Web3.providers.HttpProvider("http://localhost:8545"))
```

















