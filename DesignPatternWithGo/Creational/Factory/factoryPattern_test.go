package Factory

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T){
	payment, err := GetPaymentMethod(Cash)
	if err != nil{
		t.Error("A payment method of cash must exist")
	}
	msg := payment.Pay(20.0)
	if !strings.Contains(msg, "paid using cash"){
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("Log:", msg)
}

func TestCreatePaymentMethodDebitCard(t *testing.T){
	payment, err := GetPaymentMethod(DebitCard)
	if err != nil{
		t.Error("A payment of method debit card doesn't exist")
	}
	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodNotExistent(t *testing.T){
	_, err := GetPaymentMethod(20)
	if err == nil{
		t.Error("A payment method with Id 20 must return a error")
	}
	t.Log("Log: ", err)
}