package main

import(
	"fmt"
	"strconv"
	"stan"
)

func main(){
	sc, _:=stan.Connect("prod","simple=pub")
	defer sc.Close()

	for i:=1; ;i++{
		sc.Publish("bestellugen", []byte("besstelung"+strconv.Itoa(i))
	time.Sleep(2*time.Second)
	}
}