package main

import "fmt"

func findElement(arr [10]int, k int) bool {
	for i := 0; i < 10; i++ {
		if arr[i] == k {
			return true
		}
	}
	return false
}

func main() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	check := findElement(arr, 2)
	fmt.Println(check)

	check2 := findElement(arr, 11)
	fmt.Println(check2)
}
