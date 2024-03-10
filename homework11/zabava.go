package homework11

import (
	"bufio"
	"fmt"
	"strings"
)

func stringToBytes(str string) []byte {
	return []byte(str)
}

func bytesToString(b []byte) string {
	return string(b)
}

func encryptCaesar(text string) string {
	var result []byte
	for i := range text {
		if text[i] >= 'A' && text[i] <= 'Z' {
			result = append(result, 'A'+(text[i]-'A'+3)%26)
		} else if text[i] >= 'a' && text[i] <= 'z' {
			result = append(result, 'a'+(text[i]-'a'+3)%26)
		} else {
			result = append(result, text[i])
		}
	}
	return string(result)
}

func decryptCaesar(text string) string {
	var result []byte
	for i := range text {
		if text[i] >= 'A' && text[i] <= 'Z' {
			result = append(result, 'A'+(text[i]-'A'+26-3)%26)
		} else if text[i] >= 'a' && text[i] <= 'z' {
			result = append(result, 'a'+(text[i]-'a'+26-3)%26)
		} else {
			result = append(result, text[i])
		}
	}
	return string(result)
}

func replacePluses(text string) string {
	return strings.ReplaceAll(text, "+", " ")
}

func ExecuteStringOperations(reader *bufio.Reader) {

	fmt.Println("Выберите задачу:")
	fmt.Println("1 - Конвертация строки в слайс байтов и обратно")
	fmt.Println("2 - Шифр Цезаря")
	fmt.Println("3 - Замена '+' на пробелы")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("Введите строку:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		bytes := stringToBytes(input)
		fmt.Println("Слайс байтов:", bytes)
		str := bytesToString(bytes)
		fmt.Println("Строка:", str)
	case "2":
		fmt.Println("Выберите действие: 1 - шифровать, 2 - расшифровать")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)
		fmt.Println("Введите текст:")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if action == "1" {
			encrypted := encryptCaesar(text)
			fmt.Println("Зашифрованный текст:", encrypted)
		} else if action == "2" {
			decrypted := decryptCaesar(text)
			fmt.Println("Расшифрованный текст:", decrypted)
		}
	case "3":
		fmt.Println("Введите текст с '+' вместо пробелов:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		corrected := replacePluses(input)
		fmt.Println("Исправленный текст:", corrected)
	default:
		fmt.Println("Неверный выбор.")
	}
}
