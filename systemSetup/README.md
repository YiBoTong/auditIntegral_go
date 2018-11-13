# SystemSetup Service

This is the SystemSetup service

Generated with

```
micro new auditIntegral/systemSetup --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.systemSetup
- Type: srv
- Alias: systemSetup

## Dependencies

Micro services depend on service discovery. The default is consul.  
[down URL](https://www.consul.io/downloads.html)
```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./systemSetup-srv
```

Build a docker image
```
make docker
```

## 编译proto  
进入到proto的上层目录然后执行
```cmd
protoc -I. --go_out=plugins=micro:. --micro_out=. proto/example/example.proto
```

## 运行
```cmd
MICRO_SERVER_ADDRESS=:50052 MICRO_REGISTRY=mdns go run main.go
```

## 服务检查
```cmd
micro list services
```

## API运行
```cmd
MICRO_API_HANDLER=rpc MICRO_API_NAMESPACE=go.micro.srv micro api
```
## API 测试
```
curl -XPOST -H 'Content-Type: application/json' -d '{"name": "John"}' http://localhost:8080/systemSetup/example/call
```