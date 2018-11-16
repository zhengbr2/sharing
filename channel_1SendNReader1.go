package main

import (
	"math/rand"
	"time"
)

func read(id int, ci chan int) {
	for i := 0; i < 10; i++ {
		println(id, ":", <-ci)
	}
}

func main() {

	ci := make(chan int)
	rand.Seed(time.Now().UnixNano())

	go func() {
		for {
			ci <- rand.Intn(100)
		}
	}()

	for id := 0; id < 10; id++ {
		go read(id, ci)
	}
	println("quiting main()")
}
