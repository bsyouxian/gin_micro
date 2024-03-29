package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"user/core"
	"user/servies"
)

func main() {
	//etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"))
	//微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),//微服务名字
	micro.Address("127.0.0.1:8082"),
	micro.Registry(etcdReg))//etcd注册件
	//微服务初始化
	microService.Init()
	//服务注册
	_=servies.RegisterUserServiceHandler(microService.Server(),new(core.UserServicer))
	_=microService.Run()
}

