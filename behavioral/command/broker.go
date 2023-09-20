package main

// Broker represents the invoker that queues and executes payment commands.
type Broker struct {
	commands []Command
}

func (b *Broker) AddCommand(command Command) {
	b.commands = append(b.commands, command)
}

func (b *Broker) ProcessPayments() {
	for _, command := range b.commands {
		command.Execute()
	}
	b.commands = nil // Clear the command queue after processing.
}
