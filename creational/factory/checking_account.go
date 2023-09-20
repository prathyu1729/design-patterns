package factory

import "fmt"

// CheckingAccount represents a checking account.
type CheckingAccount struct {
	balance float64
}

func NewCheckingAccount() *CheckingAccount {
	return &CheckingAccount{}
}

func (c *CheckingAccount) Deposit(amount float64) string {
	c.balance += amount
	return "Deposited $" + formatAmount(amount)
}

func (c *CheckingAccount) Withdraw(amount float64) string {
	if c.balance >= amount {
		c.balance -= amount
		return "Withdrawn $" + formatAmount(amount)
	}
	return "Insufficient funds"
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}

// Helper function to format the amount.
func formatAmount(amount float64) string {
	return fmt.Sprintf("%.2f", amount)
}
