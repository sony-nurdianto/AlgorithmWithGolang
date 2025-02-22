package main

import "fmt"

func PowerSeries(a int) (int, int) {
	return a * a, a * a * a
}

func main() {
	square, cube := PowerSeries(3)

	fmt.Printf("Square: %d\nCube: %d\n", square, cube)
}
