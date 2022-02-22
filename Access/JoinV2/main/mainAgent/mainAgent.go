package main

import (
	ser "Access/service"
)

//本"10.243.55.132:50051"
//agent"10.243.156.108:50051"
func main() {
	//Agent的端口设定
	ser.RunAgentToJoin("127.0.0.1:50051")
	//ser.RunAgentToJoin("0.0.0.0:8080")
}
