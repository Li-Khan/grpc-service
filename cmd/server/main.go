package main

import (
	"google.golang.org/grpc"
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

	log.Println(server.Serve(listener))
}
