package week3

import (
	"assignment/week3/httpserver"
	"assignment/week3/server"
	"log"
)

func main() {
	srv := server.NewServer()
	srv.Add(httpserver.NewHttpServer(httpserver.WithAddress(":1234")))
	srv.Add(httpserver.NewHttpServer(httpserver.WithAddress(":1235")))
	if err := srv.Start(); err != nil {
		log.Panic(err)
	}
}
