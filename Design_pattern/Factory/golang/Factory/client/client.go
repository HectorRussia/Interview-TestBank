package main

import (
	"fmt"
	"log"
)

// Payment interface
type Payment interface {
	Pay(amount float64)
}

// Concrete payment types
type CreditCardPayment struct{}

func (c CreditCardPayment) Pay(amount float64) {
	fmt.Println("Pay with Credit Card:", amount)
}

type PaypalPayment struct{}

func (p PaypalPayment) Pay(amount float64) {
	fmt.Println("Pay with Paypal:", amount)
}

type PromptPayPayment struct{}

func (p PromptPayPayment) Pay(amount float64) {
	fmt.Println("Pay with PromptPayPayment:", amount)
}

type CryptoPayment struct{}

func (c CryptoPayment) Pay(amount float64) {
	fmt.Println("Pay with Crypto:", amount)
}

// Factory function
func NewPayment(method string) (Payment, error) {
	switch method {
	case "card":
		return CreditCardPayment{}, nil
	case "paypal":
		return PaypalPayment{}, nil
	case "promptpay":
		return PromptPayPayment{}, nil
	case "crypto":
		return CryptoPayment{}, nil
	default:
		return nil, fmt.Errorf("unknown payment method: %s", method)
	}
}

func Checkout(amount float64, method string) {
	p, err := NewPayment(method)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	p.Pay(amount) // client ไม่ต้องรู้ว่าคือ struct อะไร
}

func main() {
	Checkout(100.0, "card")
}
