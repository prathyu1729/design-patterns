package factory

// CheckingAccountFactory creates checking accounts.
type CheckingAccountFactory struct{}

func (caf *CheckingAccountFactory) CreateAccount() Account {
	return NewCheckingAccount()
}

// SavingsAccountFactory creates savings accounts.
type SavingsAccountFactory struct{}

func (saf *SavingsAccountFactory) CreateAccount() Account {
	return NewSavingsAccount()
}
