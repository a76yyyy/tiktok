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
 * @Date: 2022-06-11 10:20:50
 * @LastEditors: a76yyyy q981331502@163.com
 * @LastEditTime: 2022-06-19 00:53:50
 * @FilePath: /tiktok/pkg/middleware/common.go
 * @Description: RPC Common Middleware
 */

package middleware

import (
	"context"

	"github.com/a76yyyy/tiktok/pkg/dlog"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var _ endpoint.Middleware = CommonMiddleware

func init() {
	var logger = dlog.InitLog(3)
	defer logger.Sync()

	klog.SetLogger(logger)
}

// CommonMiddleware common middleware print some rpc info、real request and real response
func CommonMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get real request
		klog.Debugf("real request: %+v", req)
		// get remote service information
		klog.Debugf("remote service name: %s, remote method: %s", ri.To().ServiceName(), ri.To().Method())
		if err = next(ctx, req, resp); err != nil {
			return err
		}
		// get real response
		klog.Infof("real response: %+v", resp)
		return nil
	}
}
