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
 * @Date: 2022-06-08 16:22:35
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:49:57
 * @FilePath: /tiktok/pkg/errno/errno.go
 * @Description: 错误码报错业务逻辑
 */

package errno

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrCode int
	ErrMsg  string
}

// Err represents an error
type Err struct {
	ErrCode int
	ErrMsg  string
	Err     error
}

type HttpErr struct {
	StatusCode int
	ErrNo      ErrNo
}

func NewErrNo(code int, msg string) ErrNo {
	return ErrNo{code, msg}
}

func NewHttpErr(code int, httpcode int, msg string) HttpErr {
	return HttpErr{
		StatusCode: httpcode,
		ErrNo:      ErrNo{ErrCode: code, ErrMsg: msg},
	}
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrMsg = msg
	return e
}

func NewErr(errno *ErrNo, err error) *Err {
	return &Err{ErrCode: errno.ErrCode, ErrMsg: errno.ErrMsg, Err: err}
}

func (err *Err) Add(message string) error {
	//err.ErrMsg = fmt.Sprintf("%s %s", err.ErrMsg, message)
	err.ErrMsg += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	//return err.ErrMsg = fmt.Sprintf("%s %s", err.ErrMsg, fmt.Sprintf(format, args...))
	err.ErrMsg += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.ErrCode, err.ErrMsg, err.Err)
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.ErrCode
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return Success.ErrCode, Success.ErrMsg
	}

	switch typed := err.(type) {
	case *Err:
		return typed.ErrCode, typed.ErrMsg
	case *ErrNo:
		return typed.ErrCode, typed.ErrMsg
	default:
	}

	return ErrUnknown.ErrCode, err.Error()
}

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ErrUnknown
	s.ErrMsg = err.Error()
	return s
}
