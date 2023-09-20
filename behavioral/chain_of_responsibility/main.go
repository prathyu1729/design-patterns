package main

import "fmt"

// PaymentRequest represents a payment request with an amount.
type PaymentRequest struct {
	Amount float64
}

// PaymentHandler is the interface for payment request handlers.
type PaymentHandler interface {
	SetNext(handler PaymentHandler)
	Handle(request *PaymentRequest)
}

// BasePaymentHandler is the base struct for payment handlers.
type BasePaymentHandler struct {
	nextHandler PaymentHandler
	threshold   float64
}

func (b *BasePaymentHandler) SetNext(handler PaymentHandler) {
	b.nextHandler = handler
}

// ConcretePaymentHandler handles payments up to a certain threshold.
type ConcretePaymentHandler struct {
	BasePaymentHandler
}

func NewConcretePaymentHandler(threshold float64) *ConcretePaymentHandler {
	return &ConcretePaymentHandler{
		BasePaymentHandler: BasePaymentHandler{threshold: threshold},
	}
}

func (c *ConcretePaymentHandler) Handle(request *PaymentRequest) {
	if request.Amount <= c.threshold {
		fmt.Printf("Payment handled by ConcretePaymentHandler (Threshold: $%.2f)\n", c.threshold)
	} else if c.nextHandler != nil {
		fmt.Printf("Payment passed from ConcretePaymentHandler to next handler (Threshold: $%.2f)\n", c.threshold)
		c.nextHandler.Handle(request)
	} else {
		fmt.Println("Payment cannot be handled.")
	}
}

func main() {
	// Create payment handlers with different thresholds.
	handler1 := NewConcretePaymentHandler(100.0)
	handler2 := NewConcretePaymentHandler(500.0)
	handler3 := NewConcretePaymentHandler(1000.0)

	// Construct the chain of responsibility.
	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	// Create payment requests.
	request1 := &PaymentRequest{Amount: 50.0}
	request2 := &PaymentRequest{Amount: 300.0}
	request3 := &PaymentRequest{Amount: 1500.0}

	// Process payment requests through the chain.
	handler1.Handle(request1)
	handler1.Handle(request2)
	handler1.Handle(request3)
}
