# go-ethereum

## 下载安装

```bash
go get -u github.com/ethereum/go-ethereum
```

## 参考文档

[ Ethereum Development with Go](https://goethereumbook.org/zh/transfer-eth/)

# 基本信息相关

## 获取以太币数量

* 使用`rpc.Dial()`连接以太坊网络，获取客户端连接
* 使用`.Call()`使用以太坊客户端

```go
import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
    //连接以太坊网络获取客户端
	client, err := rpc.Dial("http://127.0.0.1:8545")
	defer client.Close()
	if err != nil {
		fmt.Println(err)
	}
	
	var result string
    //使用_代替.通过Call获取数据，放入接收变量，地址
	err = client.Call(&result, "eth_getBalance","0x...", "latest")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
```

## 通过keystore文件和密码获取私钥

```go
func KeystoreToPrivateKey(privateKeyFile, password string) (string, string, error) {
	keyjson, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		fmt.Println("read keyjson file failed：", err)
	}
	unlockedKey, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		return "", "", err
	}
	privKey := hex.EncodeToString(unlockedKey.PrivateKey.D.Bytes())
	addr := crypto.PubkeyToAddress(unlockedKey.PrivateKey.PublicKey)
	return privKey, addr.String(), nil
}
```





# 账号操作相关

## 创建新的以太坊账户

* 使用`keystore.NewKeyStore()`获取`keystore`对象
* `keystore.NewKeyStore()`第一个参数是`keystore`目录地址，最后两个参数是定义密钥库加密的资源密集程度的加密参数
* 使用`.NewAccount()`创建账户，需要输入密码

```go
import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)
func main() {
	ks := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	na, err := ks.NewAccount("123456")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(na)
}
```

## 转账

* 获取以太坊客户端`client`
* 使用`keystore`和密码获取私钥
* 使用`types.NewTx()`设置交易订单参数
* 使用`types.SignTx()`注册订单
* 使用`client.SendTransaction()`发送交易

```go
func main() {
	//获取以太坊客户端
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取私钥
	pk, err := KeystoreToPrivateKey("./keystore/UTC--2023-02-21T03-28-15.707247923Z--9165c7c6e5e7f21b7ef1a074c94815fb0f510298", "123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	//交易对象
	fromAddress := common.HexToAddress("0x9165c7c6e5e7f21b7ef1a074c94815fb0f510298")
    toAddress := common.HexToAddress("0xac9a87643f14a0906ecf15912ca8b4b63d6f547c")
	//设置交易参数
	ctx := context.Background()
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	value := big.NewInt(10000000) // in wei
	gasLimit := uint64(210000)
	gasPrice := big.NewInt(1000)
	//填写交易订单
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    value,
		Data:     make([]byte, 0),
	})
	//注册交易
	transaction, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(15)), pk.PrivateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
    //发送交易
	err = client.SendTransaction(context.Background(), transaction)
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(time.Second)
    //查看转账结果
	rpcConn, err := rpc.Dial("http://127.0.0.1:8545")
	defer rpcConn.Close()

	if err != nil {
		fmt.Println(err)
	}
	res, _ := GetBalance(rpcConn, "0xac9a87643f14a0906ecf15912ca8b4b63d6f547c")
    fmt.Println(strconv.ParseInt(res[2:], 16, 64))

}
func GetBalance(rpcClient *rpc.Client, userAddress string) (res string, err error) {
	err = rpcClient.Call(&res, "eth_getBalance", userAddress, "latest")
	return
}
```

# 部署合约

**一、准备工作**

* 在[remix开发界面](http://remix.ethereum.org/)编写solidity脚本并编译
* 找到编译界面，查看编译详情，将`WEB3DEPLOY`的`data`数据复制下来为`bin`文件
* 在编译界面最下面找到`ABI`按钮，复制ABI文件的内容并保存

**二、定义BIN和ABI全局变量**

* 使用`const`定义全局变量

```go
const BIN=""
const ABI=""
```

**三、使用账户文件和密码获取私钥**

* 需要`keystore`文件和密码
* 使用`ioutil.ReadFile()`读取`keystore`文件
* 使用`keystore.DecryptKey()`方法获取私钥

```go
func KeystoreToPrivateKey(privateKeyFile, password string) (privKey *keystore.Key, err error) {
	keyJson, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		return nil, err
	}
	privKey, err = keystore.DecryptKey(keyJson, password)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}
```

**四、获取以太坊客户端连接**

* 使用`ethclient.Dial()`获取以太坊客户端连接
* `url`必须有`http`协议和端口

```go
client, err := ethclient.Dial("http://127.0.0.1:8545")
defer client.Close()
if err != nil {
    fmt.Println(err)
}
```

**五、获取交易对象并设置订单信息**

* 使用`bind.NewKeyedTransactorWithChainID()`获取交易对象
* `GasLimit`对于大合约建议设置到`200w`以上
* <font color=red>GasLimit过低会导致合约部署失败</font>
* 需要挖矿，否则合约不能上链

```go
//获取上下文
ctx := context.Background()
//nonce递增随机数
nonce, err := client.PendingNonceAt(ctx, pk.Address)
if err != nil {
    fmt.Println(err)
    return
}
//获取预估消耗的gasPrice
gasPrice, err := client.SuggestGasPrice(ctx)
if err != nil {
    fmt.Println(err)
    return
}
//获取交易对象
auth, err := bind.NewKeyedTransactorWithChainID(pk.PrivateKey, big.NewInt(15))
//设置部署合约的配置
auth.Nonce = big.NewInt(int64(nonce))
auth.Value = big.NewInt(0)     // in wei
auth.GasLimit = uint64(300000) // in units
auth.GasPrice = gasPrice
```

**六、发起交易部署合约**

* 使用`abi.JSON()`解析`ABI`内容获取`ABI`对象
* 使用`bind.DeployContract()`部署合约（需要参数：交易订单、ABI对象、BIN对象、客户端对象）

```go
func deploy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *bind.BoundContract, *types.Transaction, error) {
	//解析abi
	Abi, err := abi.JSON(strings.NewReader(PROJECTABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
    //部署合约
	address, tx, contract, err := bind.DeployContract(auth, Abi, common.FromHex(BIN), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, contract, tx, nil
}
```

**七、项目实践**

```go
import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
	"strings"
)

const BIN = `0x608060405234801561001057600080fd5b5060e28061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063c2bc2efc14602d575b600080fd5b606b6038366004607e565b60408051602080820183526000918290526001600160a01b03939093168152808352819020815192830190915254815290565b6040519051815260200160405180910390f35b600060208284031215608f57600080fd5b81356001600160a01b038116811460a557600080fd5b939250505056fea264697066735822122057fb326f966f039b2013261419cdc97cec8a1b27f0e1d3979c6071b1dcdaad7864736f6c63430008110033`
const PROJECTABI = `[{"inputs":[{"internalType":"address","name":"ad","type":"address"}],"name":"get","outputs":[{"components":[{"internalType":"uint256","name":"num","type":"uint256"}],"internalType":"struct demo.v","name":"","type":"tuple"}],"stateMutability":"view","type":"function"}]`
//部署合约
func deploy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *bind.BoundContract, *types.Transaction, error) {
	//解析abi
	Abi, err := abi.JSON(strings.NewReader(PROJECTABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, Abi, common.FromHex(BIN), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, contract, tx, nil
}
//获取私钥
func KeystoreToPrivateKey(privateKeyFile, password string) (privKey *keystore.Key, err error) {
	keyJson, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		return nil, err
	}
	privKey, err = keystore.DecryptKey(keyJson, password)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}

func main() {
    //获取以太坊客户端连接
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	defer client.Close()
	if err != nil {
		fmt.Println(err)
	}
	//获取私钥
	pk, err := KeystoreToPrivateKey("./keystore/UTC--2023-02-20T00-54-44.966527958Z--ac9a87643f14a0906ecf15912ca8b4b63d6f547c", "123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pk)
	//获取上下文
	ctx := context.Background()
	nonce, err := client.PendingNonceAt(ctx, pk.Address)
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取gasPrice建议
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取交易对象
	auth, err := bind.NewKeyedTransactorWithChainID(pk.PrivateKey, big.NewInt(15))
	//设置部署合约的配置
    auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
    //部署合约，得到合约地址，合约对象，交易信息
	contractAddr, contract, tx, err := deploy(auth, client)
	if err != nil {
		fmt.Println(err, 4)
		return
	}
	fmt.Println(contractAddr, contract, tx)
}
```

# 调用合约

**一、获取abi接口文件**

* [官方github下载源码](https://github.com/ethereum/go-ethereum)
* 进入`~/go-ethereum/cmd/abigen`目录
* 对里面的文件进行编译得到`abigen`二进制文件
* 通过bin文件和abi文件获取合约接口

```bash
./abigen --abi abi文件 --pkg 接口包名 --type 合约名 --out api文件.go
```

**二、字符串地址转为common类型地址**

* 使用`common.HexToAddress()`实现

```go
useAddr := "0xAC9a87643F14A0906eCf15912CA8b4B63D6f547C"
addr:=common.HexToAddress(useAddr)
```

**三、获取合约对象**

* 导入api文件
* 使用`New~()`方法获取合约对象，需要合约地址和客户端对象

```go
contractObj, err := api.NewGetDataFunc(contractAddr, client)
if err != nil {
    fmt.Println(err)
    return
}
```

**四、调用不修改合约状态的方法**

* 直接`.方法名()`获取值
* 需要输入`&bind.CallOpts{}`调用参数
* 需要用到地址时放入地址

```go
//获取数字
res1, err := contractObj.Num(&bind.CallOpts{Pending: true})
if err != nil {
    fmt.Println(err, 0)
    return
}
fmt.Println(res1)
//获取调用者地址
res2, err := contractObj.Myaddr(&bind.CallOpts{Pending: true, From: userAddr})
if err != nil {
    fmt.Println(err, 0)
    return
}
fmt.Println(res2)
```

**五、调用修改合约状态的方法**

* 通过发起交易修改状态，需要花费以太币
* 需要解锁账户获取私钥进行操作
* 使用` bind.NewKeyedTransactorWithChainID()`发起交易
* 需要挖矿，否则不能修改
* 好像不能获取返回值

```go
tx, err := bind.NewKeyedTransactorWithChainID(pk.PrivateKey, big.NewInt(15))
res3, err := contractObj.Add(tx)
if err != nil {
    fmt.Println(err, 2)
    return
}
fmt.Println(res3.Data())
```

**六、项目实践**

```go
import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-eth/api"
	"io/ioutil"
	"math/big"
	"strings"
)

const BIN = `0x60806040526000805534801561001457600080fd5b5060e2806100236000396000f3fe6080604052348015600f57600080fd5b5060043610603c5760003560e01c80634e70b1dc1460415780634f2be91f14605c578063f0759a4d146062575b600080fd5b604960005481565b6040519081526020015b60405180910390f35b6049606f565b6040513381526020016053565b60008054607c9060016086565b6000819055919050565b8082018082111560a657634e487b7160e01b600052601160045260246000fd5b9291505056fea26469706673582212200dd5c6e6ccf1614e5f9c135012ff9d0b54d9c8b9e5a49d3dd0b65fb03d57cbf664736f6c63430008110033`
const PROJECTABI = `[{"inputs":[],"name":"add","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"myaddr","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"num","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`

func deploy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *bind.BoundContract, *types.Transaction, error) {
	//解析abi
	Abi, err := abi.JSON(strings.NewReader(PROJECTABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, Abi, common.FromHex(BIN), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, contract, tx, nil
}

func KeystoreToPrivateKey(privateKeyFile, password string) (privKey *keystore.Key, err error) {
	keyJson, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		return nil, err
	}
	privKey, err = keystore.DecryptKey(keyJson, password)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}

func useApi(userAddr common.Address, client bind.ContractBackend, contractAddr common.Address, pk *keystore.Key) {
	//contractObj, err := api.NewTeststruct(contractAddr, client) 另一个合约
	contractObj, err := api.NewGetDataFunc(contractAddr, client)
	if err != nil {
		fmt.Println(err)
		return
	}
	res1, err := contractObj.Num(&bind.CallOpts{Pending: true})
	if err != nil {
		fmt.Println(err, 0)
		return
	}
	fmt.Println(res1)

	res2, err := contractObj.Myaddr(&bind.CallOpts{Pending: true, From: userAddr})
	if err != nil {
		fmt.Println(err, 0)
		return
	}
	fmt.Println(res2)
	tx, err := bind.NewKeyedTransactorWithChainID(pk.PrivateKey, big.NewInt(15))
	res3, err := contractObj.Add(tx)
	if err != nil {
		fmt.Println(err, 2)
		return
	}
	fmt.Println(res3.Data())
}

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	defer client.Close()
	if err != nil {
		fmt.Println(err)
	}

	pk, err := KeystoreToPrivateKey("./keystore/UTC--2023-02-20T00-54-44.966527958Z--ac9a87643f14a0906ecf15912ca8b4b63d6f547c", "123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	//
	//ctx := context.Background()
	//nonce, err := client.PendingNonceAt(ctx, pk.Address)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//gasPrice, err := client.SuggestGasPrice(ctx)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//auth, err := bind.NewKeyedTransactorWithChainID(pk.PrivateKey, big.NewInt(15))
	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.Value = big.NewInt(0)     // in wei
	//auth.GasLimit = uint64(600000) // in units
	//auth.GasPrice = gasPrice
	//
	//contractAddr, _, _, err := deploy(auth, client)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(contractAddr)

	useAddr := "0xAC9a87643F14A0906eCf15912CA8b4B63D6f547C"
	contractAddr := "0xF7AE4AC3163bfb1Ffc394a8dc0436b206a0794e4"
	useApi(common.HexToAddress(useAddr), client, common.HexToAddress(contractAddr), pk)
}

```



# 类型转换

**一、**



[面向 Go 开发者的以太坊 | ethereum.org](https://ethereum.org/zh/developers/docs/programming-languages/golang/)

[【智能合约】Go语言调用智能合约 | geth-阿里云开发者社区 (aliyun.com)](https://developer.aliyun.com/article/822806)