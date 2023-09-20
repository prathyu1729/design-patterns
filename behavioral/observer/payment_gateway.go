package main

// PaymentGateway represents the payment gateway (subject) that maintains a list of observers.
type PaymentGateway struct {
	observers map[Observer]bool
}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{
		observers: make(map[Observer]bool),
	}
}

func (pg *PaymentGateway) RegisterObserver(observer Observer) {
	pg.observers[observer] = true
}

func (pg *PaymentGateway) RemoveObserver(observer Observer) {
	delete(pg.observers, observer)
}

func (pg *PaymentGateway) NotifyObservers(paymentID string, status string) {
	for observer := range pg.observers {
		observer.Update(paymentID, status)
	}
}
