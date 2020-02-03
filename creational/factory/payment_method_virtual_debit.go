package factory

import "fmt"

type VirtualDebitPM struct{}

func (c *VirtualDebitPM) Pay(amount float32) string {
	return fmt.Sprintf("%.02f paid using debit (virtual)\n", amount)
}
