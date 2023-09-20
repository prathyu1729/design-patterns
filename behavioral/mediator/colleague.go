package main

import "fmt"

// Colleague interface defines methods for communication with other colleagues.
type Colleague interface {
	GetName() string
	Send(message string)
	Receive(message string)
}

// ConcreteColleague represents a concrete colleague that communicates through the mediator.
type ConcreteColleague struct {
	name     string
	mediator Mediator
}

func NewConcreteColleague(name string, mediator Mediator) *ConcreteColleague {
	return &ConcreteColleague{
		name:     name,
		mediator: mediator,
	}
}

func (c *ConcreteColleague) GetName() string {
	return c.name
}

func (c *ConcreteColleague) Send(message string) {
	fmt.Printf("%s sends: %s\n", c.name, message)
	c.mediator.NotifyColleague(c, message)
}

func (c *ConcreteColleague) Receive(message string) {
	fmt.Printf("%s receives: %s\n", c.name, message)
}
