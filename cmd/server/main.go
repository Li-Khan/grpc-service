package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/Li-Khan/grpc-service/api/protobuf/calendar"
	"github.com/Li-Khan/grpc-service/configs"
	"github.com/Li-Khan/grpc-service/internal/calendar/server"
	"github.com/Li-Khan/grpc-service/internal/repository/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	db := postgres.GetDb()
	err := postgres.Migration(db)
	if err != nil {
		panic(err)
	}
}

func main() {
	address := fmt.Sprintf("localhost:%d", configs.GetConfig().BindAddr)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}

	srv := grpc.NewServer()
	reflection.Register(srv)

	calendarServer := server.NewCalendarServer(postgres.GetDb())

	pb.RegisterCalendarServer(srv, calendarServer)

	log.Printf("starting the grpc server on :%d\n", configs.GetConfig().BindAddr)
	log.Println(srv.Serve(listener))
}
