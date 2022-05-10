package main

import "fmt"

func main() {
	simpleRecursionExample(10)
	x := powerOfTwoRecursive(10)
	fmt.Println(x)
	y := powerOfTwoIterative(10)
	fmt.Println(y)
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
