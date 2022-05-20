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
	endDate string
}

func (ftEmployee fullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type temporalyEmployee struct {
	person
	employee
	extRate int
}

func (tEmployee temporalyEmployee) getMessage() string {
	return "Temporaly Employee"
}

type printInfo interface {
	getMessage() string
}

func getMessage(p printInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := fullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 22
	ftEmployee.id = 1
	fmt.Printf(" Full time employee %v\n", ftEmployee) //imprime claramente la composicion claramente la composición a la herencia que aplica go
	tEmployee := temporalyEmployee{}
	getMessage(ftEmployee)
	getMessage(tEmployee)
}
