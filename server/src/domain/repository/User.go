package repository

import (
	"context"

	"github.com/Mizumaki/Go_and_gRPC_sample/server/src/domain/entity"
)

type User interface {
	Get(ctx context.Context, email string) (*entity.User, error)
}
