package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	//establish a connection nc with the nats server (gnatsd) on default port :4222 with name "KWnats"
	nc, _ := nats.Connect("localhost", nats.Name("KWnats"), nats.Timeout(10*time.Second))
	defer nc.Close()			                      // connection will be closed after end of main

	// content
	subj, payload := "KWsubj", []byte("KWmsg")

	// send request, receive reply, set timeout of 2 sec
	reply, _ := nc.Request(subj, payload, 2*time.Second)

	// log request and reply
	log.Printf("Published [%s] : '%s'", subj, payload)
	log.Printf("Received  [%v] : '%s'", reply.Subject, string(reply.Data))
}
