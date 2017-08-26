package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/go-nats"
)

func main() {
	go createPublisher()

	for {

	}
}

func createPublisher() {

	log.Println("pub started")

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	msg := make([]byte, 16)

	for i := 0; i < 100000; i++ {
		if errPub := nc.Publish("alenSub", msg); errPub != nil {
			panic(errPub)
		}

		if (i % 100) == 0 {
			fmt.Println("i", i)
		}
		time.Sleep(time.Millisecond * 1)
	}

	log.Println("pub finish")

	errFlush := nc.Flush()
	if errFlush != nil {
		panic(errFlush)
	}

	errLast := nc.LastError()
	if errLast != nil {
		panic(errLast)
	}

}
