package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

/*
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
*/
func main() {
	pings := make(chan string)
	//	pongs := make(chan string)
	go ping(pings, "passed message")
	//	pong(pings, pongs)

	//	pings <- "aaaa"

	fmt.Println(<-pings)
}
