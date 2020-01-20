package main

import (
	"designpatternwithgo/Creational/Builder/builder_example"
	"fmt"
	)

func main(){
	normalBuilder := builder_example.GetBuilder("normal")
	iglooBuilder := builder_example.GetBuilder("igloo")

	director := builder_example.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()
	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.DoorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)
}

