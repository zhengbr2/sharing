package main

import (
	"math/rand"
	"time"
)

func init() { rand.Seed(time.Now().UnixNano()) }

func gen(ci chan int) {
	for {
		ci <- rand.Intn(100)
	}
}

func main() {
	ci := make(chan int)

	for i := 0; i < 10; i++ {
		go gen(ci)
	}

	for i := 0; i < 100; i++ {
		println(i, ":", <-ci)
	}
}
