package main

import (
	"fmt"
	"math/rand"
)

func GenNum(done <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < 10; i++ {
			select {
			case <-done:
				return
			case out <- rand.Intn(10):
			}
		}
	}()

	return out
}

func GenArr(done <-chan struct{}, n <-chan int) <-chan []int {
	out := make(chan []int)

	go func() {
		defer close(out)

		var arr []int
		for v := range n {
			arr = append(arr, v)
		}

		fmt.Println(arr)

		select {
		case <-done:
			return
		case out <- arr:
		}
	}()

	return out
}

func orDone(done <-chan struct{}, c <-chan []int) <-chan []int {
	valueStream := make(chan []int)

	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case valueStream <- v:
				case <-done:
				}
			}
		}
	}()

	return valueStream
}

func TeeMinMax(done <-chan struct{}, in <-chan []int) (_, _ <-chan int) {
	out1 := make(chan int)
	out2 := make(chan int)

	go func() {
		defer close(out1)
		defer close(out2)

		for val := range orDone(done, in) {
			ch1, ch2 := out1, out2

			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case ch1 <- GetMin(val):
					ch1 = nil
				case ch2 <- GetMax(val):
					ch2 = nil
				}
			}
		}
	}()

	return out1, out2
}

func GetMin(arr []int) int {
	k := arr[0]

	for i := range arr {
		if k > arr[i] {
			k = arr[i]
		}
	}

	return k
}

func GetMax(arr []int) int {
	k := arr[0]

	for i := range arr {
		if k < arr[i] {
			k = arr[i]
		}
	}

	return k
}

func main() {
	done := make(chan struct{})
	defer close(done)

	minimum, maximum := TeeMinMax(done, GenArr(done, GenNum(done)))

	fmt.Printf("Minimum value: %d, Maximum value: %d\n", <-minimum, <-maximum)
}
