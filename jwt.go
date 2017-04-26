package jwt

import (
	"encoding/json"
)

func Hello() string {
	return "hello world!"
}

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type JWT struct {
	Key string `json:"key"`
	Hdr string `json:"hdr"`
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
		res = []byte(err.Error())
	}
	jwt.Hdr = string(res)

	return jwt
}

func (j *JWT) Encode() string {
	return "Encode"
}
