package bridge_practice

import "fmt"

type Windows struct {
	printer Printer
}

func (w *Windows) Print(){
	fmt.Println("Print request from windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer){
	w.printer = p
}
