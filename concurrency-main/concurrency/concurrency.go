package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func Channels_buffer_Main() {
	// unbuffered channel
	//c := make(chan int)
	//c <- 1
	//fmt.Println(<-c)

	// bufferd channel
	c := make(chan int, 1)

	c <- 1

	fmt.Println(<-c)

}

func WaitGroup_Main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}

	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("started in %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("ended")
}

func Sem() {
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
	c := make(chan int, 1)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething2(i, &wg, c)
	}
	wg.Wait()
}

func doSomething2(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("id %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("id %d finished\n", i)
	<-c
}

// generator(channel escritura)
func generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

// double(lectura, escritura)
func double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- 2 * value
	}
	close(out)
}

// print(lectura)
func print(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func Pipelines_Main() {
	v_generator := make(chan int)
	doubles := make(chan int)
	go generator(v_generator)
	go double(v_generator, doubles)
	print(doubles)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("worker with id %d started fib with %d\n", id, job)
		fib := fibonacci(job)
		fmt.Printf("worker with id %d, job %d and fib %d\n", id, job, fib)
		result <- fib
	}
}

func Worker_pools_MAin() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go worker(i, jobs, results)
	}

	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}
