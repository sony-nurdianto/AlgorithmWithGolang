package main

import "fmt"

func h(x int, y int) int {
	return x * y
}

func g(l int, m int) (x int, y int) {
	return 2 * l, 4 * m
}

func main() {
	fmt.Println(h(g(1, 2)))
}
