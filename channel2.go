package main

import (
	"math/rand"
	"time"
)

func main() {

	ci := make(chan int)
	rand.Seed(time.Now().UnixNano())

	go func() {
		ci <- rand.Intn(100) //send to channel
		ci <- rand.Intn(100)
	}()

	b := <-ci // receive from channel
	println(b)
	println(<-ci)
}
