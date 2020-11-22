package main

import (
	"context"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/chernovsergey/dockerized_service/api"
	"github.com/chernovsergey/dockerized_service/app/finder"
)

const (
	portGRPC = ":5050"
	portHTTP = ":7070"
)

func RunListener(port string) *net.Listener {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Cannot listen %s", port)
	}
	log.Printf("Listening port: %s", port)
	return &lis
}

func grpcStart(ctx context.Context, port string) error {
	listener := RunListener(port)

	// Configure main service logic
	finderServer := &finder.FinderServer{}
	finderServer.Init()

	// Create GRPC service and wrap main service in it
	grpcServer := grpc.NewServer(
		// provides Prometheus monitoring for Unary RPCs
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	api.RegisterFinderServer(grpcServer, finderServer)

	grpc_prometheus.EnableHandlingTimeHistogram(
		grpc_prometheus.WithHistogramBuckets([]float64{
			.001, .005, .01, .025, .05, .1,
		}),
	)
	grpc_prometheus.Register(grpcServer)

	log.Println("Starting GRPC server.", "Port:", port)
	select {
	case <-ctx.Done():
		grpcServer.GracefulStop()
		return ctx.Err()
	case err := <-ErrChannel(func() error { return grpcServer.Serve(*listener) }):
		return err
	}
}

func httpStart(ctx context.Context, port string) error {
	mux := http.DefaultServeMux
	mux.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Handler: mux,
		Addr:    port,
	}

	log.Println("Starting http server for Prometheus.", "Port:", port)
	e, _ := errgroup.WithContext(ctx)
	e.Go(func() error {
		listener := RunListener(port)
		return srv.Serve(*listener)
	})

	e.Go(func() error {
		<-ctx.Done()
		return srv.Shutdown(ctx)
	})

	return e.Wait()
}

func ErrChannel(in func() error) <-chan error {
	out := make(chan error)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		out <- in()
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errsGRPC := ErrChannel(func() error { return grpcStart(ctx, portGRPC) })
	errsHTTP := ErrChannel(func() error { return httpStart(ctx, portHTTP) })

	select {
	case reason := <-errsGRPC:
		log.Println("grpc server is down. reason = ", reason)
	case reason := <-errsHTTP:
		log.Println("http server is down. reason = ", reason)
	case <-ctx.Done():
		log.Println("context canceled. reason = ", ctx.Err())
	}

	<-ctx.Done()
}
