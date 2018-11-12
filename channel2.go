package main

import (
	"math/rand"
	"time"
)

func main(){

	var ci chan int;
	// ci is a channel for type int
	ci =make(chan int)

	//or in this way
	// ci:=make(chan int)

	// send to channel
	rand.Seed(time.Now().UnixNano())
	go func(a int){
		ci<-a
	}(rand.Intn(100))

	b:= <- ci  // receive from channel
	println(b)
}
