package client

import (
	"context"
	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	_ "google.golang.org/grpc"
)

// LoadEvents liuxd: DDD event storage get aggregate root by id
func (c *GRPCClient) LoadEvents(ctx context.Context, req *pb.LoadEventRequest) (*pb.LoadEventResponse, error) {
	return c.protoClient.LoadEvents(ctx, req)
}

// SaveSnapshot liuxd: DDD event storage save event snapshot
func (c *GRPCClient) SaveSnapshot(ctx context.Context, in *pb.SaveSnapshotRequest) (*pb.SaveSnapshotResponse, error) {
	return c.protoClient.SaveSnapshot(ctx, in)
}

// ExistAggregate liuxd: DDD event storage get event by id
func (c *GRPCClient) ExistAggregate(ctx context.Context, in *pb.ExistAggregateRequest) (*pb.ExistAggregateResponse, error) {
	return c.protoClient.ExistAggregate(ctx, in)
}

// ApplyEvent liuxd: DDD event storage apply event
func (c *GRPCClient) ApplyEvent(ctx context.Context, in *pb.ApplyEventRequest) (*pb.ApplyEventResponse, error) {
	return c.protoClient.ApplyEvent(ctx, in)
}
