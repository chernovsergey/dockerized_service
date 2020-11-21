package main

import (
	"context"
	"log"
	"time"

	"github.com/chernovsergey/dockerized_service/api"
	"google.golang.org/grpc"
)

const (
	port    = "localhost:5050"
	text    = "banana"
	pattern = "ana"
)

func precalc(client api.FinderClient, timeout time.Duration,
	text string) (*api.Status, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return client.BuildIndex(ctx, &api.Text{Value: text})
}

func lookup(client api.FinderClient, timeout time.Duration,
	pattern string) (*api.Offsets, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return client.Lookup(ctx, &api.Pattern{Value: pattern})
}

func main() {

	// Create connection to grpc server, working on
	// specific adress
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect to %s", port)
	}
	defer conn.Close()

	// Create service client instance
	client := api.NewFinderClient(conn)

	// Build index
	_, err = precalc(client, time.Second, text)
	if err != nil {
		log.Fatalf("BuildIndex failed. Reason = %v ", err)
	}

	// Estimate rps
	log.Println("start")
	success := 0
	go func() {
		for {
			if _, err := lookup(client, time.Second, pattern); err == nil {
				success++
			}
		}
	}()

	select {
	case <-time.After(10 * time.Second):
		log.Println("stop")
		log.Printf("Estimated requests per second = %v", success)
	}
}
