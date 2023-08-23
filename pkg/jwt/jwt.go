package jwt

import (
	"github.com/golang-jwt/jwt"
	"tiktok-backend/pkg/constants"
	"tiktok-backend/pkg/errno"
	"time"
)

type CustomClaims struct {
	Id  int64
	Exp int64
	jwt.StandardClaims
}

var signingKey = []byte(constants.SecretKey)

// CreateToken create a new token
func CreateToken(claims CustomClaims) (string, error) {
	expire := time.Now().Add(time.Hour)
	claims.Exp = expire.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

// ParseToken parses the token for id
func ParseToken(tokenString string) (int64, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return signingKey, nil
		},
	)
	if err != nil {
		return -1, errno.TokenErr
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.Id, nil
	}
	return -1, errno.TokenErr
}
