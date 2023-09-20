package main

// Command interface defines the Execute method.
type Command interface {
	Execute()
}

// PaymentCommand represents a concrete command for processing payments.
type PaymentCommand struct {
	paymentProcessor *PaymentProcessor
	accountID        string
	amount           float64
}

func NewPaymentCommand(paymentProcessor *PaymentProcessor, accountID string, amount float64) *PaymentCommand {
	return &PaymentCommand{
		paymentProcessor: paymentProcessor,
		accountID:        accountID,
		amount:           amount,
	}
}

func (c *PaymentCommand) Execute() {
	c.paymentProcessor.ExecutePayment(c.accountID, c.amount)
}
