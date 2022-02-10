package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "Recv/JoinV2"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

var (
	serverAddr = flag.String("addr", "localhost:50052", "The server address in the format of host:port")
)

type ReportServer struct {
	pb.UnimplementedReportServer
}

func runJointToJudge(reportReq *pb.ReportReq) (*pb.ReportRsp, error) {

	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewReportClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feature, err := client.Send(ctx, reportReq)
	if err != nil {
		log.Fatalf("%v.Send(_) = _, %v: ", client, err)
	}
	log.Println(feature)

	return feature, err
}

//等待Agent端将信息传输过来后，即把信息通过runRecordRouteRAS传输到判定服务（judge）
func (s *ReportServer) Send(ctx context.Context, reportReq *pb.ReportReq) (*pb.ReportRsp, error) {

	fmt.Println("ad %v", reportReq)

	feature, err := runJointToJudge(reportReq)

	return feature, err

}

///开启连接
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	///////

	//s := grpc.NewServer()
	pb.RegisterReportServer(grpcServer, &ReportServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
