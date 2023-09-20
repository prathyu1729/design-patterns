package prototype

// Transaction defines the transaction interface.
type Transaction interface {
	Clone() Transaction
}
