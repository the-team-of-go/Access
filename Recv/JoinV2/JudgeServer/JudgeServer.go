package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "Recv/JoinV2"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50052, "The server port")
)

type ReportServer struct {
	pb.UnimplementedReportServer
}

func (s *ReportServer) Send(ctx context.Context, reportReq *pb.ReportReq) (*pb.ReportRsp, error) {

	fmt.Println("ad %v", reportReq)

	var c int32 = 1
	var m string = "m"
	return &pb.ReportRsp{
		Code: c,
		Msg:  m,
	}, nil

}

///judeg（判定服务）作为服务端（Server）开启连接
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterReportServer(grpcServer, &ReportServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
