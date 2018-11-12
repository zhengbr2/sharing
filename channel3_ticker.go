package main

import (
	"log"
	"time"
)

func main() {

	ticker := time.Tick(time.Second * 1)

	for i := 0; i < 4; i++ {
		log.Println(i, ":", <-ticker)
	}
}
