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

buf generate ./proto #  ./proto 代表 proto 文件所在路径 那么 buf.yaml 需要放在 ./proto  下面  buf.gen.yaml 放在根目录
```

```yaml
# buf.gen.yaml 
version: v1
plugins:
  - plugin: go
    out: gen/go
    opt:
      - paths=source_relative
  - plugin: go-grpc
    out: gen/go
    opt:
      - paths=source_relative
  - plugin: grpc-gateway
    out: gen/go
    opt:
      - paths=source_relative
      - generate_unbound_methods=true
```

```yaml
# buf.yaml
version: v1
deps:
  - buf.build/googleapis/googleapis
breaking:
  use:
    - FILE
lint:
  use:
    - DEFAULT
```

###

```shell
github.com/googleapis/googleapis
```