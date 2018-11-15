package main

import (
	"auditIntegral/_public/config"
	"auditIntegral/_public/log"
	"auditIntegral/systemSetup/handler"

	"auditIntegral/systemSetup/subscriber"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"go.uber.org/zap"

	dictionaries "auditIntegral/systemSetup/proto/dictionaries"
	example "auditIntegral/systemSetup/proto/example"
)

func main() {
	log.Init(config.SystemSetupNameSpace)
	logger := log.Instance()
	// New Service
	service := micro.NewService(
		micro.Name(config.NameSpace+config.SystemSetupNameSpace),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(context *cli.Context) {
			logger.Info("info", zap.Any(config.SystemSetupNameSpace, "service start"))
			// Register Handler
			example.RegisterExampleHandler(service.Server(), new(handler.Example))
			// 字典管理
			//dictionaries.RegisterDictionariesHandler(service.Server(), new(handler.Dictionaries))
			dictionaries.RegisterDictionariesHandler(service.Server(), handler.NewDictionariesExtHandler())

			// Register Struct as Subscriber
			micro.RegisterSubscriber(config.NameSpace+config.SystemSetupNameSpace, service.Server(), new(subscriber.Example))

			// Register Function as Subscriber
			micro.RegisterSubscriber(config.NameSpace+config.SystemSetupNameSpace, service.Server(), subscriber.Handler)
		}),
		micro.AfterStop(func() error {
			logger.Info("info", zap.Any(config.SystemSetupNameSpace, "service stop"))
			return nil
		}),
	)

	// Run service
	if err := service.Run(); err != nil {
		logger.Panic(config.SystemSetupNameSpace + " service startup failure")
		logger.Error("error", zap.Error(err))
	}
}
