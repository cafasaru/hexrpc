package ports

import (
	"context"
	"hexrpc/internal/adapters/framework/left/grpc/pb"
)

type GRPCPort interface {
	Run()
	Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error)
}
