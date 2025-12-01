package checkout

import "fmt"

type CredtiCardPayment struct{}

func (c CredtiCardPayment) Pay(amount float64) {
	fmt.Println("Paying with Credit Card:", amount)
}

type PaypalPayment struct{}

func (p PaypalPayment) Pay(amount float64) {
	fmt.Println("Paying with Credit Card:", amount)
}

type PromptPayPayment struct{}

func (p PromptPayPayment) Pay(amount float64) {
	fmt.Println("Paying with Credit Card:", amount)
}

func Checkout(amount float64, method string) {
	if method == "card" {
		p := CredtiCardPayment{}
		p.Pay(amount)
	} else if method == "Paypal" {
		p := PaypalPayment{}
		p.Pay(amount)
	} else if method == "Promptpay" {
		p := PromptPayPayment{}
		p.Pay(amount)
	}
}

// ❌ ปัญหา

// if-else ยาว

// ต้อง import concrete struct ทุกครั้ง

// ถ้าเพิ่ม CryptoPayment → ต้องแก้ทุกไฟล์ที่ใช้ payment

// โค้ดติดกับ concrete struct (coupling สูง)

// พอนาน ๆ ไปจะเละมาก
