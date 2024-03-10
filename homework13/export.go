package homework13

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ExecuteQRInterfaceOperations(reader *bufio.Reader, wallet *Wallet) {
	fmt.Println("Выберите операцию:\n1 - Проверить баланс\n2 - Оплатить\n3 - Перевести средства")

	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	switch operation {
	case "1":
		balance, err := wallet.CheckBalance()
		if err != nil {
			fmt.Println("Ошибка при проверке баланса:", err)
			return
		}
		fmt.Println("Текущий баланс:", balance)
	case "2":
		fmt.Println("Введите сумму для оплаты:")
		amountStr, _ := reader.ReadString('\n')
		amount, err := strconv.ParseFloat(strings.TrimSpace(amountStr), 64)
		if err != nil {
			fmt.Println("Ошибка при конвертации введенной суммы:", err)
			return
		}
		err = wallet.Pay(amount)
		if err != nil {
			fmt.Println("Ошибка при оплате:", err)
			return
		}
		fmt.Println("Оплата прошла успешно")
	case "3":
		// Для простоты примера, создадим второй кошелек внутри этого случая
		targetWallet := &Wallet{Balance: 100}
		fmt.Println("Введите сумму для перевода:")
		amountStr, _ := reader.ReadString('\n')
		amount, err := strconv.ParseFloat(strings.TrimSpace(amountStr), 64)
		if err != nil {
			fmt.Println("Ошибка при конвертации введенной суммы:", err)
			return
		}
		err = wallet.Transfer(targetWallet, amount)
		if err != nil {
			fmt.Println("Ошибка при переводе средств:", err)
			return
		}
		fmt.Println("Перевод средств выполнен успешно")
	default:
		fmt.Println("Неверная операция")
	}
}
