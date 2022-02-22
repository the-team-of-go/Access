package service

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/robfig/cron"
	"google.golang.org/grpc"

	pb "Access/service/pb/agent"
	"Access/service/pb/gopool"
	db "Access/service/pb/judge"
	"Access/service/pb/queue"
)

type ReportServer struct {
	pb.UnimplementedReportServer
}

var ReportReqStream []*pb.ReportReq
var ReportReqsq = queue.NewSliceQueue(0)

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

	//fmt.Println("ad %v", reportReq)

	//ReportReqStream = append(ReportReqStream, reportReq)
	ReportReqsq.Enqueue(reportReq)

	ind++
	fmt.Println("Agent send ", ind, " ", reportReq)

	return &pb.ReportRsp{
		Code: reportReq.MachineId,
		Msg:  "ok",
	}, nil

}

func RunGetUserInfo() {

	if ReportReqsq.Getqsize() > 0 {

		pool := gopool.New(PoolSize)

		conn, err := grpc.Dial(JudgeAddr, grpc.WithInsecure())
		if err != nil {
			log.Println("fail to dial: ", err)
		}
		defer conn.Close()
		client := db.NewUserInfoServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		for ReportReqsq.Getqsize() > 0 {
			fmt.Println("  ind is ", ind, " qsize is ", ReportReqsq.Getqsize())
			RRsq := ReportReqsq.Dequeue()
			pool.Add(1)
			go func(RRsq *pb.ReportReq) {
				userinfo := &db.UserRequest{
					Id:       RRsq.MachineId,
					Cpuused:  RRsq.CpuPercent,
					Memused:  RRsq.MemPercent,
					DiskUsed: RRsq.DiskPercent,
					Timeout:  RRsq.TimeStamp,
				}
				feature, err := client.GetUserInfo(ctx, userinfo)
				if err != nil {
					log.Println("%v.GetUserInfo(_) = _, %v: ", client, err)
				}

				log.Println(feature)
				pool.Done()
			}(RRsq)
		}

		pool.Wait()
	}
}

func RunJointToJudge(AgentAddr string) {

	c := cron.New()                            // 新建一个定时任务对象
	c.AddFunc("*/3 * * * * ?", RunGetUserInfo) // 给对象增加定时任务
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

// func RunGetUserInfo() {
// 	if ind < len(ReportReqStream) {

// 		pool := gopool.New(PoolSize)

// 		conn, err := grpc.Dial(JudgeAddr, grpc.WithInsecure())
// 		if err != nil {
// 			log.Fatalf("fail to dial: %v", err)
// 		}
// 		defer conn.Close()
// 		client := db.NewUserInfoServiceClient(conn)
// 		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 		defer cancel()

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
