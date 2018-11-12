package main

import (
	"math/rand"
	"time"
	"sync"
)

func read2(id int, ci chan int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		println(id, ":", <-ci)
	}
	wg.Done()   //one go routine done
}

func main() {

	ci := make(chan int)
	rand.Seed(time.Now().UnixNano())

	go func() {
		for {
			ci <- rand.Intn(100)
		}
	}()

	var wg sync.WaitGroup   //declare
	wg.Add(10)    // will initiate 10 go routine
	for id := 0; id < 10; id++ {
		go read2(id, ci, &wg)    // initiating
	}
	wg.Wait()    // wait here
}
