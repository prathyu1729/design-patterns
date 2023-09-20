package singleton

import (
	"sync"
	"time"
)

// TransactionLogger represents the singleton transaction logger.
type TransactionLogger struct {
	log []string
	mu  sync.Mutex
}

var instance *TransactionLogger
var once sync.Once

// GetTransactionLogger returns the singleton instance of the TransactionLogger.
func GetTransactionLogger() *TransactionLogger {
	once.Do(func() {
		instance = &TransactionLogger{}
	})
	return instance
}

// LogTransaction logs a transaction.
func (t *TransactionLogger) LogTransaction(message string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.log = append(t.log, time.Now().Format(time.RFC3339)+" - "+message)
}

// GetLog retrieves the transaction log.
func (t *TransactionLogger) GetLog() []string {
	return t.log
}
