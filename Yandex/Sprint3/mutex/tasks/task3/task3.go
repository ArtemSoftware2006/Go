package task3

import "sync"

type Queue interface {
	Enqueue(element interface{}) // положить элемент в очередь
	Dequeue() interface{}        // забрать первый элемент из очереди
}
type ConcurrentQueue struct {
	queue []interface{} // здесь хранить элементы очереди
	mutex sync.Mutex
}

func (q *ConcurrentQueue) Enqueue(element interface{}) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.queue = append(q.queue, element)
}

func (q *ConcurrentQueue) Dequeue() interface{} {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	elem := q.queue[0]
	q.queue = q.queue[1:]
	return elem
}
