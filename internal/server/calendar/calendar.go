package calendar

import (
	"context"
	pb "github.com/Li-Khan/grpc-service/api/protobuf/calendar"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PbServer struct {
	pb.UnimplementedCalendarServer
	db *pgxpool.Pool
}

func NewCalendarServer(db *pgxpool.Pool) *PbServer {
	return &PbServer{
		db: db,
	}
}

func (c PbServer) Add(ctx context.Context, event *pb.Event) (*pb.Event, error) {
	stmt := `INSERT INTO "event" (
		"name",
		"date"
	) VALUES ($1, $2) RETURNING "id"`

	err := c.db.QueryRow(ctx, stmt, event.Name, event.Date.AsTime()).Scan(&event.Id)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (c PbServer) Update(ctx context.Context, request *pb.UpdateEventRequest) (*pb.Event, error) {
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
