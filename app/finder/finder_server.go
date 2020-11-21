package finder

import (
	"context"
	"errors"
	"github.com/chernovsergey/dockerized_service/api"
)

type FinderServer struct {
	finder Finder
}

func (s *FinderServer) Init() {
	s.finder = Finder{}
}

func (s *FinderServer) BuildIndex(c context.Context, in *api.Text) (*api.Status, error) {
	err := s.finder.BuildIndex([]byte(in.GetValue()))
	return &api.Status{Value: 0}, err
}

func (s *FinderServer) Lookup(c context.Context, p *api.Pattern) (*api.Offsets, error) {

	//now := time.Now()
	//defer func() {
	//	metrics.ResponseLatency("make_lookup", time.Since(now).Seconds())
	//}()

	if s.finder.isEmpty() {
		return nil, errors.New("index is empty")
	}

	pbyte := []byte(p.GetValue())
	if len(pbyte) == 0 {
		return nil, errors.New("pattern is empty")
	}

	offsets := s.finder.Lookup(pbyte, -1)

	result := api.Offsets{Value: make([]int32, len(offsets))}
	for i := 0; i < len(offsets); i++ {
		result.Value[i] = int32(offsets[i])
	}
	return &result, nil
}
