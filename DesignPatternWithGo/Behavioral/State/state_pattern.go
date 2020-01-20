package main

import (
	"fmt"
	"log"
)

type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state
	currentState  state
	itemCount     int
	itemPrice     int
}

func newVendingMachine(itemCount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}

	hasItemState := &hasItemState{vendingMachine: v}
	itemRequestedState := &itemRequestedState{vendingMachine: v}
	hasMoneyState := hasMoneyState{vendingMachine: v}
	noItemState := noItemState{vendingMachine: v}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = &noItemState
	return v

}

func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}

type state interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

type hasItemState struct {
	vendingMachine *vendingMachine
}

func (h hasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	h.vendingMachine.incrementItemCount(count)
	return nil
}

func (h hasItemState) requestItem() error {
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
		return fmt.Errorf("no item present")
	}
	fmt.Printf("Item requestd\n")
	h.vendingMachine.setState(h.vendingMachine.itemRequested)
	return nil
}

func (h hasItemState) insertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (h hasItemState) dispenseItem() error {
	return fmt.Errorf("please select item first")
}

type noItemState struct {
	vendingMachine *vendingMachine
}

func (n *noItemState) addItem(item int) error {
	n.vendingMachine.incrementItemCount(item)
	n.vendingMachine.setState(n.vendingMachine.hasItem)
	return nil
}

func (n *noItemState) requestItem() error {
	return fmt.Errorf("item out of stock")
}

func (n *noItemState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (n *noItemState) dispenseItem() error {
	return fmt.Errorf("item out of stock")
}

type itemRequestedState struct {
	vendingMachine *vendingMachine
}

func (i *itemRequestedState) requestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *itemRequestedState) addItem(count int) error {
	return fmt.Errorf("item Dispense in progress")
}

func (i *itemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *itemRequestedState) dispenseItem() error {
	return fmt.Errorf("please insert money first")
}

type hasMoneyState struct {
	vendingMachine *vendingMachine
}

func (h hasMoneyState) addItem(int) error {
	return fmt.Errorf("item dispense in progress")
}

func (h hasMoneyState) requestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (h hasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("item dispense in progress")
}

func (h hasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	h.vendingMachine.itemCount = h.vendingMachine.itemCount - 1
	if h.vendingMachine.itemCount == 0 {
		h.vendingMachine.setState(h.vendingMachine.noItem)
	} else {
		h.vendingMachine.setState(h.vendingMachine.hasItem)
	}
	return nil
}


func main(){
	vendingMachine := newVendingMachine(1,10)
	err := vendingMachine.requestItem()
	if err != nil{
		log.Fatalln(err)
	}
	err = vendingMachine.insertMoney(5)
	if err != nil{
		log.Fatal(err)
	}

}
