package task4

import "sync"

var mux sync.Mutex
var Buf []int

func Write(num int) {
	mux.Lock()
	defer mux.Unlock()
	Buf = append(Buf, num)
}

func Consume() int {
	mux.Lock()
	defer mux.Unlock()

	first := Buf[0]
	Buf = Buf[1:]
	return first
}
