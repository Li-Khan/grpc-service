package main

import (
	"github.com/Li-Khan/grpc-service/api/protobuf/calendar"
	"github.com/Li-Khan/grpc-service/internal/calendar_server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Println(err)
		return
	}

	server := grpc.NewServer()
	reflection.Register(server)

	calendar.RegisterCalendarServer(server, calendar_server.CalendarServer{})

	log.Println(server.Serve(listener))
}
