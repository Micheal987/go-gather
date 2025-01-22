package jwtCombined

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateToken() {
	mySigningKey := []byte("AllYourBase")
	myclaims := MyCustomClaims{
		"admin",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}
	fmt.Printf("foo: %v\n", myclaims.Username)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myclaims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Println(ss, err)
	SampleToken(ss)
}

func SampleToken(tokenString string) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Println(token.Claims.(*MyCustomClaims).Username)
}
