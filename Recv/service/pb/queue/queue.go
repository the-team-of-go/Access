package queue

import (
	pb "Access/service/pb/agent"
	"sync"
)

// //使用锁实现一个队列
// type SliceQueue struct {
// 	data []interface{}
// 	mu   sync.Mutex
// }

// func NewSliceQueue(n int) *SliceQueue {
// 	return &SliceQueue{data: make([]interface{}, 0, n)}
// }

// func (sq *SliceQueue) Enqueue(v interface{}) {
// 	sq.mu.Lock()
// 	sq.data = append(sq.data, v)
// 	sq.mu.Unlock()
// }

// //出队
// func (sq *SliceQueue) Dequeue() interface{} {
// 	sq.mu.Lock()
// 	defer sq.mu.Unlock()
// 	v := sq.data[0]
// 	sq.data = sq.data[1:]
// 	return v
// }

//使用锁实现一个队列
type SliceQueue struct {
	data  []*pb.ReportReq
	mu    sync.Mutex
	qsize int
}

func NewSliceQueue(n int) *SliceQueue {
	return &SliceQueue{data: make([]*pb.ReportReq, 0, n), qsize: 0}
}

func (sq *SliceQueue) Enqueue(v *pb.ReportReq) {
	sq.mu.Lock()
	sq.data = append(sq.data, v)
	sq.qsize = sq.qsize + 1
	sq.mu.Unlock()
}

//出队
func (sq *SliceQueue) Dequeue() *pb.ReportReq {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	v := sq.data[0]
	sq.data = sq.data[1:]
	sq.qsize = sq.qsize - 1
	return v
}

func (sq *SliceQueue) Getqsize() int {
	sq.mu.Lock()
	defer sq.mu.Unlock()
	return sq.qsize
}
