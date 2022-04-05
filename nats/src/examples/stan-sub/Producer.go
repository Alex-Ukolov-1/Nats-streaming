package main

import(
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
)

const{
	clusterID="test-cluster"
	clientID="Hellowworld-pub"
	svrtURL="nats://192.168.99.100:4222"
}

func main(){
	fmt.Println("hellow world publisher")
	sc, _:=stan.Connect(clusterOD,clientID,stan.NatsURL(svrtURL))
	defer sc.Close()
	for n:=0;n<3;n++{
		msg, _:=time.Now().MarshalText()
		sc.Publish(subject,msg)
		fmt.Printf("%s: %s\n",subject,msg)
	}
	
}