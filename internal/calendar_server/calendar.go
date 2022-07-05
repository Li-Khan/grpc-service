package calendar_server

import (
	"context"
	"github.com/Li-Khan/grpc-service/api/protobuf/calendar"
)

type CalendarServer struct {
}

func (c CalendarServer) Add(ctx context.Context, event *calendar.Event) (*calendar.Events, error) {
	//TODO implement me
	panic("implement me")
}

func (c CalendarServer) Update(ctx context.Context, event *calendar.Event) (*calendar.Events, error) {
	//TODO implement me
	panic("implement me")
}

func (c CalendarServer) List(ctx context.Context, event *calendar.Event) (*calendar.Events, error) {
	//TODO implement me
	panic("implement me")
}

func (c CalendarServer) Delete(ctx context.Context, event *calendar.Event) (*calendar.Events, error) {
	//TODO implement me
	panic("implement me")
}
