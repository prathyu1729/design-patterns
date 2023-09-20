package main

// LoanProduct is a concrete financial product that uses a payment method.
type LoanProduct struct {
	paymentMethod PaymentMethod
}

func NewLoanProduct(paymentMethod PaymentMethod) *LoanProduct {
	return &LoanProduct{
		paymentMethod: paymentMethod,
	}
}

func (p *LoanProduct) Apply() string {
	return "Loan application submitted"
}

func (p *LoanProduct) Pay(amount float64) (string, error) {
	return p.paymentMethod.ProcessPayment(amount)
}
