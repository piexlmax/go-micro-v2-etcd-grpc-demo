package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	hello "grpc2/server/proto"
)

func main(){
	reg := etcdv3.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
		)
	service := micro.NewService(
		micro.Registry(reg),
		)
	client := hello.NewSayService("hello",service.Client())
	res,err := client.Hello(context.Background(),&hello.Req{Name: "test"})
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(res)
}
