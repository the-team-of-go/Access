package main

import (
	ser "Access/service"
)

//"10.243.55.132:50051"
func main() {
	//判定服务的端口设定
	ser.RunJudge("127.0.0.1:50052")
	//ser.RunJudge("0.0.0.0:8081")

}
