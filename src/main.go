package main

import (
	"fmt"
	"strconv"
)

func main() {
	// === repaso ===
	// variables:
	//-declaración y asignación explicita
	var x int = 8
	//-declaración y asignacion implicita
	y := 7
	//impreción rapida
	print(x, " ")
	println(y)

	// go maneja errores de manera expicita
	//-ejm
	value, err := strconv.ParseInt("NaN", 0, 64)
	// nil: valor nulo, si retorna nil no hay error
	if err != nil {
		fmt.Printf("%v'\n\n", err)
	} else {
		fmt.Println(value)
	}

	// map
	m := make(map[string]int)
	m["key"] = 1
	fmt.Println(m["key"])

	// slice
	s := []int{1, 2, 3}
	//-recorrido de slice
	for i, value := range s {
		fmt.Print(i)
		fmt.Println(value)
	}
	//-agregar un valor al fina del slice
	s = append(s, 17)
	for i, value := range s {
		fmt.Print(i)
		fmt.Println(value)
	}
}

// go run para e desarrollo
// go build para producción
