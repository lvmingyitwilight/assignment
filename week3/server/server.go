package server

import (
	"assignment/week3/httpserver"
	"context"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	servers []*httpserver.Server
	cancel  func()
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Add(srv *httpserver.Server) {
	s.servers = append(s.servers, srv)
}

func (s *Server) Start() error {
	var ctx context.Context
	ctx, s.cancel = context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	for _, srv := range s.servers {
		eg.Go(func() error {
			return srv.Start()
		})

		eg.Go(func() error {
			<-ctx.Done()
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()
			return srv.Shutdown(ctx)
		})
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	eg.Go(func() error {
		select {
		case <-ch:
			s.Stop()
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	})
	return eg.Wait()
}

func (s *Server) Stop() {
	if s.cancel != nil {
		s.cancel()
	}
}
