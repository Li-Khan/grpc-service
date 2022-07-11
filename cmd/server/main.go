package main

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net"
	"sync"

	pb "github.com/Li-Khan/grpc-service/api/protobuf/calendar"
	"github.com/Li-Khan/grpc-service/configs"
	"github.com/Li-Khan/grpc-service/internal/calendar/server"
	"github.com/Li-Khan/grpc-service/internal/repository/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const cfgPath string = "./configs/config.json"

var db *pgxpool.Pool
var cfg *configs.Config
var onceDb sync.Once
var onceCfg sync.Once

func getDb() *pgxpool.Pool {
	onceDb.Do(func() {
		var err error
		cfg = getConfig()
		db, err = postgres.NewPostgresRepository(cfg)
		if err != nil {
			panic(err)
		}
	})

	return db
}

func getConfig() *configs.Config {
	onceCfg.Do(func() {
		var err error

		cfg, err = configs.LoadConfig(cfgPath)
		if err != nil {
			panic(err)
		}
	})
	return cfg
}

func init() {
	db = getDb()
	//	TODO migration
}

func main() {
	address := fmt.Sprintf("localhost:%d", getConfig().BindAddr)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}

	srv := grpc.NewServer()
	reflection.Register(srv)

	calendarServer := server.NewCalendarServer(getDb())

	pb.RegisterCalendarServer(srv, calendarServer)

	log.Printf("starting the grpc server on :%d\n", getConfig().BindAddr)
	log.Println(srv.Serve(listener))
}
