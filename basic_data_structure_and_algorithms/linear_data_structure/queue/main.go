package main

import "fmt"

type Queue []*Order

type Order struct {
	Priority     int
	Quantity     int
	Product      string
	CustomerName string
}

func NewOrder(pty int, qty int, pr string, cn string) *Order {
	return &Order{
		Priority:     pty,
		Quantity:     qty,
		Product:      pr,
		CustomerName: cn,
	}
}

func (queue *Queue) Add(order *Order) {
	for i, addedOrder := range *queue {
		if order.Priority > addedOrder.Priority {

			*queue = append((*queue)[:i], append([]*Order{order}, (*queue)[i:]...)...)

			return
		}
	}

	*queue = append(*queue, order)
}

func main() {
	queue := &Queue{}

	orderOne := NewOrder(2, 20, "Computer", "Greg White")
	orderTwo := NewOrder(1, 10, "Monitor", "Jhon Smith")

	queue.Add(orderOne)
	queue.Add(orderTwo)

	for _, v := range *queue {
		fmt.Println(v)
	}
}
