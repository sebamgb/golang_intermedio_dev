package main

import (
	"fmt"
)

type employee struct {
	id       int
	name     string
	vacation bool
}

func newEmployee(id int, name string, vacation bool) *employee {
	return &employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// forma 1
	e := employee{}
	fmt.Println(e)
	// forma 2
	e2 := employee{
		id:       1,
		name:     "Name",
		vacation: true,
	}
	fmt.Println(e2)
	// forma 3 (equivalente a forma 1)
	//-instancia apunta al struct
	e3 := new(employee)
	fmt.Println(e3)  //imprime una referencia, mas explicita que la dirección de memoria
	fmt.Println(*e3) //imprime el valor
	//-explicita luego de haber instanciado
	e3.id = 1
	e3.name = "Name 3"
	e3.vacation = true
	fmt.Println(*e3)
	// forma 4
	e4 := newEmployee(4, "Name 4", true)
	fmt.Printf("%v\n", *e4)
}
