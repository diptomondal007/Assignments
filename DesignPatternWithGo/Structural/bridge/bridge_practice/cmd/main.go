package main

import (
	"designpatternwithgo/Structural/bridge/bridge_practice"
	"fmt"
)

func main(){
	hpPrinter := &bridge_practice.Hp{}
	epsonPrinter := &bridge_practice.Epson{}

	macComputer := bridge_practice.Mac{}
	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()
	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
}
