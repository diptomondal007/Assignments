package bridge_practice

import "fmt"

type Epson struct {

}

func (e *Epson) PrintFile(){
	fmt.Println("Printing by a epson printer")
}