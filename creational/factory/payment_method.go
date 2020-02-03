package factory

import (
	"errors"
	"fmt"
)

const (
	Cash  string = "cash"
	Debit string = "debit"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetPaymentMethod(m string) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case Debit:
		// return new(DebitPM), nil
		return new(VirtualDebitPM), nil
	default:
		return nil, errors.New(
			fmt.Sprintf("Payment method %s not recognized\n", m),
		)
	}
}
