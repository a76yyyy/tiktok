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

var _ endpoint.Middleware = ClientMiddleware

// ClientMiddleware client middleware print server address 、rpc timeout and connection timeout
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
