package factory

import "fmt"

func main() {
	// Create a checking account using the factory
	checkingFactory := &CheckingAccountFactory{}
	checkingAccount := checkingFactory.CreateAccount()

	// Deposit and withdraw from the checking account
	fmt.Println(checkingAccount.Deposit(1000.0))
	fmt.Println(checkingAccount.Withdraw(500.0))
	fmt.Println("Checking Account Balance:", checkingAccount.GetBalance())

	// Create a savings account using the factory
	savingsFactory := &SavingsAccountFactory{}
	savingsAccount := savingsFactory.CreateAccount()

	// Deposit and withdraw from the savings account
	fmt.Println(savingsAccount.Deposit(2000.0))
	fmt.Println(savingsAccount.Withdraw(1000.0))
	fmt.Println("Savings Account Balance:", savingsAccount.GetBalance())
}
