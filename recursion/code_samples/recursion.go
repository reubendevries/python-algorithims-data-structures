package main

import (
	"fmt"
	"log"
)

func main() {
	a := factorial_number(5)
	fmt.Println(a)
	simpleRecursionExample(10)
	b := powerOfTwoRecursive(10)
	fmt.Println(b)
	c := powerOfTwoIterative(10)
	fmt.Println(c)
	d := fibonacci_function(1)
	log.Printf("Result: %d\n", d)
}

func simpleRecursionExample(n int) {
	if n == 1 {
		fmt.Println(n)
	} else {
		simpleRecursionExample(n - 1)
	}
}

func powerOfTwoRecursive(n int) int {
	if n == 0 {
		return 1
	} else {
		power := powerOfTwoRecursive(n - 1)
		return power * 2
	}
}

func powerOfTwoIterative(n int) int {
	i := 0
	power := 1
	for i < n {
		power = power * 2
		i++
	}
	return power
}

func factorial_number(n int) int {
	if n < 0 {
		return -1
	} else if n < 2 {
		return 1
	} else {
		return (n * factorial_number(n-1))
	}
}

func fibonacci_function(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci_function(n-2) + fibonacci_function(n-1)
}
