package main

import (
	"fmt"
)

type employee struct {
	id   int
	name string
}

// en los structs los metodos se asignan utilizando receiver functions

func (e *employee) setId(id int) {
	e.id = id
}

func (e *employee) setName(name string) {
	e.name = name
}

func (e employee) getId() int {
	return e.id
}

func (e employee) getName() string {
	return e.name
}

func main() {
	e := employee{}
	fmt.Println(e)
	e.id = 1
	e.name = "Name"
	fmt.Println(e)
	e.setId(5)
	e.setName("Name 2")
	fmt.Println(e)
	fmt.Println()
	fmt.Println(e.getId())
	fmt.Println(e.getName())
}
