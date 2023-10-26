package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	//	connect to server
	nc, _ := nats.Connect("localhost:4222")
	defer nc.Close()

	// Subscribe a subject
	sub, _ := nc.SubscribeSync("KWsubj")

	// Wait for a message
	msg, _ := sub.NextMsg(30 * time.Second)

	// Use the message
	log.Printf("Got message: %s", msg.Data)
}
