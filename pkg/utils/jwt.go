package utils

import (
	"time"
	"github.com/GallenHu/bookmarkgo/pkg/setting"
	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.AppSecret)

type Claims struct {
	Id int `json:"id"`
	Ctime int64 `json:"ctime"`
	jwt.StandardClaims
}

func GenerateToken(id int) (string, error) {
	ctime := time.Now().Unix() // currenttime 单位s,打印结果:1491888244
	claims := Claims{
		id,
		ctime,
		jwt.StandardClaims {
			ExpiresAt : 0,
			Issuer : "bookmarkgo",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
	})

	if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            return claims, nil
        }
	}

	return nil, err
}
