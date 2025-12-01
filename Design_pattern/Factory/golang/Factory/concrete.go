package payment

import "fmt"

type CreditCardPayment struct{}

func (c CreditCardPayment) Pay(amount float64) {
	fmt.Println("Pay wih Credit Card:", amount)
}

type PaypalPayment struct{}

func (p PaypalPayment) Pay(amount float64) {
	fmt.Println("Pay wih Paypal:", amount)
}

type PromptPayPayment struct{}

func (p PromptPayPayment) Pay(amount float64) {
	fmt.Println("Pay wih PromptPayPayment:", amount)
}

// เพิ่ม payment ใหม่ → แก้แค่ที่เดียว

type CryptoPayment struct{}

func (c CryptoPayment) Pay(amount float64) {
	fmt.Println("Pay with Crypto:", amount)
}
