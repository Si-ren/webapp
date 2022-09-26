package utils

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTToken struct {
	privateKey *rsa.PrivateKey
	Issuer     string
	nowFunc    func() time.Time
}

func NewJWTToken(issuer string, key *rsa.PrivateKey) *JWTToken {
	return &JWTToken{
		privateKey: key,
		Issuer:     issuer,
		nowFunc:    time.Now,
	}

}

func (t *JWTToken) TokenGenerate(accountID string, expire time.Duration) (string, error) {
	nowSec := t.nowFunc().Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		ExpiresAt: nowSec + int64(expire.Seconds()),
		IssuedAt:  nowSec,
		Issuer:    t.Issuer,
		Subject:   accountID,
	})
	return tkn.SignedString(t.privateKey)
}
