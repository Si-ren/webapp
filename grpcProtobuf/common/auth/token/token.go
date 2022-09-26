package token

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

// https://pkg.go.dev/github.com/golang-jwt/jwt/v4
type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

func (j *JWTTokenVerifier) Verify(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwt.SigningMethodRS512, nil
	})
	if err != nil {
		logrus.Error("JWTToken Verify err: ", err)
		return "", err
	}
	if !t.Valid {
		logrus.Error("JWTToken Verify Failed: ", err)
		return "", err
	}
	//断言
	clm, ok := t.Claims.(*jwt.StandardClaims)
	if !ok {
		logrus.Error("JWTToken Verify Assert Failed: ", err)
		return "", err
	}
	//valid函数根据exp、iat、nbf判断jwt是否有效
	if err := clm.Valid(); err != nil {
		logrus.Error("JWTToken Verify Claim Valid: ", err)
		return "", err
	}

	return clm.Subject, nil
}
