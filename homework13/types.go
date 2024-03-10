package homework13

// QR определяет интерфейс для работы с кошельком
type QR interface {
	Pay(amount float64) error
	CheckBalance() (float64, error)
	Transfer(target *Wallet, amount float64) error
}

// Wallet реализует интерфейс QR
type Wallet struct {
	Balance float64
}
