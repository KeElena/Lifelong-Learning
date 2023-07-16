package JWT

import (
	"encoding/json"
	"testing"
)

func BenchmarkGetToken(b *testing.B) {
	//privKey
	key := []byte("123456")
	//payload
	user := `{"uid":1,"name":"demo","expireAt":1684408658}`
	payload, _ := json.Marshal(&user)
	for i := 0; i < b.N; i++ {
		GetToken(payload, key, SHA256)
	}
}

func BenchmarkCheck(b *testing.B) {
	//privKey
	key := []byte("123456")
	//payload
	user := `{"uid":1,"name":"demo","expireAt":1684408658}`
	payload, _ := json.Marshal(&user)
	token := GetToken(payload, key, SHA256)
	for i := 0; i < b.N; i++ {
		Check(token, key)
	}
}
