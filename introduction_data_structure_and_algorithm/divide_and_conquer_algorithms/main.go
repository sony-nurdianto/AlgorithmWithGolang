package main

import "fmt"

func fibonacci(k int) int {
	if k <= 1 {
		return 1
	}

	return fibonacci(k-1) + fibonacci(k-2)
}

func main() {
	m := 5
	for m = 0; m < 8; m++ {
		fib := fibonacci(m)
		fmt.Println(fib)
	}
}
