package main

import "fmt"

func main() {
	creditCardPayment := &CreditCardPayment{}
	bankTransferPayment := &BankTransferPayment{}
	digitalWalletPayment := &DigitalWalletPayment{}

	loanWithCreditCard := NewLoanProduct(creditCardPayment)
	loanWithBankTransfer := NewLoanProduct(bankTransferPayment)
	loanWithDigitalWallet := NewLoanProduct(digitalWalletPayment)

	// Process payments for loans using different payment methods
	result1, _ := loanWithCreditCard.Pay(1000.0)
	result2, _ := loanWithBankTransfer.Pay(1500.0)
	result3, _ := loanWithDigitalWallet.Pay(800.0)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
