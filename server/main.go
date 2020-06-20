package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"grpc2/server/model"
	hello "grpc2/server/proto"
)

//和v1的差别  proto解释器要用go-micro v2的包 官方示例不可以直接go get 需要到本目录下 强制打开mod模式 然后go get
// 这种情况下去编译proto就是v2版本的可以和v1一样使用  其他流程同v1版本


func main()  {
	//使用etcd创建注册中心
	reg  := etcdv3.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
		)
	//注册服务端
	servers := micro.NewService(
		micro.Registry(reg),
		micro.Name("hello"),  // 注意name用来做服务发现
		)
	servers.Init()
	//同样是调用proto给我们生成的方法来注册handler
	hello.RegisterSayHandler(servers.Server(),new(model.Say))
	if err:=servers.Run();err!=nil{
		panic(err)
	}
}