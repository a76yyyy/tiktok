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
 * @Date: 2022-06-12 10:00:59
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 01:15:17
 * @FilePath: /tiktok/pkg/minio/minio_test.go
 * @Description: minio 测试
 */

package minio

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// (TiktokTesst)bucket name  ccontains invalid characters
// bucket name 只能用小写字母
func TestCreateBucket(t *testing.T) {
	CreateBucket("tiktoktest")
}

func TestUploadLocalFile(t *testing.T) {
	info, err := UploadLocalFile("tiktoktest", "test.mp4", "./test.mp4", "video/mp4")
	fmt.Println(info, err)
}

func TestUploadFile(t *testing.T) {
	file, _ := os.Open("./test.mp4")
	defer file.Close()
	fi, _ := os.Stat("./test.mp4")
	err := UploadFile("tiktoktest", "ceshi2", file, fi.Size())
	fmt.Println(err)
}

func TestGetFileUrl(t *testing.T) {
	url, err := GetFileUrl("tiktoktest", "test.mp4", 0)
	fmt.Println(url, err, strings.Split(url.String(), "?")[0])
	fmt.Println(url.Path, url.RawPath)
}
