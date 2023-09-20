package main

// Observer interface defines the method for receiving payment updates.
type Observer interface {
	Update(paymentID string, status string)
}

// Subject interface defines methods for managing observers.
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(paymentID string, status string)
}
