package rpc

import (
	"context"
	"hexrpc/internal/adapters/framework/left/grpc/pb"
)

func (grpca Adapter) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	message, err := grpca.api.Ping()
	if err != nil {
		return nil, err
	}

	return &pb.PingResponse{
		Message: message,
	},nil
}