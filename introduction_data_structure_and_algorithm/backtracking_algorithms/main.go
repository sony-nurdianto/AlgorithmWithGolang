// main package untuk contoh Backtracking
package main

// Mengimpor fmt package
import (
	"fmt"
)

// Fungsi rekursif untuk mencari kombinasi elemen yang jumlahnya sama dengan k
func findElementsWithSum(arr [10]int, combinations [19]int, size int, k int, addValue int, l int, m int) int {
	num := 0
	if addValue > k {
		return -1
	}
	if addValue == k {
		num = num + 1
		p := 0
		for p = 0; p < m; p++ {
			fmt.Printf("%d,", arr[combinations[p]])
		}
		fmt.Println(" ")
	}
	var i int
	for i = l; i < size; i++ {
		combinations[m] = l
		findElementsWithSum(arr, combinations, size, k, addValue+arr[i], l, m+1)
		l = l + 1
	}
	return num
}

// Fungsi utama (main)
func main() {
	arr := [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	addedSum := 18
	var combinations [19]int
	findElementsWithSum(arr, combinations, 10, addedSum, 0, 0, 0)
}
