package mock

import (
	"context"

	"github.com/Mizumaki/Go_and_gRPC_sample/server/src/domain/entity"
	"github.com/Mizumaki/Go_and_gRPC_sample/server/src/domain/repository"
)

type userMock struct{}

func (*userMock) Get(ctx context.Context, email string) (*entity.User, error) {
	return &entity.User{
		ID:    "userID",
		Email: "user@gmail.com",
	}, nil
}

func NewUserMock() repository.User {
	return &userMock{}
}
