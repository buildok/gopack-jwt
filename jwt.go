/**
 * Simple encoding/decoding JSON Web Token
 * Support HS256
 * RFC 7519
 */

package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
)

/**
 * JWT header
 */
type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

/**
 * JWT claims
 */
type Claims struct {
	Iss string
	Sub string
	Aud string
	Exp int
	Nbf int
	Iat int
	Jti string
}

type JWT struct {
	Key      []byte
	Segments [3]string
	Claims
}

/**
 * Create new encoder/decoder.
 */
func New(key string) *JWT {
	j := new(JWT)
	j.Key = []byte(key)

	h, _ := json.Marshal(header{Alg: "HS256", Typ: "JWT"})
	j.Segments[0] = base64.StdEncoding.EncodeToString(h)

	return j
}

/**
 * Return JWT token
 */
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

/**
 * Decode data from token to payload
 */
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

	var fields map[string]interface{}
	json.Unmarshal(p, &fields)

	for i, v := range fields {
		switch i {

		case "iss":
			j.Iss = v.(string)
		case "sub":
			j.Sub = v.(string)
		case "aud":
			j.Aud = v.(string)
		case "exp":
			j.Exp = v.(int)
		case "nbf":
			j.Nbf = v.(int)
		case "iat":
			j.Iat = v.(int)
		case "jti":
			j.Jti = v.(string)

		}
	}

	return nil
}

/**
 * Validate payload data
 */
func (j *JWT) Validate(token string) (bool, error) {
	segs := strings.Split(token, ".")

	sign := hmac.New(sha256.New, j.Key)
	sign.Write([]byte(strings.Join(segs[:2], ".")))

	mac1 := []byte(segs[2])
	mac2 = []byte(base64.StdEncoding.EncodeToString(sign.Sum(nil)))

	if !hmac.Equal(mac1, mac2) {
		return false, nil
	}

	return true, nil
}
