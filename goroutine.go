package main

import (
	"time"
)

func double(a int) int { return a * 2 }

func main() {


	for i := 0; i < 2000000; i++ {
		go println(double(i))
	}
	time.Sleep(time.Second * 15)
}
