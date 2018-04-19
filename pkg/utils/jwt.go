package utils

import (
	"github.com/GallenHu/bookmarkgo/pkg/setting"
	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.AppSecret)

type Claims struct {
	Mail string `json:"mail"`
	Id int `json:"id"`
	jwt.StandardClaims
}

func GenerateToken(mail string, id int) (string, error) {
	claims := Claims{
		mail,
		id,
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
