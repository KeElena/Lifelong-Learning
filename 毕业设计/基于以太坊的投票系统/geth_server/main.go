package main

import (
	"GethServe/accountFunc"
	"GethServe/connectFunc"
	contractFunc "GethServe/contrcatFunc"
	"GethServe/infoFunc"
	pb "GethServe/protoFile"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"google.golang.org/grpc"
	"net"
	"strings"
)

var ethClient *ethclient.Client
var rpcClient *rpc.Client
var EthNetAddr = "http://0.0.0.0:8545"

type server struct {
	pb.UnimplementedVoteServer
}

// CreateAccount 创建帐号
func (s *server) CreateAccount(ctx context.Context, in *pb.AccountPwd) (*pb.AccountInfo, error) {
	account, err := accountFunc.CreateAccount("./keystore", in.Password)
	if err != nil {
		return nil, err
	}
	return &pb.AccountInfo{AccountAddr: account.Address.String(), KeyStoreName: account.URL.String()[strings.LastIndex(account.URL.String(), "/")+1:]}, nil
}

// DeployContract 部署合约
func (s *server) DeployContract(ctx context.Context, in *pb.CreateTransaction) (*pb.ContractInfo, error) {
	contractAddr, err := contractFunc.DeployContract(ethClient, "./keystore/"+in.KeyStoreName, in.Password)
	if err != nil {
		return nil, err
	}
	return &pb.ContractInfo{ContractAddr: contractAddr, Ok: true}, nil
}

// InitContractInfo 初始化合约信息
func (s *server) InitContractInfo(ctx context.Context, in *pb.InitInfo) (*pb.Success, error) {
	pk, err := accountFunc.KeystoreToPrivateKey("./keystore/"+in.KeyStoreName, in.Password)
	if err != nil {
		return nil, err
	}
	err = contractFunc.SetExpireAndOptions(in.ContractAddr, pk, ethClient, in.ExpireTime, in.Options)
	if err != nil {
		return nil, err
	}
	return &pb.Success{Ok: true}, nil
}

// Vote 投票
func (s *server) Vote(ctx context.Context, in *pb.InitVote) (*pb.Success, error) {
	pk, err := accountFunc.KeystoreToPrivateKey("./keystore/"+in.KeyStoreName, in.Password)
	if err != nil {
		return nil, err
	}
	err = contractFunc.Vote(in.ContractAddr, pk, ethClient, uint8(in.Idx))
	if err != nil {
		return nil, err
	}
	return &pb.Success{Ok: true}, nil
}

// GetBalance 获取以太币数量
func (s *server) GetBalance(ctx context.Context, in *pb.AccountAddress) (*pb.Balance, error) {
	balanceNUm, err := infoFunc.GetBalance(rpcClient, in.AccountAddr)
	if err != nil {
		return nil, err
	}
	return &pb.Balance{Balance: balanceNUm}, nil
}

// GetContractContent 获取合约内容
func (s *server) GetContractContent(ctx context.Context, in *pb.ContractAddr) (*pb.ContractContent, error) {
	voting, options, expireTime, err := infoFunc.GetContractContent(in.ContractAddr, ethClient)
	if err != nil {
		return nil, err
	}
	return &pb.ContractContent{Voting: voting, Options: options, ExpireTime: expireTime}, nil
}

// GetVote 获取票据
func (s *server) GetVote(ctx context.Context, in *pb.AllAddress) (*pb.Prove, error) {
	prove, err := infoFunc.GetVote(in.ContactAddr, ethClient, in.AccountAddr)
	if err != nil {
		return nil, err
	}
	var data pb.Prove
	data.Idx = fmt.Sprintf("%d", prove.Idx)
	data.Voted = prove.Voted
	data.Time = fmt.Sprintf("%d", prove.Time)
	return &data, nil
}

// WithDraw 水龙头转账
func (s *server) WithDraw(ctx context.Context, in *pb.AccountAddress) (*pb.Null, error) {
	from := "0x6f0009c31f116d9ec3ed7e76436203467092cef0"
	err := contractFunc.WithDraw(ethClient, from, in.AccountAddr)
	return &pb.Null{}, err
}

func main() {
	var err error
	var listener net.Listener
	listener, err = net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		fmt.Println(err)
		panic("listener error")
	}
	defer listener.Close()

	rpcClient, err = connectFunc.GetRpcClient(EthNetAddr)
	if err != nil {
		fmt.Println(err)
		panic("rpcClient error")
	}

	ethClient, err = connectFunc.GetEthClient(EthNetAddr)
	if err != nil {
		fmt.Println(err)
		panic("ethClient error")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterVoteServer(grpcServer, &server{})
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println(err)
		panic("grpc server run error")
	}
}
