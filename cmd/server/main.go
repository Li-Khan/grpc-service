package main

import (
	"context"
	"github.com/Li-Khan/grpc-service/internal/server/calendar"
	"log"
	"net"

	pb "github.com/Li-Khan/grpc-service/api/protobuf/calendar"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type a struct {
	pb.UnimplementedCalendarServer
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Println(err)
		return
	}

	server := grpc.NewServer()
	reflection.Register(server)

	calendarServer := calendar.NewCalendarServer()

	pb.RegisterCalendarServer(server, calendarServer)
	calendarServer.Add(context.Background(), &pb.Event{})

	log.Println(server.Serve(listener))
}
