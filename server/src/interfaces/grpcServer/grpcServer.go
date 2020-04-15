package grpcserver

import (
	"github.com/Mizumaki/Go_and_gRPC_sample/server/src/interfaces/grpcServer/router"
	"google.golang.org/grpc"
)

func New() (*grpc.Server, error) {
	server := grpc.NewServer()
	err := router.Register(server)
	if err != nil {
		return nil, err
	}

	return server, nil
}
