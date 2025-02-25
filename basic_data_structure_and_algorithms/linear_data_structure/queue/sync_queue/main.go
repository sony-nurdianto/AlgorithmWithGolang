package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MessagePassStart = iota
	MessageTicketStart
	MessagePassEnd
	MessageTicketEnd
)

type Queue struct {
	WaitPass    int
	WaitTicket  int
	PlayPass    int
	PlayTicket  int
	QueuePass   chan int
	QueueTicket chan int
	Message     chan int
}

func NewQueue() *Queue {
	return &Queue{
		QueuePass:   make(chan int),
		Message:     make(chan int),
		QueueTicket: make(chan int),
	}
}

func (queue *Queue) StartTicketIssue() {
	queue.Message <- MessageTicketStart
	<-queue.QueueTicket
}

func (queue *Queue) EndTicketIssue() {
	queue.Message <- MessageTicketEnd
}

func ticketIsue(queue *Queue) {
	for {
		time.Sleep(time.Duration(rand.Intn(10_000)) * time.Millisecond)
		queue.StartTicketIssue()
		fmt.Println("Ticket Issue starts")
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("Ticket Issue End")
		queue.EndTicketIssue()
	}
}

func (queue *Queue) StartPass() {
	queue.Message <- MessagePassStart
	<-queue.QueuePass
}

func (queue *Queue) EndPass() {
	queue.Message <- MessagePassEnd
}

func passenger(Queue *Queue) {
	for {
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		Queue.StartPass()
		fmt.Println(" Passenger starts")
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println(" Passenger ends")
		Queue.EndPass()
	}
}

func main() {
	Queue := NewQueue()
	fmt.Println(Queue)
	var i int
	for i = 0; i < 10; i++ {
		// fmt.Println(i, "passenger in the Queue")
		go passenger(Queue)
	}
	// close(Queue.queuePass)
	var j int
	for j = 0; j < 5; j++ {
		// fmt.Println(i, "ticket issued in the Queue")
		go ticketIsue(Queue)
	}

	go func() {
		for {
			msg := <-Queue.Message // Selalu menunggu pesan masuk
			switch msg {
			case MessagePassStart:
				Queue.WaitPass++
			case MessagePassEnd:
				Queue.PlayPass = 0
			case MessageTicketStart:
				Queue.WaitTicket++
			case MessageTicketEnd:
				Queue.PlayTicket = 0
			}

			// Jika ada penumpang dan tiket yang menunggu, izinkan mereka masuk
			if Queue.WaitPass > 0 && Queue.WaitTicket > 0 && Queue.PlayPass == 0 && Queue.PlayTicket == 0 {
				Queue.PlayPass = 1
				Queue.PlayTicket = 1
				Queue.WaitTicket--
				Queue.WaitPass--
				Queue.QueuePass <- 1
				Queue.QueueTicket <- 1
			}
		}
	}()

	select {}
}
