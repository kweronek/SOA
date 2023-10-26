package main

import (
	"log"
	"sync"

	//	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	//	connect to server
	nc, _ := nats.Connect("localhost:4222")
	defer nc.Close()

	// Use a WaitGroup to wait for a message to arrive
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe
	if _, err := nc.Subscribe("KWsubj", func(m *nats.Msg) {
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()

	/*

		// Subscribe
		for true {
			if _, err == nc.Subscribe("KWtest", func(m *nats.Msg) {
				log.Printf("Received a message: %s\n", string(m.Data))
			} else {
				log.Fatal(err)
			})
		}
	*/
}
