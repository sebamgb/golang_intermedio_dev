package main

import (
	"time"
)

type person struct {
	name string
	age  int
	DNI  string
}
type employee struct {
	id       int
	position string
}
type fullTimeEmployee struct {
	person
	employee
}

var GetPersonByDNI = func(dni string) (person, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM PERSON WHERE ...
	return person{}, nil
}

var GetEmployeeByID = func(id int) (employee, error) {
	time.Sleep(5 * time.Second)
	return employee{}, nil
}

func GetFullTimeEmloyeeById(id int, dni string) (fullTimeEmployee, error) {
	var ftEmployee fullTimeEmployee

	e, err := GetEmployeeByID(id)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.employee = e

	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.person = p
	return ftEmployee, nil
}
