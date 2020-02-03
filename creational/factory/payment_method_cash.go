package factory

import "fmt"

type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%.02f paid using cash\n", amount)
}
