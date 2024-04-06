package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println("Error connecting to NATS: ", err)
		return
	}
	defer nc.Close()

	jsCtx, err := nc.JetStream()
	if err != nil {
		fmt.Println("Error creating jetstream context: ", err)
		return
	}

	_, err = jsCtx.AddStream(&nats.StreamConfig{
		Name:     "my-stream",
		Subjects: []string{"my-subject"},
	})
	if err != nil {
		fmt.Println("Error creating stream:", err)
		return
	}

	_, err = jsCtx.Publish("my-subject", []byte("HELLO WORLD!... FROM NATS JET-STREAM"))
	if err != nil {
		fmt.Println("Error publishing message:", err)
		return
	}

	sub, err := jsCtx.Subscribe("my-subject", func(m *nats.Msg) {
		fmt.Println("Received message:", string(m.Data))
	})
	if err != nil {
		fmt.Println("Error subscribing:", err)
		return
	}

	time.Sleep(10 * time.Second)

	sub.Drain()
}
