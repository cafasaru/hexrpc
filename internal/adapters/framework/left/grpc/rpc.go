package rpc

import (
	"context"
	"hexrpc/internal/adapters/framework/left/grpc/pb"
	"log"
)

func (grpca Adapter) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	message, err := grpca.api.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("IN GRPC PING FUNC")

	return &pb.PingResponse{
		Message: message,
	}, nil
}
