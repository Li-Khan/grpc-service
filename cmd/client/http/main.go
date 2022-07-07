package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	pb "github.com/Li-Khan/grpc-service/api/protobuf/calendar"
	"github.com/Li-Khan/grpc-service/configs"
	"github.com/Li-Khan/grpc-service/internal/http_handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfgPath := flag.String("c", "./configs/config.json", "path to the config file")
	flag.Parse()

	cfg, err := configs.LoadConfig(*cfgPath)
	if err != nil {
		log.Println(err)
		return
	}

	address := fmt.Sprintf("localhost:%d", cfg.BindAddr)

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer func() {
		_ = conn.Close()
	}()
	if err != nil {
		log.Println(err)
		return
	}

	c := pb.NewCalendarClient(conn)

	handler := http_handler.NewHandler(c)

	address = fmt.Sprintf("localhost:%d", cfg.ClientAddr)
	srv := &http.Server{
		Addr:           address,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Printf("starting the client on :%d\n", cfg.ClientAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
			os.Exit(0)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	timeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(timeout); err != nil {
		log.Println(err)
		return
	}
}
