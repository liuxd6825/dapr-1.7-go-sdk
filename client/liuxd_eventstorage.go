package client

import (
	"context"
	pb "github.com/liuxd6825/dapr/pkg/proto/runtime/v1"
	_ "google.golang.org/grpc"
)

//
// LoadEvents
// @Description:  DDD event storage get aggregate root by id
// @receiver c
// @param ctx
// @param req
// @return *pb.LoadEventResponse
// @return error
//
func (c *GRPCClient) LoadEvents(ctx context.Context, req *pb.LoadEventRequest) (*pb.LoadEventResponse, error) {
	return c.protoClient.LoadEvents(c.withAuthToken(ctx), req)
}

//
// SaveSnapshot
// @Description: liuxd: DDD event storage apply event
// @receiver c
// @param ctx
// @param in
// @return *pb.SaveSnapshotResponse
// @return error
//
func (c *GRPCClient) SaveSnapshot(ctx context.Context, in *pb.SaveSnapshotRequest) (*pb.SaveSnapshotResponse, error) {
	return c.protoClient.SaveSnapshot(c.withAuthToken(ctx), in)
}

//
// ApplyEvent
// @Description: liuxd: DDD event storage apply event
// @receiver c
// @param ctx
// @param in
// @return *pb.ApplyEventsResponse
// @return error
//
func (c *GRPCClient) ApplyEvent(ctx context.Context, in *pb.ApplyEventRequest) (*pb.ApplyEventResponse, error) {
	return c.protoClient.ApplyEvent(c.withAuthToken(ctx), in)
}

//
// CreateEvent
// @Description:
// @receiver c
// @param ctx
// @param request
// @return *pb.CreateEventResponse
// @return error
//
func (c *GRPCClient) CreateEvent(ctx context.Context, request *pb.CreateEventRequest) (*pb.CreateEventResponse, error) {
	return c.protoClient.CreateEvent(ctx, request)
}

//
// DeleteEvent
// @Description:
// @receiver c
// @param ctx
// @param request
// @return *pb.DeleteEventResponse
// @return error
//
func (c *GRPCClient) DeleteEvent(ctx context.Context, request *pb.DeleteEventRequest) (*pb.DeleteEventResponse, error) {
	return c.protoClient.DeleteEvent(ctx, request)
}

func (c *GRPCClient) GetRelations(ctx context.Context, request *pb.GetRelationsRequest) (*pb.GetRelationsResponse, error) {
	return c.protoClient.GetRelations(ctx, request)
}
