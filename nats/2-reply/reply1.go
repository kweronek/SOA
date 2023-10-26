package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	nc, _ := nats.Connect("localhost:4222")
	defer nc.Close()

	// Subscribe
	sub, _ := nc.SubscribeSync("KWsubj")

	for true {
		// Read a message
		msg, _ := sub.NextMsg(10 * time.Second)
		log.Println("received: ", string(msg.Data))

		// Send the time as the response.
		msg.Respond([]byte("reply-message"))
	}
}
