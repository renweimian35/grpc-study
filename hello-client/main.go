package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "grpc-study/hello-client/proto"
	"log"
)

func main() {
	//连接到server端

	creds, _ := credentials.NewClientTLSFromFile("E:\\workspaces\\goland\\grpc-study\\key\\public.pem", "www.baidu.club")

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(creds))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	conn, err := grpc.Dial("127.0.0.1:8080", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewSayHelloClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "miloyang"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.GetResponseMsg())
}

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"APPID":  "milo",
		"APPKEY": "123456",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return true
}
