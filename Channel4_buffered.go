package main

import (
	"math/rand"
	"time"
)

func main() {

	ci:=make(chan int,1)

	// send to channel
	rand.Seed(time.Now().UnixNano())

	ci <- rand.Intn(100)  //send to channel

	b := <-ci // receive from channel
	println(b)
}
