package main

import "fmt"

func main() {
	// unbuffered channel
	//c := make(chan int)
	//c <- 1
	//fmt.Println(<-c)

	// bufferd channel
	c := make(chan int, 1)

	c <- 1

	fmt.Println(<-c)
}
