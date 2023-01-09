package main

import (
	"goods_services/proto"
	"google.golang.org/grpc"
	"net"
)

// @program:     goods_servies
// @file:        main.go
// @author:      bowen
// @create:      2023-01-08 17:03
// @description:
type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	// 创建grpc服务
	s := grpc.NewServer()
	// 注册服务
	proto.RegisterGoodsServer(s, &GoodsServer{})
	//跑起来
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
