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
 * @Date: 2022-06-19 01:08:23
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 01:17:09
 * @FilePath: /tiktok/pkg/jwt/jwt_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package jwt

import (
	"fmt"
	"testing"
)

func TestNewJWT(t *testing.T) {
	JWT := NewJWT([]byte("test"))
	fmt.Println(JWT)
}

func TestCreateToken(t *testing.T) {
	JWT := NewJWT([]byte("test"))
	fmt.Println(JWT)
	token, err := JWT.CreateToken(CustomClaims{
		Id: int64(10010),
	})
	fmt.Println(token, err)
}

func TestParseToken(t *testing.T) {
	JWT := NewJWT([]byte("test"))
	fmt.Println(JWT)
	token, err := JWT.CreateToken(CustomClaims{
		Id: int64(10010),
	})
	fmt.Println(token, err)
	claim, err := JWT.ParseToken(token)
	fmt.Println(claim, err)
}
