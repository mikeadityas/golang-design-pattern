package factory

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}
	t.Log("LOG: ", msg)
}

func TestGetPaymentMethodDebit(t *testing.T) {
	payment, err := GetPaymentMethod(Debit)
	if err != nil {
		t.Fatal("A payment method of type 'Debit' must exist")
	}

	msg := payment.Pay(22.30)
	if !strings.Contains(msg, "paid using debit") {
		t.Error("The debit payment method message wasn't correct")
	}
	t.Log("LOG: ", msg)
}

func TestGetPaymentMethodNonExistent(t *testing.T) {
	_, err := GetPaymentMethod("nonExistent")

	if err == nil {
		t.Error("A payment method with ID 'nonExistent' must return an error")
	}
	t.Log("LOG: ", err)
}
