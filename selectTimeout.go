package main

import (
	"fmt"
	"time"
)

func doTask(job string)  <- chan string{
	result:=make( chan string );
	go func () { time.Sleep(time.Second * 2); result <- job + " task is done!"}()
	return result
}

func main() {

	for {
		select {
		case s := <- doTask("Joe"):
			fmt.Println(s)
			return
		case <-time.After(time.Second * 1):
			fmt.Println("You're too slow.")
			return
		}
	}
}