package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/robfig/cron"
	"google.golang.org/grpc"

	pb "Recv/service/pb/agent"
	"Recv/service/pb/gopool"
	db "Recv/service/pb/judge"
)

// var (
// 	port = flag.Int("port", 50051, "The server port")
// )

// var (
// 	serverAddr = flag.String("addr", "localhost:50052", "The server address in the format of host:port")
// )

type ReportServer struct {
	pb.UnimplementedReportServer
}

var ReportReqStream []*pb.ReportReq
var ind int = 0
var PoolSize int
var JudgeAddr string

func SetUserInfoServiceAddr(judgeAddr string) {
	JudgeAddr = judgeAddr
	fmt.Println("JudgeAddr setted: ", JudgeAddr)
}

func SetPoolSize(poolSize int) {
	PoolSize = poolSize
	fmt.Println("PoolSize setted: ", PoolSize)
}

//等待Agent端将信息传输过来后，即把信息通过runRecordRouteRAS传输到判定服务（judge）
func (s *ReportServer) Send(ctx context.Context, reportReq *pb.ReportReq) (*pb.ReportRsp, error) {

	fmt.Println("ad %v", reportReq)

	ReportReqStream = append(ReportReqStream, reportReq)
	for i, j := range ReportReqStream {
		fmt.Println("Agent send ", i, " ", j)
	}

	return &pb.ReportRsp{
		Code: reportReq.MachineId,
		Msg:  "ok",
	}, nil

}

// func RunGetUserInfo() {
// 	if ind < len(ReportReqStream) {
// 		pool := gopool.New(PoolSize)
// 		for ; ind < len(ReportReqStream); ind++ {
// 			pool.Add(1)
// 			go func(ind int) {
// 				userinfo := &db.UserRequest{
// 					Id:       ReportReqStream[ind].MachineId,
// 					Cpuused:  ReportReqStream[ind].CpuPercent,
// 					Memused:  ReportReqStream[ind].MemPercent,
// 					DiskUsed: ReportReqStream[ind].DiskPercent,
// 					Timeout:  ReportReqStream[ind].TimeStamp,
// 				}
// 				// conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())
// 				conn, err := grpc.Dial(JudgeAddr, grpc.WithInsecure())
// 				if err != nil {
// 					log.Fatalf("fail to dial: %v", err)
// 				}
// 				defer conn.Close()
// 				client := db.NewUserInfoServiceClient(conn)
// 				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 				defer cancel()
// 				feature, err := client.GetUserInfo(ctx, userinfo)
// 				if err != nil {
// 					log.Fatalf("%v.Send(_) = _, %v: ", client, err)
// 				}
// 				log.Println(feature)
// 				pool.Done()
// 			}(ind)
// 		}
// 		pool.Wait()
// 	}
// }

func RunGetUserInfo() {
	if ind < len(ReportReqStream) {

		pool := gopool.New(PoolSize)

		conn, err := grpc.Dial(JudgeAddr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()
		client := db.NewUserInfoServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		for ; ind < len(ReportReqStream); ind++ {
			pool.Add(1)
			go func(ind int) {
				userinfo := &db.UserRequest{
					Id:       ReportReqStream[ind].MachineId,
					Cpuused:  ReportReqStream[ind].CpuPercent,
					Memused:  ReportReqStream[ind].MemPercent,
					DiskUsed: ReportReqStream[ind].DiskPercent,
					Timeout:  ReportReqStream[ind].TimeStamp,
				}
				// conn, err := grpc.Dial("127.0.0.1:50052", grpc.WithInsecure())

				feature, err := client.GetUserInfo(ctx, userinfo)
				if err != nil {
					log.Fatalf("%v.Send(_) = _, %v: ", client, err)
				}
				log.Println(feature)
				pool.Done()
			}(ind)
		}
		pool.Wait()
	}
}

func RunJointToJudge(AgentAddr string) {

	c := cron.New()                             // 新建一个定时任务对象
	c.AddFunc("*/10 * * * * ?", RunGetUserInfo) // 给对象增加定时任务
	c.Start()

	// lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:50051"))
	lis, err := net.Listen("tcp", fmt.Sprintf(AgentAddr))
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

///开启连接
// func main() {
// 	flag.Parse()

// 	var hostaddr string = "127.0.0.1:50051"
// 	//var aport int = 50051
// 	// lis, err := net.Listen("tcp", fmt.Sprintf("10.243.55.132:%d", *port))
// 	// lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", aport))
// 	lis, err := net.Listen("tcp", fmt.Sprintf(hostaddr))
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	grpcServer := grpc.NewServer()
// 	///////

// 	//s := grpc.NewServer()
// 	pb.RegisterReportServer(grpcServer, &ReportServer{})

// 	log.Printf("server listening at %v", lis.Addr())
// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}

// }
