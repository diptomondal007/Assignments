package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List

	intList.PushBack(10)
	intList.PushBack(110)
	intList.PushBack(20)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}

}
