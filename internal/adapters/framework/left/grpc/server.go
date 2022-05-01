package rpc

import (
	"hexrpc/internal/adapters/framework/left/grpc/pb"
	"hexrpc/internal/ports"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{
		api: api,
	}
}

func (grpca Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("failed to listen on port 9000 ", err)

	}

	pignServiceServer := grpca

	grpcServer := grpc.NewServer()
	pb.RegisterPingServiceServer(grpcServer, pignServiceServer)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("failed to serve grpc over port 9000", err)
	}
}
