package exchange

import (
	"github.com/golang-queue/rabbitmq/module/queue"
)

// Exchange contains all information needed for exchange
type Exchange struct {
	queues []*queue.Queue
}

// NewExchange start a new exchange with single queue attach to it
func NewExchange(routingKey string) *Exchange {
	return &Exchange{}
}

// NewQueue attach a new queue to exchange
func (e *Exchange) NewQueue(routingKey string) {
	e.queues = append(e.queues, queue.Create(routingKey))
}

// NewMessage add a new message to queue, default using fan out mode
func (e *Exchange) NewMessage(mode, key string, v interface{}) {
	switch mode {
	case "direct":
		for i := range e.queues {
			if e.queues[i].RoutingKey() == key {
				e.queues[i].Enqueue(v)
				break
			}
		}
	default:
		for i := range e.queues {
			e.queues[i].Enqueue(v)
		}
	}
}
