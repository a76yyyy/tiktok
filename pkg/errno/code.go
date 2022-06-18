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
 * @Date: 2022-06-08 16:23:30
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:49:39
 * @FilePath: /tiktok/pkg/errno/code.go
 * @Description: 错误码, 设计逻辑:https://github.com/a76yyyy/ErrnoCode
 */

// 错误码
package errno

import code "github.com/a76yyyy/ErrnoCode"

// HTTP Error
var (
	HttpSuccess                  = NewHttpErr(code.ErrSuccess, 200, "OK")
	ErrHttpUnknown               = NewHttpErr(code.ErrUnknown, 500, "Internal server error")
	ErrHttpBind                  = NewHttpErr(code.ErrBind, 400, "Error occurred while binding the request body to the struct")
	ErrHttpValidation            = NewHttpErr(code.ErrValidation, 400, "Validation failed")
	ErrHttpTokenInvalid          = NewHttpErr(code.ErrTokenInvalid, 401, "Token invalid")
	ErrHttpDatabase              = NewHttpErr(code.ErrDatabase, 500, "Database error")
	ErrHttpRecordNotFound        = NewHttpErr(code.ErrRecordNotFound, 500, "Record not found")
	ErrHttpInvalidTransaction    = NewHttpErr(code.ErrInvalidTransaction, 500, "Invalid transaction when you are trying to `Commit` or `Rollback`")
	ErrHttpNotImplemented        = NewHttpErr(code.ErrNotImplemented, 500, "Not implemented")
	ErrHttpMissingWhereClause    = NewHttpErr(code.ErrMissingWhereClause, 500, "Missing where clause")
	ErrHttpUnsupportedRelation   = NewHttpErr(code.ErrUnsupportedRelation, 500, "Unsupported relations")
	ErrHttpPrimaryKeyRequired    = NewHttpErr(code.ErrPrimaryKeyRequired, 500, "Primary keys required")
	ErrHttpModelValueRequired    = NewHttpErr(code.ErrModelValueRequired, 500, "Model value required")
	ErrHttpInvalidData           = NewHttpErr(code.ErrInvalidData, 500, "Unsupported data")
	ErrHttpUnsupportedDriver     = NewHttpErr(code.ErrUnsupportedDriver, 500, "Unsupported driver")
	ErrHttpRegistered            = NewHttpErr(code.ErrRegistered, 500, "Registered")
	ErrHttpInvalidField          = NewHttpErr(code.ErrInvalidField, 500, "Invalid field")
	ErrHttpEmptySlice            = NewHttpErr(code.ErrEmptySlice, 500, "Empty slice found")
	ErrHttpDryRunModeUnsupported = NewHttpErr(code.ErrDryRunModeUnsupported, 500, "Dry run mode unsupported")
	ErrHttpInvalidDB             = NewHttpErr(code.ErrInvalidDB, 500, "Invalid db")
	ErrHttpInvalidValue          = NewHttpErr(code.ErrInvalidValue, 500, "Invalid value")
	ErrHttpInvalidValueOfLength  = NewHttpErr(code.ErrInvalidValueOfLength, 500, "Invalid values do not match length")
	ErrHttpPreloadNotAllowed     = NewHttpErr(code.ErrPreloadNotAllowed, 500, "Preload is not allowed when count is used")
	ErrHttpEncrypt               = NewHttpErr(code.ErrEncrypt, 401, "Error occurred while encrypting the user password")
	ErrHttpSignatureInvalid      = NewHttpErr(code.ErrSignatureInvalid, 401, "Signature is invalid")
	ErrHttpExpired               = NewHttpErr(code.ErrExpired, 401, "Token expired")
	ErrHttpInvalidAuthHeader     = NewHttpErr(code.ErrInvalidAuthHeader, 401, "Invalid authorization header")
	ErrHttpMissingHeader         = NewHttpErr(code.ErrMissingHeader, 401, "The `Authorization` header was empty")
	ErrHttporExpired             = NewHttpErr(code.ErrorExpired, 401, "Token expired")
	ErrHttpPasswordIncorrect     = NewHttpErr(code.ErrPasswordIncorrect, 401, "Password was incorrect")
	ErrHttpPermissionDenied      = NewHttpErr(code.ErrPermissionDenied, 403, "Permission denied")
	ErrHttpEncodingFailed        = NewHttpErr(code.ErrEncodingFailed, 500, "Encoding failed due to an error with the data")
	ErrHttpDecodingFailed        = NewHttpErr(code.ErrDecodingFailed, 500, "Decoding failed due to an error with the data")
	ErrHttpInvalidJSON           = NewHttpErr(code.ErrInvalidJSON, 500, "Data is not valid JSON")
	ErrHttpEncodingJSON          = NewHttpErr(code.ErrEncodingJSON, 500, "JSON data could not be encoded")
	ErrHttpDecodingJSON          = NewHttpErr(code.ErrDecodingJSON, 500, "JSON data could not be decoded")
	ErrHttpInvalidYaml           = NewHttpErr(code.ErrInvalidYaml, 500, "Data is not valid Yaml")
	ErrHttpEncodingYaml          = NewHttpErr(code.ErrEncodingYaml, 500, "Yaml data could not be encoded")
	ErrHttpDecodingYaml          = NewHttpErr(code.ErrDecodingYaml, 500, "Yaml data could not be decoded")
	ErrHttpUserNotFound          = NewHttpErr(code.ErrUserNotFound, 404, "User not found")
	ErrHttpUserAlreadyExist      = NewHttpErr(code.ErrUserAlreadyExist, 400, "User already exist")
	ErrHttpReachMaxCount         = NewHttpErr(code.ErrReachMaxCount, 400, "Secret reach the max count")
	ErrHttpSecretNotFound        = NewHttpErr(code.ErrSecretNotFound, 404, "Secret not found")
	ErrHttpVideoNotFound         = NewHttpErr(code.ErrVideoNotFound, 400, "Video not found")
	ErrHttpInvalidHash           = NewHttpErr(code.ErrInvalidHash, 500, "Encoded hash is not in the correct format")
	ErrHttpIncompatibleVersion   = NewHttpErr(code.ErrIncompatibleVersion, 500, "Encoded hash is not in the correct format")
)

// Server Error
var (
	Success         = NewErrNo(code.ErrSuccess, "OK")
	ErrUnknown      = NewErrNo(code.ErrUnknown, "Internal server error")
	ErrBind         = NewErrNo(code.ErrBind, "Error occurred while binding the request body to the struct")
	ErrValidation   = NewErrNo(code.ErrValidation, "Validation failed")
	ErrTokenInvalid = NewErrNo(code.ErrTokenInvalid, "Token invalid")

	ErrDatabase              = NewErrNo(code.ErrDatabase, "Database error")
	ErrRecordNotFound        = NewErrNo(code.ErrRecordNotFound, "Record not found")
	ErrInvalidTransaction    = NewErrNo(code.ErrInvalidTransaction, "Invalid transaction when you are trying to `Commit` or `Rollback`")
	ErrNotImplemented        = NewErrNo(code.ErrNotImplemented, "Not implemented")
	ErrMissingWhereClause    = NewErrNo(code.ErrMissingWhereClause, "Missing where clause")
	ErrUnsupportedRelation   = NewErrNo(code.ErrUnsupportedRelation, "Unsupported relations")
	ErrPrimaryKeyRequired    = NewErrNo(code.ErrPrimaryKeyRequired, "Primary keys required")
	ErrModelValueRequired    = NewErrNo(code.ErrModelValueRequired, "Model value required")
	ErrInvalidData           = NewErrNo(code.ErrInvalidData, "Unsupported data")
	ErrUnsupportedDriver     = NewErrNo(code.ErrUnsupportedDriver, "Unsupported driver")
	ErrRegistered            = NewErrNo(code.ErrRegistered, "Registered")
	ErrInvalidField          = NewErrNo(code.ErrInvalidField, "Invalid field")
	ErrEmptySlice            = NewErrNo(code.ErrEmptySlice, "Empty slice found")
	ErrDryRunModeUnsupported = NewErrNo(code.ErrDryRunModeUnsupported, "Dry run mode unsupported")
	ErrInvalidDB             = NewErrNo(code.ErrInvalidDB, "Invalid db")
	ErrInvalidValue          = NewErrNo(code.ErrInvalidValue, "Invalid value")
	ErrInvalidValueOfLength  = NewErrNo(code.ErrInvalidValueOfLength, "Invalid values do not match length")
	ErrPreloadNotAllowed     = NewErrNo(code.ErrPreloadNotAllowed, "Preload is not allowed when count is used")

	ErrEncrypt           = NewErrNo(code.ErrEncrypt, "Error occurred while encrypting the user password")
	ErrSignatureInvalid  = NewErrNo(code.ErrSignatureInvalid, "Signature is invalid")
	ErrExpired           = NewErrNo(code.ErrExpired, "Token expired")
	ErrInvalidAuthHeader = NewErrNo(code.ErrInvalidAuthHeader, "Invalid authorization header")
	ErrMissingHeader     = NewErrNo(code.ErrMissingHeader, "The `Authorization` header was empty")
	ErrorExpired         = NewErrNo(code.ErrorExpired, "Token expired")
	ErrPasswordIncorrect = NewErrNo(code.ErrPasswordIncorrect, "Password was incorrect")
	ErrPermissionDenied  = NewErrNo(code.ErrPermissionDenied, "Permission denied")

	ErrEncodingFailed = NewErrNo(code.ErrEncodingFailed, "Encoding failed due to an error with the data")
	ErrDecodingFailed = NewErrNo(code.ErrDecodingFailed, "Decoding failed due to an error with the data")
	ErrInvalidJSON    = NewErrNo(code.ErrInvalidJSON, "Data is not valid JSON")
	ErrEncodingJSON   = NewErrNo(code.ErrEncodingJSON, "JSON data could not be encoded")
	ErrDecodingJSON   = NewErrNo(code.ErrDecodingJSON, "JSON data could not be decoded")
	ErrInvalidYaml    = NewErrNo(code.ErrInvalidYaml, "Data is not valid Yaml")
	ErrEncodingYaml   = NewErrNo(code.ErrEncodingYaml, "Yaml data could not be encoded")
	ErrDecodingYaml   = NewErrNo(code.ErrDecodingYaml, "Yaml data could not be decoded")

	ErrUserNotFound        = NewErrNo(code.ErrUserNotFound, "User not found")
	ErrUserAlreadyExist    = NewErrNo(code.ErrUserAlreadyExist, "User already exist")
	ErrReachMaxCount       = NewErrNo(code.ErrReachMaxCount, "Secret reach the max count")
	ErrSecretNotFound      = NewErrNo(code.ErrSecretNotFound, "Secret not found")
	ErrVideoNotFound       = NewErrNo(code.ErrVideoNotFound, "Video not found")
	ErrInvalidHash         = NewErrNo(code.ErrInvalidHash, "Encoded hash is not in the correct format")
	ErrIncompatibleVersion = NewErrNo(code.ErrIncompatibleVersion, "Encoded hash is not in the correct format")
)
