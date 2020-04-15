package main

import (
	"log"
	"net"
	"os"

	grpcServer "github.com/Mizumaki/Go_and_gRPC_sample/server/src/interfaces/grpcServer"
)

func main() {
	PORT := ":" + os.Getenv("PORT")
	log.Printf("listening on: tcp" + PORT)
	listenPort, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalln(err)
	}
	server, err := grpcServer.New()
	if err == nil {
		server.Serve(listenPort)
	}
}
