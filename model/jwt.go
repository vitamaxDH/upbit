package model

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/url"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
)

const (
	SHA512 Algorithm = "SHA512"
)

type Algorithm string

type AuthTokenClaims struct {
	AccessKey    string `json:"access_key"`
	Nonce        string `json:"nonce"`
	QueryHash    string `json:"query_hash"`
	QueryHashAlg string `json:"query_hash_alg"`
	jwt.StandardClaims
}

func NewClaim() jwt.MapClaims {
	return jwt.MapClaims{
		"access_key": "YOUR ACCESS KEY",
		"nonce":      uuid.NewString(),
	}
}

func NewClaimForQuery(queries url.Values) jwt.MapClaims {
	claim := NewClaim()
	claim["query_hash"] = HashQuery(queries)
	claim["query_hash_alg"] = SHA512
	return claim
}

func GetBearerToken(token *jwt.Token) string {
	tokenStr, _ := token.SignedString([]byte("YOUR SECRET KEY"))
	return "Bearer " + tokenStr
}

func GetQueryString(queryStrings map[string]string) string {
	length := len(queryStrings)
	if length == 0 {
		return ""
	}
	var result string
	cnt := 0
	for key, value := range queryStrings {
		cnt++
		result += fmt.Sprintf("key[%s]=%s", key, value)
		if cnt != length {
			result += "&"
		}
	}
	return result
}

func HashQuery(query url.Values) string {
	sha := sha512.New()
	sha.Write([]byte(query.Encode()))
	return hex.EncodeToString(sha.Sum(nil))
}
