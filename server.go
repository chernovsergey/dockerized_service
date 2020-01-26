package main

import (
	"context"
	"errors"
	"index/suffixarray"
	"log"
	"net"
	"sync"

	"github.com/chernovsergey/dockerized_service/api"
	"google.golang.org/grpc"
)

const (
	port = ":50505"
)

// revive:disable:exported

type Finder struct {
	arr *suffixarray.Index
}

func (s *Finder) BuildIndex(c context.Context,
	in *api.Text) (*api.Status, error) {
	s.arr = suffixarray.New([]byte(in.GetValue()))
	return &api.Status{Value: 0}, nil
}

func (s *Finder) Lookup(c context.Context,
	p *api.Pattern) (*api.Offsets, error) {

	if s.arr == nil {
		return nil, errors.New("index is empty")
	}

	pbyte := []byte(p.GetValue())
	if len(pbyte) == 0 {
		return nil, errors.New("patter is empty")
	}

	offsets := s.arr.Lookup(pbyte, -1)

	result := api.Offsets{Value: make([]int32, len(offsets))}
	for i := 0; i < len(offsets); i++ {
		result.Value[i] = int32(offsets[i])
	}
	return &result, nil
}

func RunListener(port string) *net.Listener {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Panicln(err)
		log.Fatalf("Cannot listen %s", port)
	}
	log.Printf("Listening port: %s", port)
	return &lis
}

func grpcStart(ctx context.Context, port string) error {
	listener := RunListener(port)

	instance := &Finder{}
	server := grpc.NewServer()
	api.RegisterFinderServer(server, instance)
	select {
	case <-ctx.Done():
		server.GracefulStop()
		return ctx.Err()
	case err := <-Errch(func() error { return server.Serve(*listener) }):
		return err
	}
}

func Errch(in func() error) <-chan error {
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

	errors := Errch(func() error { return grpcStart(ctx, port) })

	select {
	case reason := <-errors:
		log.Println("grpc server is down. reason = ", reason)
	case <-ctx.Done():
		log.Println("context canceled. reason = ", ctx.Err())
	}

	<-ctx.Done()
}
