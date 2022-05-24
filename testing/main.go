package main

func sum(x, y int) int {

	return x + y
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
