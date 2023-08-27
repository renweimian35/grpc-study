package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	pb "grpc-study/hello-server/proto"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	//创建grpc
	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
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

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("not token")
	}
	fmt.Println(md)
	var appId string
	var appKey string
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}

	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}

	fmt.Println("appid = ", appId)
	fmt.Println("appKey = ", appKey)

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
