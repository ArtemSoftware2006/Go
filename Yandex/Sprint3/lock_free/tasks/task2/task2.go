package main

import (
	"sync/atomic"
)

func AtomicSwap(ptr1, ptr2 *int32) {
	val1 := atomic.LoadInt32(ptr1)
	val2 := atomic.LoadInt32(ptr2)

	atomic.StoreInt32(ptr1, val2)
	atomic.StoreInt32(ptr2, val1)
}
