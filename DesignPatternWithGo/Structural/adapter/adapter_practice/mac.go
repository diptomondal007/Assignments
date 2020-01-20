package adapter_practice

import "fmt"

type Mac struct {

}

func (m *Mac) InsertInSquarePort(){
	fmt.Println("Insert Square port into mac machine!")
}
