package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	pb "Gocodes/grpcdemo/pb"

	"google.golang.org/grpc"
)

// greeterServer 是我们自己实现的服务端结构体。
// 嵌入 UnimplementedGreeterServer 是 gRPC Go 的推荐写法：
// 以后 proto 里新增 RPC 方法时，旧代码仍然能编译，并返回明确的未实现错误。
type greeterServer struct {
	pb.UnimplementedGreeterServer
}

// SayHello 实现 hello.proto 中定义的 Greeter/SayHello RPC。
// context.Context 可以用于超时、取消请求、传递 metadata 等。
func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	name := strings.TrimSpace(req.GetName())
	if name == "" {
		name = "gRPC learner"
	}

	age := req.GetAge()
	if age == 0 {
		age = 18
	}

	time.Sleep(time.Second * 4)

	return &pb.HelloReply{
		Message: "Hello, " + name + ", you are " + fmt.Sprintf("%d", age) + " years old",
	}, nil
}

func (s *greeterServer) SayGoodbye(ctx context.Context, req *pb.GoodbyeRequest) (*pb.HelloReply, error) {
	name := req.GetName()
	if name == "" {
		name = "gRPC learner"
	}

	return &pb.HelloReply{
		Message: "Goodbye " + name,
	}, nil
}

func main() {
	// gRPC 默认使用 HTTP/2。这里监听本机 50051 端口，这是示例里常用的端口。
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("listen failed: %v", err)
	}

	grpcServer := grpc.NewServer()

	// 把我们实现的 greeterServer 注册到 gRPC 服务器中。
	// 注册后，客户端才能调用 Greeter 服务里的 SayHello 方法。
	pb.RegisterGreeterServer(grpcServer, &greeterServer{})

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("serve failed: %v", err)
	}
}
