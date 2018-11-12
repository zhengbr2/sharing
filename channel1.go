package main

func main() {

	var ci chan int
	// ci is a channel for type int
	ci = make(chan int)

	//or in this way
	// ci:=make(chan int)

	ci <- 100  // send to channel
	b := <-ci  // receive from channel
	println(b) // why it doesn' work?
}
