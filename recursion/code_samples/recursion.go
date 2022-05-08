package main

import "fmt"

func main() {
	simpleRecursionExample(10)
	powerOfTwoRecursive(10)
	powerOfTwoIterative(10)
}

func simpleRecursionExample(n int) {
	if n == 1 {
		fmt.Println(n)
	} else {
		fmt.Println(n)
		simpleRecursionExample(n - 1)
	}
}

func powerOfTwoRecursive(n int) int {
	if n == 0 {
		return 1
	} else {
		return (powerOfTwoRecursive(n-1) * 2)
	}
}

func powerOfTwoIterative(n int) int {
	i := 0
	power := 1
	if i < n {
		power = power * 2
		i++
	}
	return power
}
