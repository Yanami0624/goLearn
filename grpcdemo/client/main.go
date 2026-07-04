package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "Gocodes/grpcdemo/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := flag.String("addr", "localhost:50051", "gRPC server address")
	name := flag.String("name", "Go", "name sent to SayHello")
	age := flag.Int("age", 0, "age_")
	flag.Parse()

	// 为了让学习示例更短，这里使用 insecure 连接，也就是不启用 TLS。
	// 生产环境通常应该配置 TLS 或放在可信内网/网关之后。
	conn, err := grpc.NewClient(
		*addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("create client failed: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// 给这次 RPC 调用设置超时时间，避免服务端不可达时一直等待。
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	reply, err := client.SayHello(ctx, &pb.HelloRequest{Name: *name, Age: (int32)(*age)})
	if err != nil {
		log.Fatalf("SayHello failed: %v", err)
	}

	log.Printf("server replied: %s", reply.GetMessage())

	reply, _ = client.SayGoodbye(ctx, &pb.GoodbyeRequest{Name: *name})

	log.Printf("server replied: %s", reply.GetMessage())
}
