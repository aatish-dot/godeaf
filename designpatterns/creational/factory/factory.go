package factory

import (
	"fmt"
)

//PaymentMethod uber type which will define the methods to implement
type PaymentMethod interface {
	Pay(amount float32) string
}

//Our current implemented Payment methods are defined as constants
const (
	Cash      = 1
	DebitCard = 2
)

//GetPaymentMethod returns a pointer to a PaymentMethod object or an error
//if the method is not registered. We used "new" operator to return the pointer
//but we could also used &Type{} although new makes it more readable for
//newcomers as it could be confusing
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitCardPM), nil
	default:
		return nil, fmt.Errorf("Payment method %v not defined", m)
	}
}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}
