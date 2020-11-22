package main

import (
	"context"
	"flag"
	"log"
	"sync/atomic"
	"time"

	"github.com/chernovsergey/dockerized_service/api"
	"google.golang.org/grpc"
)

const (
	port    = "192.168.0.106:5050"
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
	//ctx, cancel := context.WithTimeout(context.Background(), timeout)
	ctx := context.Background()
	return client.Lookup(ctx, &api.Pattern{Value: pattern})
}

func main() {
	nWorkers := flag.Int("workers", 100, "-w")
	flag.Parse()

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

	done := make(chan int32, 100000)
	errs := make(chan int32, 100000)

	workers := *nWorkers
	log.Println("Using %v workers", workers)

	for i := 0; i < workers; i++ {
		go func() {
			for {
				if _, err := lookup(client, 200 * time.Millisecond, pattern); err == nil {
					done <- 1
				} else {
					errs <- 1
				}
			}
		}()
	}

	var success int32
	for i := 0; i < workers; i++ {
		go func() {
			for v := range done{
				atomic.AddInt32(&success, v)
			}
		}()
	}

	var errors int32
	for i := 0; i < workers; i++ {
		go func() {
			for v := range errs{
				atomic.AddInt32(&errors, v)
			}
		}()
	}

	select {
	case <-time.After(10 * time.Second):
		close(done)
		close(errs)
		log.Println("Done %v requests", success)
		log.Println("Errs %v requests", errors)
	}
}
