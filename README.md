# GRPC FULL

- to proof of concept grpc
- client side with react-ts
- server side withg golang (echo)


## Prerequisites
- make command
- golang complier
- node.js runtime
- docker

## Installation
Protoc plugin that generates code for gRPC-Web clients
```bash
brew install protoc-gen-grpc-web
```

Install the protocol compiler plugins for Go using the following commands:
```bash
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Update your PATH so that the protoc compiler can find the plugins:
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```




    
## Run Locally

Clone the project & Go to the project directory


start product service
```bash
  cd product-service
  make run
```

start auth service
```bash
  cd auth-service
  make run
```

start envoy gateway
```bash
  cd deployment
  make start-envoy
```

start react
```bash
  cd grpc-web-client
  make run
```

gen protobuf (optional)
```bash
  make protogen
```

## Documentation

[GRPC](https://grpc.io/)
[ENVOY](https://gateway.envoyproxy.io/v0.5.0/)
[GRPC-WEB](https://github.com/grpc/grpc-web)
