package main

import "fmt"

// funciones variadicas, permiten declarar funciones con numero de parametros desconocido
func sum(values ...int) int {
	var total int
	for _, num := range values {
		total += num
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// retornos con nombre
func getValues(x int) (double int, triple int, quad int) {
	double = x * 2
	triple = x * 3
	quad = x * 4
	// con los nombres de returns utilizados la función infiere que loos estas retornando
	return
}

func main() {
	fmt.Println(sum(1, 2, 3))
	printNames("pepito", "luchito")
	fmt.Println(getValues(2))
}
