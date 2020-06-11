package utils

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	Expiration time.Time
	secret     []byte
}

type claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewJWT(exp int, sec string) (*JWT, error) {
	secFile, err := ioutil.ReadFile(sec)
	if err != nil {
		return nil, err
	}
	return &JWT{
		Expiration: time.Now().Add(time.Duration(exp) * time.Second),
		secret:     []byte(secFile),
	}, nil
}

func (j *JWT) GenerateToken(username, email, ctxID string) (tokenString string, err error) {
	claims := &claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			Issuer:    `upocwin`,
			IssuedAt:  time.Now().Unix(),
			Subject:   email,
			Id:        ctxID,
			ExpiresAt: j.Expiration.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.secret)
}

func (j *JWT) ValidateToken(tokenString string) (ctxID string, statusCode int, err error) {
	claims := &claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			statusCode = http.StatusUnauthorized
			return
		}
		statusCode = http.StatusBadRequest
		return
	}
	if !tkn.Valid {
		statusCode = http.StatusUnauthorized
		return
	}

	return claims.Id, 200, nil
}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	claims := &claims{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	claims.ExpiresAt = j.Expiration.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}
