package calendar

import (
	"context"
	pb "github.com/Li-Khan/grpc-service/api/protobuf/calendar"
)

type PbServer struct {
	pb.UnimplementedCalendarServer
}

func NewCalendarServer() *PbServer {
	return &PbServer{}
}

func (c PbServer) Add(ctx context.Context, event *pb.Event) (*pb.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (c PbServer) Update(ctx context.Context, event *pb.Event) (*pb.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (c PbServer) GetByID(ctx context.Context, request *pb.GetEventByIDRequest) (*pb.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (c PbServer) List(request *pb.ListEventsRequest, server pb.Calendar_ListServer) error {
	//TODO implement me
	panic("implement me")
}

func (c PbServer) Delete(ctx context.Context, request *pb.DeleteEventRequest) (*pb.Event, error) {
	//TODO implement me
	panic("implement me")
}
