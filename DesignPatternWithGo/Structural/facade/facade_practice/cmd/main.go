package main

import (
	"designpatternwithgo/Structural/facade/facade_practice"
	"fmt"
	"log"
)

func main(){
	wallet := facade_practice.NewWalletFacade("abc",123)
	err := wallet.AddMoneyToWallet("abc",123,20)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println()
	err = wallet.DeductMoneyFromWallet("abc", 123, 25)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
