# jwt
Golang package for JWT encode/decode

Getting
```bash
go get github.com/buildok/jwt
```

main.go
```go
package main

import (
	"fmt"
	"github.com/buildok/jwt"
)

type Model struct {
	Id   int
	Name string
}

func main() {
	jwt := jwt.New("secret")

	// some data
	p := Model{
		Id:   1,
		Name: "username",
	}

	// calculate token
	token, _ := jwt.Encode(p)
	fmt.Printf("token:\t%s\n", token)

	// decoding
	m := new(Model)
	jwt.Decode(m, token)

	fmt.Printf("data:\t%#v\n", m)
	fmt.Printf("claims:\t%#v\n", jwt.Claims)

	// validate
	ok, _ := jwt.Validate(m, token)
	fmt.Printf("result:\t%v\n", ok)
}

```

Output:
```bash
token:  eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiTmFtZSI6InVzZXJuYW1lIn0=.1nWSFwxdnI1ur2AIE4Si/McfBaiMf8pA5kk/K02SJl0=
data:   &main.Model{Id:1, Name:"username"}
claims: jwt.Claims{Iss:"", Sub:"", Aud:"", Exp:0, Nbf:0, Iat:0, Jti:""}
result: true
```