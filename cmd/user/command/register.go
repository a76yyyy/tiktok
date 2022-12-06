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
 * @Date: 2022-06-10 19:39:14
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:17:32
 * @FilePath: /tiktok/cmd/user/command/register.go
 * @Description: 注册 操作业务逻辑
 */

package command

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/a76yyyy/tiktok/dal/db"
	"github.com/a76yyyy/tiktok/kitex_gen/user"
	"github.com/a76yyyy/tiktok/pkg/errno"
	"golang.org/x/crypto/argon2"
)

type CreateUserService struct {
	ctx context.Context
}

type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *user.DouyinUserRegisterRequest, argon2Params *Argon2Params) error {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.ErrUserAlreadyExist
	}

	// Pass the plaintext password and parameters to our generateFromPassword
	// helper function.
	passWord, err := generateFromPassword(req.Password, argon2Params)
	if err != nil {
		return err
	}
	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.Username,
		Password: passWord,
	}})
}

// generateFromPassword generate the hash from the password string with salt and iterations values.
// the encrypting algorithm is Argon2id.
func generateFromPassword(password string, argon2Params *Argon2Params) (encodedHash string, err error) {
	salt, err := generateRandomBytes(argon2Params.SaltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, argon2Params.Iterations, argon2Params.Memory, argon2Params.Parallelism, argon2Params.KeyLength)

	// Base64 encode the salt and hashed password.
	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, argon2Params.Memory, argon2Params.Iterations, argon2Params.Parallelism, base64Salt, base64Hash)

	return encodedHash, nil
}

// generateRandomBytes returns a random bytes.
func generateRandomBytes(saltLength uint32) ([]byte, error) {
	buf := make([]byte, saltLength)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
