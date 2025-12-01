package payment

import "fmt"

func NewPayment(method string) (Payment, error) {

	switch method {

	case "card":
		return CreditCardPayment{}, nil
	case "paypak":
		return PaypalPayment{}, nil

	case "promptpay":
		return PromptPayPayment{}, nil

		// if you have newfeature
	case "crypto":
		return CryptoPayment{}, nil

	default:
		return nil, fmt.Errorf("unknown payment method: %s", method)
	}
}
