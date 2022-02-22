package service

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	db "Access/service/pb/judge"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

type UserInfoServiceServer struct {
	db.UnimplementedUserInfoServiceServer
}

func (s *UserInfoServiceServer) GetUserInfo(ctx context.Context, userRequest *db.UserRequest) (*db.UserResponse, error) {
	fmt.Println("Get from Join:", userRequest)
	return &db.UserResponse{
		Code: userRequest.Id,
		Mesg: "ok",
	}, nil
}

///judeg（判定服务）作为服务端（Server）开启连接
func RunJudge(judgeAddr string) {

	// lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	lis, err := net.Listen("tcp", fmt.Sprintf(judgeAddr))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	db.RegisterUserInfoServiceServer(grpcServer, &UserInfoServiceServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
