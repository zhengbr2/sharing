package main

import (
	"fmt"
	"time"
	"math/rand"
)

func doTask2(job string)  <- chan string{

	result:=make( chan string )
	go func () { time.Sleep(time.Second * time.Duration( rand.Intn(5))); result <- job + " is done!"}()
	return result
}

func fanIn(input1 , input2 <-chan string) <-chan string {

	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:  c <- s
			case s := <-input2:  c <- s
			}
		}
	}()
	return c
}
func main() {
	rand.Seed(time.Now().Unix())
	c:=fanIn(doTask2("Joe"), doTask2("Tom"))
	fmt.Println(<-c)
}