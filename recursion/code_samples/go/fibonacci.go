package main

import "fmt"

func fibonacci_function(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci_function(n-2) + fibonacci_function(n-1)
}

func main() {
	x := fibonacci_function(10)
	fmt.Println(x)
}
