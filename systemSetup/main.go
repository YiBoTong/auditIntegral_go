package main

import (
	"auditIntegral/systemSetup/handler"
	"auditIntegral/systemSetup/subscriber"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	dictionaries "auditIntegral/systemSetup/proto/dictionaries"
	example "auditIntegral/systemSetup/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.systemSetup"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))
	// 字典管理
	dictionaries.RegisterDictionariesHandler(service.Server(), new(handler.Dictionaries))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.systemSetup", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.systemSetup", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
