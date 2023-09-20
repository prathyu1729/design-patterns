package factory

// SavingsAccount represents a savings account.
type SavingsAccount struct {
	balance float64
}

func NewSavingsAccount() *SavingsAccount {
	return &SavingsAccount{}
}

func (s *SavingsAccount) Deposit(amount float64) string {
	s.balance += amount
	return "Deposited $" + formatAmount(amount)
}

func (s *SavingsAccount) Withdraw(amount float64) string {
	if s.balance >= amount {
		s.balance -= amount
		return "Withdrawn $" + formatAmount(amount)
	}
	return "Insufficient funds"
}

func (s *SavingsAccount) GetBalance() float64 {
	return s.balance
}
