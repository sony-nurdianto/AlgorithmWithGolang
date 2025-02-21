package main

import (
	"fmt"
	"sort"
	"sync"
)

func Generator(done <-chan struct{}, num [3][3]int) <-chan [3]int {
	stream := make(chan [3]int)

	go func() {
		defer close(stream)

		for _, n := range num {
			select {
			case <-done:
				return
			case stream <- n:
			}
		}
	}()

	return stream
}

func Indexing(done <-chan struct{}, num <-chan [3]int) <-chan int {
	stream := make(chan int)

	go func() {
		defer close(stream)

		for n := range num {
			for _, v := range n {
				select {
				case <-done:
					return
				case stream <- v:
				}
			}
		}
	}()

	return stream
}

func Grouping(done <-chan struct{}, numStream <-chan int) <-chan []int {
	var wg sync.WaitGroup
	out := make(chan []int)

	var arr []int
	for n := range numStream {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			select {
			case <-done:
				return
			default:
				arr = append(arr, i)
			}
		}(n)
	}

	go func() {
		wg.Wait()
		out <- arr
		close(out)
	}()

	return out
}

type SortIntArr []int

func (a SortIntArr) Len() int           { return len(a) }
func (a SortIntArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortIntArr) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	done := make(chan struct{})
	defer close(done)
	twoDArray := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	grouping := Grouping(done, Indexing(done, Generator(done, twoDArray)))

	var arr SortIntArr
	if data, ok := <-grouping; ok {
		arr = SortIntArr(data)
	}

	sort.Sort(arr)
	fmt.Println(arr)
}
