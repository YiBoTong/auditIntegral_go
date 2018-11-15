# Install Guide

## Consul
```cmd
brew install consul
consul agent -dev
```
or
```cmd
docker run consul
```
## Multicast DNS
We can use multicast DNS for zero dependency discovery
Pass `--registry=mdns` to any commands e.g micro `--registry=mdns list services`
## Protobuf
```cmd
go get github.com/micro/protoc-gen-micro
```
## Go Micro
Go Micro is an RPC framework for development microservices in G
```cmd
go get github.com/micro/go-micro
```
## Toolkit
```cmd
go get github.com/micro/micro
```
## List services
```cmd
$ micro list services
consul
go.micro.srv.xxx
```
## Get Service
```cmd
micro get service go.micro.srv.xxx
```
## Run API
`MICRO_API_HANDLER`sets the http handler  
`MICRO_API_NAMESPACE`sets the service namespace
```cmd
MICRO_API_HANDLER=rpc MICRO_API_NAMESPACE=com.ybt.ai micro api
```