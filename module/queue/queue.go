package queue

import "fmt"

// Queue contains all information needed for Queue
type Queue struct {
	routingKey string
	messages   []interface{}
	size       int
}

// Dequeue pop a message from queue
func (q *Queue) Dequeue() interface{} {
	if len(q.messages) == 0 {
		panic("Queue is empty")
	}
	msg := q.messages[0]
	q.messages = q.messages[1:]
	return msg
}

// Enqueue append a message into queue
func (q *Queue) Enqueue(v interface{}) {
	fmt.Println(q.routingKey + " gets new message")
	q.messages = append(q.messages, v)
}

// Create create a new queue
func Create(key string, msg ...interface{}) *Queue {
	return &Queue{
		routingKey: key,
		messages:   msg,
	}
}

// RoutingKey returns the routing key of the queue
func (q *Queue) RoutingKey() string {
	return q.routingKey
}
