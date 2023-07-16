package infoFunc

import (
	"GethServe/contractAPI"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetBalance(rpcClient *rpc.Client, userAddress string) (res string, err error) {
	err = rpcClient.Call(&res, "eth_getBalance", userAddress, "latest")
	return
}

//GetContractContent 获取合约信息
func GetContractContent(contractAddr string, ethClient bind.ContractBackend) (voting []uint32, options []string, expireTime uint64, err error) {
	var voteObj *contractAPI.VoteSystem
	//获取连接合约的api对象
	voteObj, err = contractAPI.NewVoteSystem(common.HexToAddress(contractAddr), ethClient)
	if err != nil {
		return
	}
	options, voting, expireTime, err = voteObj.GetVotingInfo(&bind.CallOpts{Pending: true})
	return
}

func GetVote(contractAddr string, ethClient bind.ContractBackend, userAddr string) (vote contractAPI.VoteSystemVote, err error) {
	var voteObj *contractAPI.VoteSystem
	//获取连接合约的api对象
	voteObj, err = contractAPI.NewVoteSystem(common.HexToAddress(contractAddr), ethClient)
	if err != nil {
		return
	}
	vote, err = voteObj.GetVote(&bind.CallOpts{Pending: true, From: common.HexToAddress(userAddr)})
	return
}
