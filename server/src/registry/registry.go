package registry

import (
	"github.com/pkg/errors"

	grpcHandler "github.com/Mizumaki/Go_and_gRPC_sample/server/src/handler/grpc"
	"github.com/Mizumaki/Go_and_gRPC_sample/server/src/infrastructure/mock"
	"go.uber.org/dig"
)

func New() (*dig.Container, error) {
	container := dig.New()
	
	if err := container.Provide(mock.NewUserMock); err != nil {
		return nil, errors.WithStack(err)
	}
	if err := container.Provide(grpcHandler.NewUserHandler); err != nil {
		return nil, errors.WithStack(err)
	}

	return container, nil
}
