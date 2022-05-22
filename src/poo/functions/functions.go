package main

import (
	"fmt"
	"time"
)

func main() {
	x := 5
	// doble con función anonima
	y := func() int {
		return x * 2
	}()
	fmt.Println(y)

	z := 10
	// don't repet yourself
	w := func() int {
		return z * 2
	}()
	fmt.Println(w)

	// goroutine con función anonima
	c := make(chan int, 1)
	go func() {
		fmt.Println("Starting function...")
		time.Sleep(3 * time.Second)
		fmt.Println("end")
		c <- 1
	}()
	<-c
}
