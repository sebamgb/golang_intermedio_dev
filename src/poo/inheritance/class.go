package main

import "fmt"

type person struct {
	name string
	age  int
}
type employee struct {
	id int
}
type fullTimeEmployee struct {
	person
	employee
}

func main() {
	ftEmployee := fullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 22
	ftEmployee.id = 1
	fmt.Printf(" Full time employee %v\n", ftEmployee) //imprime claramente la composicion claramente la composición a la herencia que aplica go
}
