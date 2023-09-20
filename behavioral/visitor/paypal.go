package main

// PayPal is a concrete payment method representing a PayPal payment.
type PayPal struct {
	Amount float64
}

func (pp *PayPal) Accept(visitor PaymentCalculator) {
	visitor.VisitPayPal(pp)
}
