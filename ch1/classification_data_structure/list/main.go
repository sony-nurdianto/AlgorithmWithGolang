package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List

	for i := 1; i <= 5; i++ {
		intList.PushBack(i)
	}

	for e := intList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
