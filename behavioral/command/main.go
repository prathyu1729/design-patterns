package main

func main() {
	// Create the payment processor (receiver)
	paymentProcessor := NewPaymentProcessor()

	// Create payment commands
	command1 := NewPaymentCommand(paymentProcessor, "Account1", 100.0)
	command2 := NewPaymentCommand(paymentProcessor, "Account2", 50.0)
	command3 := NewPaymentCommand(paymentProcessor, "Account1", 200.0)

	// Create the broker (invoker)
	broker := Broker{}

	// Queue payment commands
	broker.AddCommand(command1)
	broker.AddCommand(command2)
	broker.AddCommand(command3)

	// Process queued payments
	broker.ProcessPayments()
}
