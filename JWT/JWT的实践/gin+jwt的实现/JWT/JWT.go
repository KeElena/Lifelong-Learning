package JWT

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"strings"
)

type Alg int

var (
	MD5    Alg = 0
	SHA1   Alg = 1
	SHA256 Alg = 2
	SHA512 Alg = 3
)

type header struct {
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

// Base64URLEncoding Base64URL编码
func Base64URLEncoding(content []byte) string {
	return base64.RawURLEncoding.EncodeToString(content)
}

// Base64URLDecoding Base64URL解码
func Base64URLDecoding(content string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(content)
}

// HMAC HMAC消息认证
func HMAC(Content []byte, PrivKey []byte, Typ Alg) string {
	var sh hash.Hash
	switch Typ {
	case MD5:
		sh = hmac.New(md5.New, PrivKey)
	case SHA1:
		sh = hmac.New(sha1.New, PrivKey)
	case SHA256:
		sh = hmac.New(sha256.New, PrivKey)
	case SHA512:
		sh = hmac.New(sha512.New, PrivKey)
	}
	sh.Write(Content)
	return hex.EncodeToString(sh.Sum([]byte("")))
}

// GetToken 获取Token
func GetToken(Payload []byte, privKey []byte, Typ Alg) string {
	headerStruct := &header{Typ: "JWT"}
	switch Typ {
	case MD5:
		headerStruct.Alg = "HMD5"
	case SHA1:
		headerStruct.Alg = "HS1"
	case SHA256:
		headerStruct.Alg = "HS256"
	case SHA512:
		headerStruct.Alg = "HS512"
	}
	headerByte, _ := json.Marshal(&headerStruct)
	//base64URL编码
	headerStr := Base64URLEncoding(headerByte)
	payloadStr := Base64URLEncoding(Payload)
	//获取签名
	abstract := Base64URLEncoding([]byte(HMAC([]byte(strings.Join([]string{headerStr, payloadStr}, ".")), privKey, Typ)))
	//返回Token
	return strings.Join([]string{headerStr, payloadStr, abstract}, ".")
}

func Check(jwt string, privKey []byte) (payload []byte, checked bool, err error) {
	frontIdx := strings.Index(jwt, ".")
	if frontIdx == -1 {
		return nil, false, fmt.Errorf("no data")
	}
	//获取header数据
	headerByte, err := Base64URLDecoding(jwt[:frontIdx])
	if err != nil {
		return nil, false, err
	}
	var headerStruct header
	err = json.Unmarshal(headerByte, &headerStruct)
	if err != nil {
		return nil, false, err
	}
	//校验签名
	lastIdx := strings.LastIndex(jwt, ".")
	signature, err := Base64URLDecoding(jwt[lastIdx+1:])
	if err != nil {
		return nil, false, err
	}
	switch headerStruct.Alg {
	case "HMD5":
		checked = HMAC([]byte(jwt[:lastIdx]), privKey, MD5) == string(signature)
	case "HS1":
		checked = HMAC([]byte(jwt[:lastIdx]), privKey, SHA1) == string(signature)
	case "HS256":
		checked = HMAC([]byte(jwt[:lastIdx]), privKey, SHA256) == string(signature)
	case "HS512":
		checked = HMAC([]byte(jwt[:lastIdx]), privKey, SHA512) == string(signature)
	default:
		return nil, false, fmt.Errorf("no encryption algorithm matched")
	}
	//如果校验错误
	if !checked {
		return
	}
	//获取负载
	payload, err = Base64URLDecoding(jwt[frontIdx+1 : lastIdx])
	if err != nil {
		return nil, true, err
	}
	return
}
