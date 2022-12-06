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
 * @Date: 2022-06-12 09:32:40
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:03:14
 * @FilePath: /tiktok/cmd/publish/command/publish_action.go
 * @Description: 发布视频 操作业务逻辑
 */

package command

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strings"

	"github.com/a76yyyy/tiktok/dal/db"
	"github.com/a76yyyy/tiktok/kitex_gen/publish"
	"github.com/a76yyyy/tiktok/pkg/minio"
	"github.com/a76yyyy/tiktok/pkg/ttviper"
	"github.com/gofrs/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type PublishActionService struct {
	ctx context.Context
}

// NewPublishActionService new PublishActionService
func NewPublishActionService(ctx context.Context) *PublishActionService {
	return &PublishActionService{ctx: ctx}
}

// PublishAction publish video.
func (s *PublishActionService) PublishAction(req *publish.DouyinPublishActionRequest, uid int, cfg *ttviper.Config) (err error) {
	MinioVideoBucketName := minio.MinioVideoBucketName
	videoData := []byte(req.Data)

	// // 获取后缀
	// filetype := http.DetectContentType(videoData)

	// byte[] -> reader
	reader := bytes.NewReader(videoData)
	u2, err := uuid.NewV4()
	if err != nil {
		return err
	}
	fileName := u2.String() + "." + "mp4"
	// 上传视频
	err = minio.UploadFile(MinioVideoBucketName, fileName, reader, int64(len(videoData)))
	if err != nil {
		return err
	}
	// 获取视频链接
	url, err := minio.GetFileUrl(MinioVideoBucketName, fileName, 0)
	playUrl := strings.Split(url.String(), "?")[0]
	if err != nil {
		return err
	}

	u3, err := uuid.NewV4()
	if err != nil {
		return err
	}

	// 获取封面
	coverPath := u3.String() + "." + "jpg"
	coverData, err := readFrameAsJpeg(playUrl)
	if err != nil {
		return err
	}

	// 上传封面
	coverReader := bytes.NewReader(coverData)
	err = minio.UploadFile(MinioVideoBucketName, coverPath, coverReader, int64(len(coverData)))
	if err != nil {
		return err
	}

	// 获取封面链接
	coverUrl, err := minio.GetFileUrl(MinioVideoBucketName, coverPath, 0)
	if err != nil {
		return err
	}

	CoverUrl := strings.Split(coverUrl.String(), "?")[0]

	// 封装video
	videoModel := &db.Video{
		AuthorID:      uid,
		PlayUrl:       playUrl,
		CoverUrl:      CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         req.Title,
	}
	return db.CreateVideo(s.ctx, videoModel)
}

// ReadFrameAsJpeg
// 从视频流中截取一帧并返回 需要在本地环境中安装ffmpeg并将bin添加到环境变量
func readFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), err
}
