package client

import (
	"context"
	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
)

func (c *GRPCClient) WriteEventLog(ctx context.Context, in *pb.WriteEventLogRequest) (*pb.WriteEventLogResponse, error) {
	return c.protoClient.WriteEventLog(ctx, in)
}

func (c *GRPCClient) UpdateEventLog(ctx context.Context, in *pb.UpdateEventLogRequest) (*pb.UpdateEventLogResponse, error) {
	return c.protoClient.UpdateEventLog(ctx, in)
}

func (c *GRPCClient) GetEventLogByCommandId(ctx context.Context, in *pb.GetEventLogByCommandIdRequest) (*pb.GetEventLogByCommandIdResponse, error) {
	return c.protoClient.GetEventLogByCommandId(ctx, in)
}

func (c *GRPCClient) WriteAppLog(ctx context.Context, in *pb.WriteAppLogRequest) (*pb.WriteAppLogResponse, error) {
	return c.protoClient.WriteAppLog(ctx, in)
}

func (c *GRPCClient) UpdateAppLog(ctx context.Context, in *pb.UpdateAppLogRequest) (*pb.UpdateAppLogResponse, error) {
	return c.protoClient.UpdateAppLog(ctx, in)
}

func (c *GRPCClient) GetAppLogById(ctx context.Context, in *pb.GetAppLogByIdRequest) (*pb.GetAppLogByIdResponse, error) {
	return c.protoClient.GetAppLogById(ctx, in)
}
