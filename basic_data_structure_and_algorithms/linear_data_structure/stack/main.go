package main

import (
	"fmt"
	"strconv"
)

type Element struct {
	elementValue int
}

func (element *Element) String() string {
	return strconv.Itoa(element.elementValue)
}

type Stack struct {
	Elements     []*Element
	ElementCount int
}

func New() *Stack {
	return &Stack{
		Elements: make([]*Element, 0),
	}
}

func (stack *Stack) Push(el *Element) {
	stack.Elements = append(stack.Elements[:stack.ElementCount], el)
	stack.ElementCount += 1
}

func (stack *Stack) Pop() *Element {
	if len(stack.Elements) == 0 {
		return nil
	}

	element := stack.Elements[len(stack.Elements)-1]

	stack.Elements = stack.Elements[:len(stack.Elements)-1]

	return element
}

func main() {
	stack := New()

	stack.Push(&Element{3})
	stack.Push(&Element{2})
	stack.Push(&Element{1})

	stack.Pop()
	stack.Pop()

	for _, v := range stack.Elements {
		fmt.Println(*v)
	}
}
