package client

import (
	"context"
	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	_ "google.golang.org/grpc"
)

// LoadEvents
// DDD event storage get aggregate root by id
func (c *GRPCClient) LoadEvents(ctx context.Context, req *pb.LoadEventRequest) (*pb.LoadEventResponse, error) {
	return c.protoClient.LoadEvents(c.withAuthToken(ctx), req)
}

// SaveSnapshot liuxd: DDD event storage save event snapshot
func (c *GRPCClient) SaveSnapshot(ctx context.Context, in *pb.SaveSnapshotRequest) (*pb.SaveSnapshotResponse, error) {
	return c.protoClient.SaveSnapshot(c.withAuthToken(ctx), in)
}

// ExistAggregate liuxd: DDD event storage get event by id
func (c *GRPCClient) ExistAggregate(ctx context.Context, in *pb.ExistAggregateRequest) (*pb.ExistAggregateResponse, error) {
	return c.protoClient.ExistAggregate(c.withAuthToken(ctx), in)
}

// ApplyEvent liuxd: DDD event storage apply event
func (c *GRPCClient) ApplyEvent(ctx context.Context, in *pb.ApplyEventRequest) (*pb.ApplyEventResponse, error) {
	return c.protoClient.ApplyEvent(c.withAuthToken(ctx), in)
}
