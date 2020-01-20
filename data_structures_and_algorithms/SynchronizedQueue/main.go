package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	messagePassStart = iota
	messagePassEnd
	messageTicketStart
	messageTicketEnd
)

type Queue struct {
	waitPass    int
	waitTicket  int
	playPass    bool
	playTicket  bool
	queuePass   chan int
	queueTicket chan int
	message     chan int
}

func (queue *Queue) New() {
	queue.message = make(chan int)
	queue.queuePass = make(chan int)
	queue.queueTicket = make(chan int)

	go func() {
		var message int
		for {
			select {
			case message = <-queue.message:
				switch message {
				case messagePassStart:
					queue.waitPass++
				case messagePassEnd:
					queue.playPass = false
				case messageTicketStart:
					queue.waitTicket++
				case messageTicketEnd:
					queue.playTicket = false
				}
				if queue.waitPass > 0 && queue.waitTicket > 0 && !queue.playPass && !queue.playTicket {
					queue.waitPass--
					queue.waitTicket--
					queue.playPass = true
					queue.playTicket = true
					queue.queuePass <- 1
					queue.queueTicket <- 1
				}
			}
		}
	}()
}

func ticketIssue(queue *Queue) {
	for {
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		queue.TicketIssueStart()
		fmt.Println("Ticket issue start ")
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		queue.TicketIssueEnd()
	}
}

func (queue *Queue) TicketIssueStart() {
	queue.message <- messageTicketStart
	<-queue.queueTicket
}

func (queue *Queue) TicketIssueEnd() {
	queue.message <- messageTicketEnd
}

func (queue *Queue) StartPass() {
	queue.message <- messagePassStart
	<-queue.queuePass
}

func (queue *Queue) EndPass() {
	queue.message <- messagePassEnd
}

func passenger(Queue *Queue) {
	for {
		//sleep for a random time
		time.Sleep(time.Duration(rand.Intn(10000)) * time.Millisecond)
		Queue.StartPass()
		fmt.Println("Passenger starts")
		//sleep for 2 seconds
		time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		fmt.Println("Passenger Ends")
		Queue.EndPass()
	}
}

func main() {
	queue := &Queue{}
	queue.New()
	fmt.Println(queue)
	for i := 0; i < 10; i++ {
		go passenger(queue)
	}

	for j := 0; j < 5; j++ {
		go ticketIssue(queue)
	}

	select {}
}
