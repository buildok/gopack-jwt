package jwt

import (
	"encoding/json"
)

func Hello() string {
	return "hello world!"
}

type header struct {
	alg string
	typ string
}

type JWT struct {
	key string
	hdr string
}

func New(key string) *JWT {
	jwt := new(JWT)
	jwt.key = key

	hdr := header{
		alg: "HS256",
		typ: "JWT",
	}
	res, err := json.Marshal(hdr)
	if err != nil {
		res = err.Error()
	}
	jwt.hdr = string(res)

	return jwt
}

func (j *JWT) Encode() string {
	return "Encode"
}
