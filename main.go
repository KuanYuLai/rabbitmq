package main

import (
	"fmt"

	"github.com/golang-queue/rabbitmq/module/exchange"
)

func main() {
	e := exchange.NewExchange("key1")
	e.NewQueue("q1")
	e.NewMessage("", "d", "fan out message")

	fmt.Println("\n\nNew queue")
	e.NewQueue("q2")
	e.NewMessage("direct", "q2", "direct q2 message")
	e.NewMessage("direct", "q1", "direct q1 message")
}
