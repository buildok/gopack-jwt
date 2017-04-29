// package jwt
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	// "time"
)

type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type JWT struct {
	Key      []byte
	Segments [3]string
}

func New(key string) *JWT {
	j := new(JWT)
	j.Key = []byte(key)

	h, _ := json.Marshal(header{Alg: "HS256", Typ: "JWT"})
	j.Segments[0] = base64.StdEncoding.EncodeToString(h)

	return j
}

func (j *JWT) Encode(payload interface{}) (string, error) {
	p, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	j.Segments[1] = base64.StdEncoding.EncodeToString(p)

	sign := hmac.New(sha256.New, j.Key)
	sign.Write([]byte(strings.Join(j.Segments[:2], ".")))

	j.Segments[2] = base64.StdEncoding.EncodeToString(sign.Sum(nil))

	return strings.Join(j.Segments[:], "."), nil
}

func (j *JWT) Decode(payload interface{}, token string) error {
	segs := strings.Split(token, ".")

	p, err := base64.StdEncoding.DecodeString(segs[1])
	if err != nil {
		return err
	}
	err = json.Unmarshal(p, payload)
	if err != nil {
		return err
	}

	return nil
}

func (j *JWT) Validate(payload interface{}, token string) (bool, error) {
	enc_p, err := j.Encode(payload)
	if err != nil {
		return false, err
	}

	mac1 := []byte(strings.Split(enc_p, ".")[2])
	mac2 := []byte(strings.Split(token, ".")[2])

	if !hmac.Equal(mac1, mac2) {
		return false, nil
	}

	return true, nil
}

// /////////////////////////////////////////////

type Model struct {
	Id   int
	Name string
}

func main() {
	// jwt := jwt.New("secret")
	jwt := New("secret")

	p := Model{
		Id:   1,
		Name: "username",
	}

	token, err := jwt.Encode(p)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Printf("%s\n", token)

	m := new(Model)
	if jwt.Decode(m, token) != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Printf("model\t%v\n", m)

	ok, err := jwt.Validate(m, token)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	fmt.Printf("result\t%v\n", ok)
}
