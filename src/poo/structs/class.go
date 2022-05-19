package main

import (
	"fmt"
)

type employee struct {
	id   int
	name string
}

func main() {
	e := employee{}
	fmt.Println(e)
	e.id = 1
	e.name = "Name"
	fmt.Println(e)
}
