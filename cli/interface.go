/**
 * @Author: DollarKillerX
 * @Description: cli interface
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:41 2020/1/6
 */
package main

import "github.com/dollarkillerx/proto"

// 生成vodka系统
type Generator interface {
	Name() string
	Run(opt *Option, data *RPCData) error
}

type RPCData struct {
	Rpc     []*proto.RPC
	Service *proto.Service
	Message []*proto.Message
	Pkg     *proto.Package
}
