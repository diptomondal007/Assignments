package facade_practice

import "fmt"

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{
		balance: 0,
	}
}

func (w *wallet) creditBalance(amount int){
	w.balance = w.balance + amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *wallet) debitBalance(amount int) error{
	if w.balance < amount {
		return fmt.Errorf("balance is not sufficient")
	}
	w.balance = w.balance-amount
	fmt.Println("balance is sufficient")
	return nil
}
