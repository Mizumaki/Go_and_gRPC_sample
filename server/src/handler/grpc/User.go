package grpchandler

import (
	"context"
	"log"

	"github.com/Mizumaki/Go_and_gRPC_sample/server/src/domain/repository"
	pb "github.com/Mizumaki/Go_and_gRPC_sample/server/src/handler/grpc/proto/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler interface {
	GetUserBasicInfo(context.Context, *pb.UserBasicInfoRequest) (*pb.UserBasicInfo, error)
}

type userHandler struct {
	repository repository.User
}

func (handler *userHandler) GetUserBasicInfo(ctx context.Context, message *pb.UserBasicInfoRequest) (*pb.UserBasicInfo, error) {
	repository := handler.repository
	response, err := repository.Get(context.TODO(), message.Email)
	if err != nil {
		log.Printf("%+v", err)
		err := status.Error(codes.NotFound, err.Error())
		return nil, err
	}
	log.Printf("Called grpc service and respond it as")
	log.Printf("%+v", response)
	return &pb.UserBasicInfo{
		Id:                response.ID,
		Email:             response.Email,
	}, nil
}

func NewUserHandler(repo repository.User) UserHandler {
	return &userHandler{
		repository: repo,
	}
}
