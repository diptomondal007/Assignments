package main

import "fmt"

type Queue []*Order

type Order struct {
	priority     int
	quantity     int
	product      string
	customerName string
}

func (order *Order) New(priority int, quantity int, product string, customerName string) {
	order.priority = priority
	order.quantity = quantity
	order.product = product
	order.customerName = customerName
}

func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		var appended bool = false
		var i int
		var addedOrder *Order
		for i, addedOrder = range *queue {
			if order.priority > addedOrder.priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
			}
		}

		if !appended {
			*queue = append(*queue, order)
		}
	}
}

func main() {
	var queue Queue
	queue = make(Queue, 0)
	var order1 *Order = &Order{}
	var priority1 int = 2
	var quantity1 int = 20
	var product1 string = "Computer"
	var customerName1 string = "Greg White"
	order1.New(priority1, quantity1, product1, customerName1)
	var order2 *Order = &Order{}
	var priority2 int = 3
	var quantity2 int = 10

	var product2 string = "Monitor"
	var customerName2 string = "John Smith"
	order2.New(priority2, quantity2, product2, customerName2)
	queue.Add(order1)
	queue.Add(order2)
	var i int
	for i = 0; i < len(queue); i++ {
		fmt.Println(queue[i])
	}
}
