package router

import (
	grpcHandler "github.com/Mizumaki/Go_and_gRPC_sample/server/src/handler/grpc"
	pb "github.com/Mizumaki/Go_and_gRPC_sample/server/src/handler/grpc/proto/user"
	"github.com/Mizumaki/Go_and_gRPC_sample/server/src/registry"
	"google.golang.org/grpc"
)

func Register(server *grpc.Server) error {
	container, err := registry.New()

	if err != nil {
		return err
	}

	container.Invoke(func(handler grpcHandler.UserHandler) {
		pb.RegisterUserInfoAPIServer(server, handler)
	})

	return nil
}
