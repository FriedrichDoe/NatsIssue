package main

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/nats-io/go-nats"
)

var received int64

func main() {
	received = 0

	go createSubscriber()
	go check()

	for {

	}
}

func createSubscriber() {

	log.Println("sub started")

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	_, errSub := nc.Subscribe("alenSub", func(msg *nats.Msg) {
		atomic.AddInt64(&received, 1)
	})
	if errSub != nil {
		panic(errSub)
	}

	errFlush := nc.Flush()
	if errFlush != nil {
		panic(errFlush)
	}

	for {
		errLast := nc.LastError()
		if errLast != nil {
			panic(errLast)
		}
	}
}

func check() {
	for {
		fmt.Println("-----------------------")
		fmt.Println("still running")
		fmt.Println("received", atomic.LoadInt64(&received))
		fmt.Println("-----------------------")
		time.Sleep(time.Second * 2)
	}
}
