package homework13

import "fmt"

// Pay снимает средства с баланса кошелька
func (w *Wallet) Pay(amount float64) error {
	if w.Balance < amount {
		return fmt.Errorf("недостаточно средств")
	}
	w.Balance -= amount
	return nil
}

// CheckBalance возвращает текущий баланс, снимая комиссию за проверку
func (w *Wallet) CheckBalance() (float64, error) {
	const fee = 0.50 // комиссия за проверку баланса
	if w.Balance < fee {
		return 0, fmt.Errorf("недостаточно средств для проверки баланса")
	}
	w.Balance -= fee
	return w.Balance, nil
}

// Transfer переводит средства с одного кошелька на другой
func (w *Wallet) Transfer(target *Wallet, amount float64) error {
	if w.Balance < amount {
		return fmt.Errorf("недостаточно средств для перевода")
	}
	w.Balance -= amount
	target.Balance += amount
	return nil
}
