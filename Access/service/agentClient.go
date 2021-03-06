package service

import (
	"context"
	"flag"
	"log"
	"time"

	pb "Access/service/pb/agent"
	"Access/service/pb/gopool"

	"google.golang.org/grpc"
)

//传输Agent信息
func RunAgentToJoin(agentAddr string) {

	flag.Parse()

	// conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	conn, err := grpc.Dial(agentAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewReportClient(conn)

	//控制goroutine数量
	pool := gopool.New(1)
	var i int32
	for i = 1; i <= 6; i++ {

		pool.Add(1)
		//go func
		go func(i int32) {
			//runAgentToJoin传输Agent信息
			reportReq := randomReportReq(i)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			feature, err := client.Send(ctx, reportReq)
			if err != nil {
				log.Fatalf("%v.Send(_) = _, %v: ", client, err)
			}
			log.Println(feature)
			pool.Done()
		}(i)
	}
	pool.Wait()

}

//randomInfoStream用于生成信息进行测试
func randomReportReq(i int32) *pb.ReportReq {
	var machineId int32 = i
	var cpuPercent float64 = 0.1
	var memPercent float64 = 0.2
	var diskPercent float64 = 0.3
	var timeStamp int64 = 10

	return &pb.ReportReq{MachineId: machineId, CpuPercent: cpuPercent, MemPercent: memPercent, DiskPercent: diskPercent, TimeStamp: timeStamp}
}

//Agent作为客户端（Clinet）开启连接
// func main() {
// 	flag.Parse()

// 	// conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
// 	conn, err := grpc.Dial("10.243.55.132:50051", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("fail to dial: %v", err)
// 	}
// 	defer conn.Close()
// 	client := pb.NewReportClient(conn)

// 	//控制goroutine数量
// 	pool := gopool.New(3)
// 	var i int32
// 	for i = 1; i <= 6; i++ {

// 		pool.Add(1)
// 		//go func
// 		go func(i int32) {
// 			//runAgentToJoin传输Agent信息
// 			runAgentToJoin(client, i)
// 			pool.Done()
// 		}(i)
// 	}
// 	pool.Wait()

//}
