package main

import "fmt"

func powerOfTwo(n int) int {
	if n == 0 {
		return 1
	} else {
		power := powerOfTwo(n - 1)
		return power * 2
	}
}

func main() {
	x := powerOfTwo(10)
	fmt.Println(x)
}
