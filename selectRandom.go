package main

import (
	"fmt"
	"time"
)

func main() {

	var c1, c2 = make(chan int), make(chan int)

	go func() { c1 <- 100 }()
	go func() { fmt.Println("case2,received from c2:", <-c2) }()

	time.Sleep(time.Millisecond * 1)   // await ready

	select {
	case i1 := <-c1:
		fmt.Println("case1 received: ", i1, " from c1")

	case c2 <- 88:
	default:    						// never executed
		fmt.Println("in default")
	}

	time.Sleep(time.Millisecond * 1)    // await printing logs
}
