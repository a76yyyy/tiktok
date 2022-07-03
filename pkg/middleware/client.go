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
 * @LastEditTime: 2022-06-19 00:53:17
 * @FilePath: /tiktok/pkg/middleware/client.go
 * @Description: RPC Client Middleware
 */

package middleware

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

// 让编译器去检查ClientMiddleware和endpoint.Middleware是不是相同类型，如果不是则会报错
var _ endpoint.Middleware = ClientMiddleware

// ClientMiddleware client middleware print server address 、rpc timeout and connection timeout
// 相当于对Endpoint进行包装，在调用前输出一些信息
func ClientMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get server information
		klog.Infof("server address: %v, rpc timeout: %v, readwrite timeout: %v", ri.To().Address(), ri.Config().RPCTimeout(), ri.Config().ConnectTimeout())
		if err = next(ctx, req, resp); err != nil {
			return err
		}
		return nil
	}
}
