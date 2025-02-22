package main

import "fmt"

type IProcess interface {
	process()
}

type Adaptee struct {
	adapterType int
}

func (adaptee Adaptee) Convert() {
	fmt.Println("Adaptee convert method")
}

type Adapter struct {
	adaptee Adaptee
}

func (adapter Adapter) process() {
	fmt.Println("Adapter Process")
	adapter.adaptee.Convert()
}

func main() {
	processor := Adapter{}

	processor.process()
}
