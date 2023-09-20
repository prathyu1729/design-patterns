package main

func main() {
	// Create a payment gateway (subject).
	paymentGateway := NewPaymentGateway()

	// Create payment processors (observers).
	processor1 := NewPaymentProcessor("Processor 1")
	processor2 := NewPaymentProcessor("Processor 2")
	processor3 := NewPaymentProcessor("Processor 3")

	// Register payment processors as observers of the payment gateway.
	paymentGateway.RegisterObserver(processor1)
	paymentGateway.RegisterObserver(processor2)
	paymentGateway.RegisterObserver(processor3)

	// Simulate payment updates.
	paymentGateway.NotifyObservers("12345", "Payment successful")
	paymentGateway.NotifyObservers("67890", "Payment failed")

	// Unregister one observer.
	paymentGateway.RemoveObserver(processor2)

	// Simulate more updates.
	paymentGateway.NotifyObservers("54321", "Payment refunded")
}
