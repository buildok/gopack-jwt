// package jwt
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	// "encoding/hex"
	"encoding/json"
	"fmt"
)

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type JWT struct {
	Key string `json:"key"`
	Hdr []byte `json:"hdr"`
}

func New(key string) *JWT {
	jwt := new(JWT)
	jwt.Key = key

	hdr := header{
		Alg: "HS256",
		Typ: "JWT",
	}
	res, err := json.Marshal(hdr)
	if err != nil {

	}
	jwt.Hdr = res

	return jwt
}

func (j *JWT) Encode(payload interface{}) string {
	seg := base64.URLEncoding.EncodeToString(j.Hdr)

	p, err := json.Marshal(payload)
	if err != nil {

	}
	seg2 := base64.URLEncoding.EncodeToString(p)

	msg := seg + "." + seg2
	sign := hmac.New(sha256.New, []byte(j.Key))
	sign.Write([]byte(msg))

	ret := msg + "." + base64.URLEncoding.EncodeToString(sign.Sum(nil))
	return ret
}

// /////////////////////////////////////////////

type Payload struct {
	Id   int
	Name string
}

func main() {
	// jwt := jwt.New("secret")
	jwt := New("secret")

	// fmt.Printf("%#v\n", jwt)
	p := Payload{
		Id:   1,
		Name: "qqq",
	}
	fmt.Printf("%s\n", jwt.Encode(p))
}
