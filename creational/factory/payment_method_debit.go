package factory

import "fmt"

type DebitPM struct{}

func (c *DebitPM) Pay(amount float32) string {
	return fmt.Sprintf("%.02f paid using debit\n", amount)
}
