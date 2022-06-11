package jwt

import (
	"github.com/a76yyyy/errors"

	code "github.com/a76yyyy/ErrnoCode"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	SigningKey []byte
}

var (
	ErrTokenExpired     = errors.WithCode(code.ErrExpired, "Token expired")
	ErrTokenNotValidYet = errors.WithCode(code.ErrValidation, "Token is not active yet")
	ErrTokenMalformed   = errors.WithCode(code.ErrTokenInvalid, "That's not even a token")
	ErrTokenInvalid     = errors.WithCode(code.ErrTokenInvalid, "Couldn't handle this token")
)

type CustomClaims struct {
	Id          int64
	AuthorityId int64
	jwt.StandardClaims
}

func NewJWT(SigningKey []byte) *JWT {
	return &JWT{
		SigningKey,
	}
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//zap.S().Debugf(token.SigningString())
	return token.SignedString(j.SigningKey)

}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}

		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}
