package main

import (
	"fmt"
	"sync"
	"time"
)

// c := make(chan int, 1)
// c:=[]
// for -> c:=[goroutine]
// for
// c:=[] /*con doSomething(...){<-c}*/
// c:=[goroutine]
// endFor

// c := make(chan int, 2)
// c:=[][]
// for -> c:=[goroutine1][]
// for -> c:=[goroutine1][goroutine2]
// for
// c:=[][goroutine2] /*con doSomething(...){<-c}*/
// c:=[goroutine3][goroutine2]
// endFor

func main() {
	c := make(chan int, 1)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
	wg.Wait()
}
func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("id %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("id %d finished\n", i)
	<-c
}
