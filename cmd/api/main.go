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
 * @Date: 2022-06-11 10:14:40
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-12-09 23:48:52
 * @FilePath: /tiktok/cmd/api/main.go
 * @Description: 使用 Hertz 提供 API 服务将 HTTP 请求发送给 RPC 微服务端
 */

// 使用 Hertz 提供 API 服务将 HTTP 请求发送给 RPC 微服务端
package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/a76yyyy/tiktok/cmd/api/handlers"
	"github.com/a76yyyy/tiktok/cmd/api/rpc"
	"github.com/a76yyyy/tiktok/pkg/dlog"
	"github.com/a76yyyy/tiktok/pkg/jwt"
	"github.com/a76yyyy/tiktok/pkg/ttviper"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/network/netpoll"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/hertz-contrib/gzip"
	h2config "github.com/hertz-contrib/http2/config"
	"github.com/hertz-contrib/http2/factory"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"github.com/hertz-contrib/registry/etcd"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
)

type HertzCfg struct {
	UseNetpoll bool  `json:"UseNetpoll" yaml:"UseNetpoll"`
	Http2      Http2 `json:"Http2" yaml:"Http2"`
	Tls        Tls   `json:"Tls" yaml:"Tls"`
}

type Tls struct {
	Enable bool `json:"Enable" yaml:"Enable"`
	Cfg    tls.Config
	Cert   string `json:"CertFile" yaml:"CertFile"`
	Key    string `json:"KeyFile" yaml:"KeyFile"`
	ALPN   bool   `json:"ALPN" yaml:"ALPN"`
}

type Http2 struct {
	Enable           bool     `json:"Enable" yaml:"Enable"`
	DisableKeepalive bool     `json:"DisableKeepalive" yaml:"DisableKeepalive"`
	ReadTimeout      Duration `json:"ReadTimeout" yaml:"ReadTimeout"`
}

type Duration struct {
	time.Duration
}

var (
	Config      = ttviper.ConfigInit("TIKTOK_API", "apiConfig")
	ServiceName = Config.Viper.GetString("Server.Name")
	ServiceAddr = fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	EtcdAddress = fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	Jwt         *jwt.JWT
	hertzCfg    HertzCfg
)

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}

// 初始化 API 配置
func Init() {
	rpc.InitRPC(&Config)
	Jwt = jwt.NewJWT([]byte(Config.Viper.GetString("JWT.signingKey")))
}

func InitHertzCfg() {
	hertzV, err := json.Marshal(Config.Viper.Sub("Hertz").AllSettings())
	if err != nil {
		hlog.Fatalf("Error marshalling Hertz config %s", err)
	}
	if err := json.Unmarshal(hertzV, &hertzCfg); err != nil {
		hlog.Fatalf("Error unmarshalling Hertz config %s", err)
	}
}

// 初始化 Hertz
func InitHertz() *server.Hertz {
	InitHertzCfg()

	opts := []config.Option{server.WithHostPorts(ServiceAddr)}

	// 服务注册
	if Config.Viper.GetBool("Etcd.Enable") {
		r, err := etcd.NewEtcdRegistry([]string{EtcdAddress})
		if err != nil {
			hlog.Fatal(err)
		}
		opts = append(opts, server.WithRegistry(r, &registry.Info{
			ServiceName: ServiceName,
			Addr:        utils.NewNetAddr("tcp", ServiceAddr),
			Weight:      10,
			Tags:        nil,
		}))
	}

	// 链路追踪
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(ServiceName),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())
	tracer, tracerCfg := hertztracing.NewServerTracer()
	opts = append(opts, tracer)

	// 网络库
	hertzNet := standard.NewTransporter
	if hertzCfg.UseNetpoll {
		hertzNet = netpoll.NewTransporter
	}
	opts = append(opts, server.WithTransport(hertzNet))

	// TLS & Http2
	tlsEnable := hertzCfg.Tls.Enable
	h2Enable := hertzCfg.Http2.Enable
	hertzCfg.Tls.Cfg = tls.Config{
		MinVersion:       tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
	}
	if tlsEnable {
		cert, err := tls.LoadX509KeyPair(hertzCfg.Tls.Cert, hertzCfg.Tls.Key)
		if err != nil {
			hlog.Error(err)
		}
		hertzCfg.Tls.Cfg.Certificates = append(hertzCfg.Tls.Cfg.Certificates, cert)
		opts = append(opts, server.WithTLS(&hertzCfg.Tls.Cfg))

		if alpn := hertzCfg.Tls.ALPN; alpn {
			opts = append(opts, server.WithALPN(alpn))
		}
	} else if h2Enable {
		opts = append(opts, server.WithH2C(h2Enable))
	}

	// Hertz
	h := server.Default(opts...)
	h.Use(gzip.Gzip(gzip.DefaultCompression),
		hertztracing.ServerMiddleware(tracerCfg))

	// Protocol
	if h2Enable {
		h.AddProtocol("h2", factory.NewServerFactory(
			h2config.WithReadTimeout(hertzCfg.Http2.ReadTimeout.Duration),
			h2config.WithDisableKeepAlive(hertzCfg.Http2.DisableKeepalive)))
		if tlsEnable {
			hertzCfg.Tls.Cfg.NextProtos = append(hertzCfg.Tls.Cfg.NextProtos, "h2")
		}
	}

	return h
}

// 注册 Router组
func registerGroup(h *server.Hertz) {
	douyin := h.Group("/douyin")

	user := douyin.Group("/user")
	user.POST("/login/", handlers.Login)
	user.POST("/register/", handlers.Register)
	user.GET("/", handlers.GetUserById)

	video := douyin.Group("/feed")
	video.GET("/", handlers.GetUserFeed)

	publish := douyin.Group("/publish")
	publish.POST("/action/", handlers.PublishAction)
	publish.GET("/list/", handlers.PublishList)

	favorite := douyin.Group("/favorite")
	favorite.POST("/action/", handlers.FavoriteAction)
	favorite.GET("/list/", handlers.FavoriteList)

	comment := douyin.Group("/comment")
	comment.POST("/action/", handlers.CommentAction)
	comment.GET("/list/", handlers.CommentList)

	relation := douyin.Group("/relation")
	relation.POST("/action/", handlers.RelationAction)
	relation.GET("/follow/list/", handlers.RelationFollowList)
	relation.GET("/follower/list/", handlers.RelationFollowerList)
}

// 初始化 Hertz API 及 Router
func main() {
	logger := dlog.InitHertzLog(3)
	defer logger.Sync()

	hlog.SetLogger(logger)
	hlog.SetSystemLogger(logger)

	Init()

	h := InitHertz()

	pprof.Register(h)

	registerGroup(h)

	h.Spin()
}
