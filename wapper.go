/*
 * Copyrignt (c) dingdong.top. All Rights Reserved.
 * Author: xuzeshui@dingdong.top
 * Created Time: 2017-07-10 13:46:47
 * Last Modified: 2017-07-10 14:28:48
 * File Name: grpcpool/wapper.go
 * Description:
 */

//
package grpcpool

import (
	"golang.org/x/net/context"
)

type GrpcClientWapper struct {
	Ctx        context.Context
	CancelFunc context.CancelFunc
	Conn       *ClientConn
}

// 使用完之后一定要调用一次进行释放
func (p *GrpcClientWapper) Close() {
	if p.Conn != nil {
		p.Conn.Close()
	}
	if p.CancelFunc != nil {
		p.CancelFunc()
	}
}

// 如果调用方法异常,需要设置异常
func (p *GrpcClientWapper) OnCallError() {
	if p.Conn != nil {
		p.Conn.Unhealthy()
		p.Conn.ClientConn.Close()
	}
}
