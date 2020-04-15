package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/Mizumaki/Go_and_gRPC_sample/server/src/handler/grpc/proto/user"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("go_and_grpc:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewUserInfoAPIClient(conn)
	request := &pb.UserBasicInfoRequest{Email: "hello", Password: "world"}
	res, err := client.GetUserBasicInfo(context.TODO(), request)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
