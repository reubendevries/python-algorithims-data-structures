package main

import "fmt"

func main() {
	simpleRecursionExample(10)
}

func simpleRecursionExample(n int) {
	if n == 1 {
		fmt.Println(n)
	} else {
		fmt.Println(n)
		simpleRecursionExample(n - 1)
	}
}
