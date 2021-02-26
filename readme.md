# go-srv-kit

toolkit for srv_xxx

env : go version go1.15.6 with go mod

- protoc-gen-go google.golang.org/protobuf v1.25.0
- protoc-gen-go-grpc google.golang.org/grpc v1.36.0
- protoc-gen-go-tkform google.golang.org/protobuf v1.25.0 , 在message编译为go之后，添加form标签
- protoc-gen-go-tkgrpc github.com/golang/protobuf@1.4.3 , 早期的grpc生成插件
- protoc-gen-go-tktars github.com/TarsCloud/TarsGo@v1.1.6 , 适配最新的protobuf@v1.25.0

## gen protobuf

安装 protoc-gen-go , 命令如下：

- protoc-gen-go : go get google.golang.org/protobuf/cmd/protoc-gen-go
- protoc-gen-go-grpc : go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
- protoc-gen-go-tkform : go get github.com/ikaiguang/protoc-gen-go/cmd/protoc-gen-go-tkform
- protoc-gen-go-tkgrpc : go get github.com/ikaiguang/protoc-gen-go/cmd/protoc-gen-go-tkgrpc
- protoc-gen-go-tktars : go get github.com/ikaiguang/protoc-gen-go/cmd/protoc-gen-go-tktars

```shell script

# 编译示例
protoc -I. -I%GOPATH%/src --go-tkform_out=. --go-tkform_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./*.proto
protoc -I. -I%GOPATH%/src --go-tkform_out=. --go-tkform_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./*.proto
protoc -I. -I$GOPATH/src --go-tkform_out=. --go-tkform_opt=paths=source_relative --go-tktars_out=. --go-tktars_opt=paths=source_relative ./*.proto
protoc -I. -I$GOPATH/src --go-tkform_out=. --go-tkform_opt=paths=source_relative --go-tktars_out=. --go-tktars_opt=paths=source_relative ./*.proto

```

## 运行测试

go test -v ./example/testdata
