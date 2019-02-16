package main

import (
	"net"
	"log"
	"context"
)

type HelloServiceImpl struct{}

func (h *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
	reply := &String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func main() {
	grpcServer := grpc.NewServer()
	RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
