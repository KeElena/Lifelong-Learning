package accountFunc

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
)

// CreateAccount 创建账户
func CreateAccount(keyDir string, password string) (account accounts.Account, err error) {
	ks := keystore.NewKeyStore(keyDir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err = ks.NewAccount(password)
	return
}

// KeystoreToPrivateKey 获取私钥
func KeystoreToPrivateKey(privateKeyFile, password string) (privKey *keystore.Key, err error) {
	keyJson, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		return nil, err
	}
	privKey, err = keystore.DecryptKey(keyJson, password)
	if err != nil {
		return nil, err
	}
	return
}