package bridge_practice

import "fmt"

type Hp struct {

}

func (hp *Hp) PrintFile(){
	fmt.Println("Printing by a hp printer")
}
