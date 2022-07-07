package server

import (
	"context"
	pb "github.com/Li-Khan/grpc-service/api/protobuf/calendar"
	"github.com/Li-Khan/grpc-service/internal/helper"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type PbServer struct {
	pb.UnimplementedCalendarServer
	db *pgxpool.Pool
}

const format string = "2006-01-02"

func NewCalendarServer(db *pgxpool.Pool) *PbServer {
	return &PbServer{
		db: db,
	}
}

func (c PbServer) Add(ctx context.Context, event *pb.Event) (*pb.Event, error) {
	date, err := time.Parse(format, event.Date.AsTime().Format(format))
	if err != nil {
		return nil, err
	}

	stmt := `
	INSERT INTO "event" (
		"name",
		"date"
	) VALUES ($1, $2) RETURNING "id"`

	err = c.db.QueryRow(ctx, stmt, event.Name, date).Scan(&event.Id)

	if err != nil {
		if err.(*pgconn.PgError).Code == "23505" {
			return nil, helper.ErrAlreadyExist
		}
		return nil, err
	}

	return event, nil
}

func (c PbServer) Update(ctx context.Context, event *pb.Event) (*pb.Event, error) {
	stmt := `
		UPDATE "event"
		SET "name" = $1, "date" = $2
		WHERE "id" = $3;`

	_, err := c.db.Exec(ctx, stmt, event.Name, event.Date.AsTime(), event.Id)
	if err != nil {
		if err.(*pgconn.PgError).Code == "23505" {
			return nil, helper.ErrAlreadyExist
		}
		return nil, err
	}

	return event, nil
}

func (c PbServer) GetByID(ctx context.Context, request *pb.GetEventByIDRequest) (*pb.Event, error) {
	stmt := `
		SELECT * 
		FROM "event"
		WHERE "id" = $1
	`

	event := pb.Event{}
	row := c.db.QueryRow(ctx, stmt, request.Id)
	t := time.Time{}
	err := row.Scan(&event.Id, &event.Name, &t)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, helper.ErrEventNotFound
		}
		return nil, err
	}
	event.Date = timestamppb.New(t)

	return &event, nil
}

func (c PbServer) List(request *pb.ListEventsRequest, server pb.Calendar_ListServer) error {
	stmt := `SELECT * FROM "event"`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := c.db.Query(ctx, stmt)
	defer rows.Close()
	if err != nil {
		return err
	}

	date := time.Time{}
	for rows.Next() {
		event := pb.Event{}
		err = rows.Scan(&event.Id, &event.Name, &date)
		event.Date = timestamppb.New(date)

		if err != nil {
			return err
		}

		err = server.Send(&event)
		if err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func (c PbServer) Delete(ctx context.Context, request *pb.DeleteEventRequest) (*pb.Event, error) {
	stmt := `DELETE FROM "event" WHERE "id" = $1`

	event, err := c.GetByID(ctx, &pb.GetEventByIDRequest{Id: request.Id})
	if err != nil {
		return nil, err
	}

	_, err = c.db.Exec(ctx, stmt, request.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, helper.ErrEventNotFound
		}
		return nil, err
	}

	return event, nil
}
