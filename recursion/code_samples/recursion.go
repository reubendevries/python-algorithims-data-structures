package main

import "fmt"

func main() {
	recusionExample(10)
}

func recusionExample(n int) {
	if n == 1 {
		fmt.Println(n)
	} else {
		fmt.Println(n)
		recusionExample(n - 1)
	}
}
