// Copyright 2022 a76yyyy && CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * @Author: a76yyyy q981331502@163.com
 * @Date: 2022-06-11 00:10:37
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:53:05
 * @FilePath: /tiktok/pkg/jwt/jwt.go
 * @Description:  基于 http://github.com/golang-jwt/jwt 的代码封装
 */

package jwt

import (
	code "github.com/a76yyyy/ErrnoCode"
	"github.com/a76yyyy/errors"
	"github.com/golang-jwt/jwt"
)

// JWT signing Key
type JWT struct {
	SigningKey []byte
}

var (
	ErrTokenExpired     = errors.WithCode(code.ErrExpired, "Token expired")
	ErrTokenNotValidYet = errors.WithCode(code.ErrValidation, "Token is not active yet")
	ErrTokenMalformed   = errors.WithCode(code.ErrTokenInvalid, "That's not even a token")
	ErrTokenInvalid     = errors.WithCode(code.ErrTokenInvalid, "Couldn't handle this token")
)

// private claims, share information between parties that agree on using them
// CustomClaims Structured version of Claims Section, as referenced at https://tools.ietf.org/html/rfc7519#section-4.1 See examples for how to use this with your own claim types
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

// CreateToken creates a new token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// zap.S().Debugf(token.SigningString())
	return token.SignedString(j.SigningKey)

}

// ParseToken parses the token.
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
	// verify the token claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}
