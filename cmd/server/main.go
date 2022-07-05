package main

import (
	"flag"
	"fmt"
	"github.com/Li-Khan/grpc-service/internal/repository/postgres"
	"log"
	"net"

	pb "github.com/Li-Khan/grpc-service/api/protobuf/calendar"
	"github.com/Li-Khan/grpc-service/configs"
	"github.com/Li-Khan/grpc-service/internal/server/calendar"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfgPath := flag.String("c", "./configs/config.json", "path to the config file")
	flag.Parse()

	cfg, err := configs.LoadConfig(*cfgPath)
	if err != nil {
		log.Println(err)
		return
	}

	db, err := postgres.NewPostgresRepository(cfg)
	if err != nil {
		log.Println(err)
		return
	}

	address := fmt.Sprintf("localhost:%d", cfg.BindAddr)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}

	server := grpc.NewServer()
	reflection.Register(server)

	calendarServer := calendar.NewCalendarServer(db)

	pb.RegisterCalendarServer(server, calendarServer)

	log.Printf("starting the grpc server on :%d\n", cfg.BindAddr)
	log.Println(server.Serve(listener))
}
