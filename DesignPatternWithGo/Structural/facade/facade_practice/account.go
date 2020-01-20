package facade_practice

import "fmt"

type account struct {
	name string
}

func newAccount(accountName string) *account{
	return &account{name:accountName}
}

func (a *account) checkAccount(accountName string) error{
	if a.name !=accountName{
		return fmt.Errorf("acount name is not correct")
	}
	fmt.Println("Account verified")
	return nil
}
