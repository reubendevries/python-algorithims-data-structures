package main

import "fmt"

func main() {
	x := factorial_number(5)
	fmt.Println(x)
	simpleRecursionExample(10)
	y := powerOfTwoRecursive(10)
	fmt.Println(y)
	z := powerOfTwoIterative(10)
	fmt.Println(z)
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
