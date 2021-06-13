package httpserver

import (
	"context"
	"net"
	"net/http"
)

type Server struct {
	*http.Server
	options *serverOptions
}

type serverOptions struct {
	network string
	address string
}

type ServerOption func(*serverOptions)

func WithNetwork(network string) ServerOption {
	return func(o *serverOptions) {
		o.network = network
	}
}

func WithAddress(addr string) ServerOption {
	return func(o *serverOptions) {
		o.address = addr
	}
}

func NewHttpServer(options ...ServerOption) *Server {
	opts := serverOptions{
		network: "tcp",
		address: ":8080",
	}
	for _, o := range options {
		o(&opts)
	}
	mux := http.NewServeMux()
	return &Server{Server: &http.Server{Handler: mux}, options: &opts}
}

func (s *Server) Start() error {
	listener, err := net.Listen(s.options.network, s.options.address)
	if err != nil {
		return err
	}
	return s.Serve(listener)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.Shutdown(ctx)
}
