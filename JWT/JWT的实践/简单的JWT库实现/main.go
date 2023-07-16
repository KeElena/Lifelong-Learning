package main

import (
	"encoding/json"
	"fmt"
	"jwt/JWT"
	"log"
	"time"
)

type User struct {
	Uid      int    `json:"uid"`
	Name     string `json:"name"`
	ExpireAt int64  `json:"expireAt"`
}

func main() {
	//privKey
	key := []byte("123456")
	//payload
	user := User{Uid: 1, Name: "demo", ExpireAt: time.Now().Add(time.Second * 86400).Unix()}
	payload, _ := json.Marshal(&user)
	//Get Token
	jwt := JWT.GetToken(payload, key, JWT.SHA256)
	// print Token
	fmt.Println(jwt)
	//check Token
	data, check, err := JWT.Check(jwt, key)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data), check)
	//error Token
	data, check, err = JWT.Check(jwt+"1", key)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data), check)
}
