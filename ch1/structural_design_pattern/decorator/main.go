package main

import "fmt"

type IProcess interface {
	process()
}

type ProcessClass struct{}

func (process *ProcessClass) process() {
	fmt.Println("processClass process")
}

type ProcessDecorator struct {
	processInstance *ProcessClass
}

func (decorator *ProcessDecorator) process() {
	if decorator.processInstance == nil {
		fmt.Println("Process decorator process")
	} else {
		fmt.Printf("ProcessDecorator process and ")
		decorator.processInstance.process()
	}
}

func main() {
	process := &ProcessClass{}
	decorator := &ProcessDecorator{}

	decorator.process()

	decorator.processInstance = process

	decorator.process()
}
