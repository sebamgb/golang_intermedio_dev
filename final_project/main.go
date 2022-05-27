package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// ---job---
type Job struct {
	name   string
	delay  time.Duration
	number int
}

/* job's runner */
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// endJob

// ---worker---
type Worker struct {
	id         int
	jobQueue   chan Job
	workerPool chan chan Job
	quitChan   chan bool
}

func (w Worker) Start() {
	go func() {
		for {
			w.workerPool <- w.jobQueue

			select {
			case job := <-w.jobQueue:
				fmt.Printf("worker with id %d started\n", w.id)
				fib := fibonacci(job.number)
				time.Sleep(job.delay)
				fmt.Printf("worker with id %d finished with result %d\n", w.id, fib)
			case <-w.quitChan:
				fmt.Printf("worker with id %d stopped\n", w.id)
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quitChan <- true
	}()
}

/* constructor */
func NewWorker(id int, workerpool chan chan Job) *Worker {
	return &Worker{
		id:         id,
		jobQueue:   make(chan Job),
		workerPool: workerpool,
		quitChan:   make(chan bool),
	}
}

// endWorker

// ---dispatcher---
type Dispatcher struct {
	workerPool chan chan Job
	maxWorkers int
	jobQueue   chan Job
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			go func() {
				workerJobQueue := <-d.workerPool
				workerJobQueue <- job
			}()
		}
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		worker := NewWorker(i, d.workerPool)
		worker.Start()
	}

	go d.Dispatch()
}

/* constructor */
func NewDispatcher(jobqueue chan Job, maxworkers int) *Dispatcher {
	worker := make(chan chan Job, maxworkers)
	return &Dispatcher{
		jobQueue:   jobqueue,
		maxWorkers: maxworkers,
		workerPool: worker,
	}
}

// endDispatcher

// manejador de peticiones al server
func RequestHandler(w http.ResponseWriter, r *http.Request, jobqueue chan Job) {
	if r.Method != "POST" { // (X) GET, PUT, DELETE
		w.Header().Set("Alow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "invalid delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "invalid value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "invalid name", http.StatusBadRequest)
		return
	}

	job := Job{
		name:   name,
		delay:  delay,
		number: value,
	}

	jobqueue <- job
	w.WriteHeader(http.StatusCreated)
}

// endRequestHandler

func main() {
	const (
		max_workers    = 4
		max_queue_size = 20
		port           = ":8081"
	)

	job_queue := make(chan Job, max_queue_size)
	dispatcher := NewDispatcher(job_queue, max_workers)

	dispatcher.Run()

	// http://localhost:8081/fib
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, job_queue)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
