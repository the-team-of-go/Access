package main

import (
	service "Access/service"
)

//本"10.243.55.132:50051"
//agent"10.243.156.108:50051"
//judege 10.243.50.4:8080
func main() {

	//设置并发数量
	service.SetPoolSize(3)

	//判定服务的端口设定
	service.SetUserInfoServiceAddr("127.0.0.1:50052")
	//service.SetUserInfoServiceAddr("10.243.50.4:8080")

	//Agent的端口设定
	//service.RunJointToJudge("10.243.55.132:50051")
	service.RunJointToJudge("127.0.0.1:50051")

}
