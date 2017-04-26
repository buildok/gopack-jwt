// package jwt
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type payload struct {
}

type JWT struct {
	Key []byte
	Hdr header
}

func New(key []byte) *JWT {
	jwt := new(JWT)
	jwt.Key = key
	jwt.Hdr = header{Alg: "HS256", Typ: "JWT"}

	return jwt
}

func (j *JWT) Encode(payload interface{}) (string, error) {
	h, err := json.Marshal(j.Hdr)
	seg := base64.URLEncoding.EncodeToString(h)

	p, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	seg2 := base64.URLEncoding.EncodeToString(p)

	msg := seg + "." + seg2
	sign := hmac.New(sha256.New, j.Key)
	sign.Write([]byte(msg))

	token := msg + "." + base64.URLEncoding.EncodeToString(sign.Sum(nil))

	return token, nil
}

func (j *JWT) Decode(token []byte) {

}

// /////////////////////////////////////////////

type Model struct {
	Id   int
	Name string
}

func main() {
	// jwt := jwt.New("secret")
	jwt := New([]byte("secret"))

	p := Model{
		Id:   1,
		Name: "username",
	}

	token, err := jwt.Encode(p)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Printf("%s\n", token)
}
