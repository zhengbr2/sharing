package main

import (
	"math/rand"
	"time"
)


func main(){

	ci:=make(chan int)
	rand.Seed(time.Now().UnixNano())

	go func(){
		for {
			ci <- rand.Intn(100)
		}
	}()

	for i:=0;i<10;i++ {
		println( i, ":", <-ci)
	}
}
