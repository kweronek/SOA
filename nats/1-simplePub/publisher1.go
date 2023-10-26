package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	//establish a connection nc with the nats server (gnatsd) on default port :4222 with name "KWnats"
	nc, _ := nats.Connect("localhost", nats.Name("KWnats"), nats.Timeout(10*time.Second))
	defer nc.Close() // connection will be closed after end of main

	// content
	subj, payload := "KWsubj", []byte("KWmsg")

	// send request, receive reply
	nc.Publish(subj, payload)
	nc.Flush()
	// log request and reply
	log.Printf("Published [%s] : '%s'", subj, payload)
}
