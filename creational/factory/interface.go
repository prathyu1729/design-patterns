package factory

// Account represents the account interface.
type Account interface {
	Deposit(amount float64) string
	Withdraw(amount float64) string
	GetBalance() float64
}

// AccountFactory represents the account factory interface.
type AccountFactory interface {
	CreateAccount() Account
}
