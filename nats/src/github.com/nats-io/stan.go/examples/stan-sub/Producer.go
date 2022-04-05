package main

import (
	//"flag"
	"fmt"
	//"log"
	//"os"
	//"os/signal"
	//"time"

	//nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	//"github.com/nats-io/stan.go/pb"
)
//const DefaultNatsURL="nats://127.0.0.1:4222"

func main(){
sc, _ := stan.Connect("test-cluster", "wb-10-client")

cleanupDone := make(chan string)
go func() {
// Simple Synchronous Publisher
sc.Publish("foo", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming
cleanupDone <-string("foo")
cleanupDone<-string("Hello World") 
}()
fmt.Println(<-cleanupDone)

// Simple Async Subscriber
sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
    fmt.Printf("Received a message: %s\n", string(m.Data))
})

// Unsubscribe
sub.Unsubscribe()

// Close connection
sc.Close()
}