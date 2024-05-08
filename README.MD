### 使用生成工具

```shell
go install \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
google.golang.org/protobuf/cmd/protoc-gen-go \
google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### 使用 buf 安装 这个不好用

```shell
go install github.com/bufbuild/buf/cmd/buf@latest

buf generate ./proto #  ./proto 代表 proto 文件所在路径
```


###

```shell
github.com/googleapis/googleapis
```