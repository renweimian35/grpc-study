package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	pb "grpc-study/hello-server/proto"
	"log"
	"net"
)

func main() {
	// 配置好证书文件
	creds, _ := credentials.NewServerTLSFromFile("E:\\workspaces\\goland\\grpc-study\\key\\public.pem", "E:\\workspaces\\goland\\grpc-study\\key\\private.key")

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	//创建grpc
	grpcServer := grpc.NewServer(grpc.Creds(creds)) //创建gRPC服务，并把验证加入进来

	//grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	//注册grpc
	pb.RegisterSayHelloServer(grpcServer, &server{})

	//启动
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}

// hello server
type server struct {
	pb.UnimplementedSayHelloServer
}

// SayHello 直接在业务中处理，当然实际项目中，应该是在中间件中处理，业务传输层的校验
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("not token")
	}
	var appId string
	var appKey string
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}

	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}

	if appId != "milo" {
		return nil, errors.New("appid verify")
	}
	if appKey != "123456" {
		return nil, errors.New("appKey verify")
	}

	return &pb.HelloResponse{
		ResponseMsg: "hello " + req.RequestName,
	}, nil
}
