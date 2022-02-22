### 测试使用方法
1.运行mainJudge中的mainJudge.go 即判定服务打开服务端，接收接入服务发送的信息
```shell
o run Recv/JoinV2/main/mainJudge/mainJudge.go 
```
2.运行mainRecv中的mainRecv.go 即接入服务打开准备接收Agent信息并向判定服务发送信息
```shell
go run Recv/JoinV2/main/mainRecv/mainRecv.go 
```
3.运行mainAgent中的mainAgent.go 即Agent调用其客户端向接入服务发送信息
```shell
go run Recv/JoinV2/main/mainAgent/mainAgent.go 
```