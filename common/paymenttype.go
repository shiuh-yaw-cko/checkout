package common

import "fmt"

// PaymentType ...
type PaymentType string

const (
	// Card ...
	Card PaymentType = "card"
	// ApplePay ...
	ApplePay PaymentType = "ApplePay"
	// GooglePay ...
	GooglePay PaymentType = "GooglePay"
)

func (c PaymentType) String() string {
	fmt.Println("Executing String() for PaymentType!")
	return string(c)
}