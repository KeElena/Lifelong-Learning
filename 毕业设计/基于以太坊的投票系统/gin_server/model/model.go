package model

import (
	"BridgeModule/connect"
	pb "BridgeModule/protoFile"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var (
	grpcConn  *grpc.ClientConn
	mongoConn *mongo.Client
	mysqlConn *gorm.DB
	redisConn *redis.Client
)

func init() {
	//数据库连接
	grpcConn, _ = connect.GRPC()
	mongoConn, _ = connect.Mongo()
	mysqlConn, _ = connect.MYSQL()
	redisConn, _ = connect.Redis()
}

type VoteContent struct {
	Sponsor      string   `json:"sponsor"`
	Title        string   `json:"title""`
	Text         string   `json:"text"`
	Options      []string `json:"options"`
	Duration     int      `json:"duration"`
	CreateTime   int64    `json:"createTime"`
	ContractAddr string   `json:"contractAddr"`
}

type VoteSimpleInfo struct {
	Sponsor      string `json:"sponsor"`
	Title        string `json:"title"`
	ContractAddr string `json:"contractAddr"`
}

type VoteDetailContent struct {
	Sponsor string   `json:"sponsor"`
	Title   string   `json:"title"`
	Text    string   `json:"text"`
	Options []string `json:"options"`
	Expire  string   `json:"expire"`
	Result  []uint32 `json:"result"`
}

type VoteProve struct {
	Voted bool   `json:"voted"`
	Idx   string `json:"idx"`
	Time  string `json:"time"`
}

type User struct {
	Id       int64  `json:"id"`
	Uname    string `json:"uname"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Nuid     string `json:"nuid"`
	Ethaddr  string `json:"ethaddr"`
	Keystore string `json:"keystore"`
}

//连接检查
func checkMySQL() error {
	var err error
	if mysqlConn == nil {
		mysqlConn, err = connect.MYSQL()
		if err != nil {
			return err
		}
	}
	return nil
}
func checkRPC() error {
	var err error
	if grpcConn == nil {
		grpcConn, err = connect.GRPC()
		if err != nil {
			return err
		}
	}
	return nil
}
func checkMongo() error {
	var err error
	if mongoConn == nil {
		mongoConn, err = connect.Mongo()
		if err != nil {
			return err
		}
	}
	return nil
}
func checkRedis() error {
	var err error
	if redisConn == nil {
		redisConn, err = connect.Redis()
		if err != nil {
			return err
		}
	}
	return nil
}

//--------------contract设置相关-------------------//

func CreateVote(password string, keystore string, options []string, duration uint64) (contractAddr string, err error) {
	//检查rpc连接
	err = checkRPC()
	if err != nil {
		return
	}
	//获取客户端
	client := pb.NewVoteClient(grpcConn)
	//部署合约
	handleInfo, err := client.DeployContract(context.Background(), &pb.CreateTransaction{Password: password, KeyStoreName: keystore})
	if err != nil {
		return
	}
	if !handleInfo.Ok {
		err = fmt.Errorf("contract create fail")
		return
	}
	//初始化合约信息
	contractAddr = handleInfo.ContractAddr
	go func(contractAddr string, keystore string, password string, options []string) {
		//等待创建合约交易完成
		time.Sleep(time.Second * 5)
		//初始化合约信息
		_, err = client.InitContractInfo(context.Background(), &pb.InitInfo{ContractAddr: contractAddr, KeyStoreName: keystore, Password: password, ExpireTime: uint64(time.Now().Unix()) + duration*86400, Options: options})
		if err != nil {
			fmt.Println(err)
		}
	}(contractAddr, keystore, password, options)
	return
}

// Vote 投票
func Vote(password string, keystore string, contractAddr string, choose uint32) (err error) {
	//检查rpc连接
	err = checkRPC()
	if err != nil {
		return
	}
	//获取客户端
	client := pb.NewVoteClient(grpcConn)
	//投票
	info, err := client.Vote(context.Background(), &pb.InitVote{ContractAddr: contractAddr, Password: password, KeyStoreName: keystore, Idx: choose})
	if err != nil {
		return
	}
	if !info.Ok {
		return fmt.Errorf("vote handler fail")
	}
	return
}

//获取以太币
func GetETH(to string) (err error) {
	//检查连接
	err = checkRPC()
	if err != nil {
		return
	}
	//获取客户端连接
	client := pb.NewVoteClient(grpcConn)
	//发送交易
	_, err = client.WithDraw(context.Background(), &pb.AccountAddress{AccountAddr: to})
	return
}

//------------------contract取值相关-----------------------//

// GetBalance 获取以太币数量
func GetBalance(accountAddr string) (balance int64, err error) {
	err = checkRPC()
	if err != nil {
		return
	}
	client := pb.NewVoteClient(grpcConn)
	b, err := client.GetBalance(context.Background(), &pb.AccountAddress{AccountAddr: accountAddr})
	//去掉0x
	balance, _ = strconv.ParseInt(b.Balance[2:], 16, 64)
	return
}

// GetDetailContent 获取投票详细内容
func GetDetailContent(contractAddr string) (content *VoteDetailContent, err error) {
	//检查连接
	err = checkRPC()
	if err != nil {
		return
	}
	err = checkMongo()
	if err != nil {
		return
	}
	var blockInfo *pb.ContractContent
	//区块链获取信息
	blockInfo, err = pb.NewVoteClient(grpcConn).GetContractContent(context.Background(), &pb.ContractAddr{ContractAddr: contractAddr})
	if err != nil {
		return
	}
	//mongodb获取信息
	client := mongoConn.Database("vote").Collection("content")
	where := bson.M{"contractaddr": contractAddr}
	var vote VoteDetailContent
	_ = client.FindOne(context.Background(), where).Decode(&vote)
	//设置区块数据
	vote.Options = blockInfo.Options
	vote.Result = blockInfo.Voting
	vote.Expire = fmt.Sprintf("%d000", blockInfo.ExpireTime)
	return &vote, nil
}

func GetVoteProve(userAddr string, contractAddr string) (prove *VoteProve, err error) {
	var data *pb.Prove
	data, err = pb.NewVoteClient(grpcConn).GetVote(context.Background(), &pb.AllAddress{ContactAddr: contractAddr, AccountAddr: userAddr})
	if err != nil {
		return
	}
	return &VoteProve{data.Voted, data.Idx, data.Time + "000"}, nil
}

//------------------set方法-----------------------//

func Register(user *User) (err error) {
	//检查mysql连接
	err = checkMySQL()
	if err != nil {
		return
	}
	//检查grpc连接
	err = checkRPC()
	if err != nil {
		return
	}
	//插入数据并获取id
	err = mysqlConn.Create(&user).Error
	if err != nil {
		return
	}
	//创建以太坊帐号
	go func(password string, id int64, client pb.VoteClient, mysqlCoon *gorm.DB) {
		ethAccount, err := client.CreateAccount(context.Background(), &pb.AccountPwd{Password: user.Password})
		if err != nil {
			return
		}
		mysqlCoon.Where("id=?", id).Updates(&User{Ethaddr: ethAccount.AccountAddr, Keystore: ethAccount.KeyStoreName})
	}(user.Password, user.Id, pb.NewVoteClient(grpcConn), mysqlConn)
	return
}

func SetVoteContent(content *VoteContent) (err error) {
	//检查连接
	err = checkMongo()
	if err != nil {
		return
	}
	err = checkRedis()
	if err != nil {
		return
	}
	//获取处理时间
	content.CreateTime = time.Now().UnixMilli()
	//获取mongo客户端
	client := mongoConn.Database("vote").Collection("content")
	//插入数据到mongo
	res, err := client.InsertOne(context.Background(), content)
	if err != nil {
		return
	}
	//放进缓存
	err = redisConn.Set(res.InsertedID.(primitive.ObjectID).Hex(), "", time.Second*time.Duration(86400*content.Duration-180)).Err()
	if err != nil {
		return
	}
	return
}

type History struct {
	Uid          int64  `json:"uid"`
	Title        string `json:"title"`
	VoteTime     string `json:"vote_time"`
	ContractAddr string `json:"contract_addr"`
}

// SetUserHistory
func SetUserHistory(uid int64, ethAddr string, title string) (err error) {
	err = checkMongo()
	if err != nil {
		return
	}
	client := mongoConn.Database("vote").Collection("history")
	t := time.Now().Format("2006-06-02 15:01")
	data := &History{Uid: uid, Title: title, VoteTime: t, ContractAddr: ethAddr}
	_, err = client.InsertOne(context.Background(), data)
	return
}

//------------------------get方法---------------------//

// GetDataByPhone GetPhone 使用手机号获取数据
func GetDataByPhone(phone string) (*User, error) {
	var data User
	var err error
	//检查mysql连接
	err = checkMySQL()
	if err != nil {
		return nil, err
	}
	err = mysqlConn.Where("phone=?", phone).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, err
}

// GetKeystoreFromMysql mysql获取keystore
func GetKeystoreFromMysql(id int) (keystore string, err error) {
	err = checkMySQL()
	if err != nil {
		return "", err
	}
	var user User
	//获取数据
	err = mysqlConn.Where("id=?", id).Pluck("keystore", &user).Error
	if err != nil {
		return
	}
	keystore = user.Keystore
	return
}

// GetVoteList 获取投票列表
func GetVoteList() (res []*VoteSimpleInfo, err error) {
	err = checkRedis()
	if err != nil {
		return nil, err
	}
	err = checkMongo()
	if err != nil {
		return nil, err
	}
	var list []string
	res = make([]*VoteSimpleInfo, 0, 8)
	//获取投票id
	list, err = redisConn.Keys("*").Result()
	if err != nil {
		return
	}
	//mongo中获取信息
	client := mongoConn.Database("vote").Collection("content")
	length := len(list)
	if length > 8 {
		length = 8
	}
	for _, id := range list[:length] {
		oid, _ := primitive.ObjectIDFromHex(id)
		where := bson.M{"_id": oid}
		var vote = VoteSimpleInfo{}
		_ = client.FindOne(context.Background(), where).Decode(&vote)
		res = append(res, &vote)
	}
	if length < 8 {
		for i := 0; i < 8-length; i++ {
			res = append(res, &VoteSimpleInfo{})
		}
	}
	return
}

func GetHistory(uid int64) (list []*History, err error) {
	//检查连接
	err = checkMongo()
	if err != nil {
		return nil, err
	}
	//获取客户端对象
	client := mongoConn.Database("vote").Collection("history")
	//查询条件
	where := bson.M{"uid": uid}
	//设置选项
	option := &options.FindOptions{}
	option.SetLimit(3)
	option.SetSort(bson.M{"_id": -1})
	//查询数据
	cur, err := client.Find(context.Background(), where, option)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.Background()) {
		var d History
		_ = cur.Decode(&d)
		list = append(list, &d)
	}
	if len(list) < 3 {
		for i := len(list); i < 3; i++ {
			list = append(list, &History{})
		}
	}
	return list, nil
}
