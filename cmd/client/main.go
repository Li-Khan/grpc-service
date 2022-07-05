package main

import (
	"context"
	"fmt"
	"github.com/Li-Khan/grpc-service/api/protobuf/calendar"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	defer func() {
		_ = conn.Close()
	}()
	if err != nil {
		log.Println(err)
		return
	}

	c := calendar.NewCalendarClient(conn)

	events, err := c.Add(context.Background(), &calendar.Event{
		Name: "aboba",
		Date: timestamppb.Now(),
	})
	fmt.Println(events.Events)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(events)
}
