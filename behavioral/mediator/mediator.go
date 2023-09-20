package main

// Mediator interface defines methods for communication between colleagues.
type Mediator interface {
	RegisterColleague(colleague Colleague)
	NotifyColleague(colleague Colleague, message string)
}

// ConcreteMediator represents a concrete mediator that manages colleague interactions.
type ConcreteMediator struct {
	colleagues map[string]Colleague
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{
		colleagues: make(map[string]Colleague),
	}
}

func (m *ConcreteMediator) RegisterColleague(colleague Colleague) {
	m.colleagues[colleague.GetName()] = colleague
}

func (m *ConcreteMediator) NotifyColleague(colleague Colleague, message string) {
	for name, c := range m.colleagues {
		if name != colleague.GetName() {
			c.Receive(message)
		}
	}
}
