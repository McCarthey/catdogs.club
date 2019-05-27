package libs

import (
	"time"

	configs "catdogs.club/back-end/configs/common"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(configs.PwSalt)

type Claims struct {
	Openid string `json:"openid"`
	jwt.StandardClaims
}

func GenerateToken(openid string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		openid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "catdogs",
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
