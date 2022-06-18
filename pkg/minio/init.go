/*
 * @Author: a76yyyy q981331502@163.com
 * @Date: 2022-06-12 10:00:59
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:54:49
 * @FilePath: /tiktok/pkg/minio/init.go
 * @Description: Minio 对象存储初始化
 */

package minio

import (
	"github.com/a76yyyy/tiktok/pkg/ttviper"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	minioClient          *minio.Client
	Config               = ttviper.ConfigInit("TIKTOK_MINIO", "minioConfig")
	MinioEndpoint        = Config.Viper.GetString("minio.Endpoint")
	MinioAccessKeyId     = Config.Viper.GetString("minio.AccessKeyId")
	MinioSecretAccessKey = Config.Viper.GetString("minio.SecretAccessKey")
	MinioUseSSL          = Config.Viper.GetBool("minio.UseSSL")
	MinioVideoBucketName = Config.Viper.GetString("minio.VideoBucketName")
)

// Minio 对象存储初始化
func init() {
	client, err := minio.New(MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKeyId, MinioSecretAccessKey, ""),
		Secure: MinioUseSSL,
	})
	if err != nil {
		klog.Errorf("minio client init failed: %v", err)
	}
	// fmt.Println(client)
	klog.Debug("minio client init successfully")
	minioClient = client
	if err := CreateBucket(MinioVideoBucketName); err != nil {
		klog.Errorf("minio client init failed: %v", err)
	}
}
