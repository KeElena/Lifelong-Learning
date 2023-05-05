package connectFunc

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetRpcClient(network string) (client *rpc.Client, err error) {
	client, err = rpc.Dial(network)
	return
}

func GetEthClient(network string) (client *ethclient.Client, err error) {
	client, err = ethclient.Dial(network)
	return
}
