# gRPC 最小学习示例

这个目录演示一个最小的 Go gRPC 调用链：

1. 在 `proto/hello.proto` 里定义服务、请求和响应。
2. 用 `protoc` 生成 Go 代码到 `pb` 包。
3. `server/main.go` 实现并启动 gRPC 服务。
4. `client/main.go` 连接服务端并调用 `SayHello`。

## 目录结构

```text
grpcdemo/
  proto/hello.proto   # protobuf + gRPC 接口定义
  pb/                 # protoc 生成的 Go 代码
  server/main.go      # 服务端：实现 Greeter 服务
  client/main.go      # 客户端：调用 Greeter/SayHello
```

## 重新生成 pb 代码

修改 `proto/hello.proto` 后，在项目根目录运行：

```powershell
protoc --go_out=. --go_opt=module=Gocodes --go-grpc_out=. --go-grpc_opt=module=Gocodes grpcdemo/proto/hello.proto
```

这条命令会根据 `option go_package = "Gocodes/grpcdemo/pb;pb";`
把生成文件放到 `grpcdemo/pb`。

## 运行

先开一个终端启动服务端：

```powershell
go run ./grpcdemo/server
```

再开另一个终端运行客户端：

```powershell
go run ./grpcdemo/client -name Alice
```

你会看到类似输出：

```text
server replied: Hello, Alice! This response came from gRPC.
```

## 学习重点

- `.proto` 是客户端和服务端共同遵守的接口契约。
- `service` 定义可以被远程调用的方法。
- `message` 定义请求和响应的数据结构。
- `protoc` 根据 `.proto` 生成强类型 Go 代码。
- 服务端实现生成出来的 `GreeterServer` 接口。
- 客户端使用生成出来的 `GreeterClient` 发起 RPC 调用。
- `context.WithTimeout` 是 gRPC 调用里很常见的保护手段。
