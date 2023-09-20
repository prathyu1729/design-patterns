package main

func main() {
	// Create a concrete mediator.
	mediator := NewConcreteMediator()

	// Create concrete colleagues and register them with the mediator.
	colleague1 := NewConcreteColleague("Colleague 1", mediator)
	colleague2 := NewConcreteColleague("Colleague 2", mediator)

	mediator.RegisterColleague(colleague1)
	mediator.RegisterColleague(colleague2)

	// Communicate between colleagues through the mediator.
	colleague1.Send("Hello from Colleague 1!")
	colleague2.Send("Hi from Colleague 2!")
}
