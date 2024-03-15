package homework14

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func WriteMessageToFile(filename, message string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Ошибка при создании файла: %s", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(message + "\n")
	if err != nil {
		log.Fatalf("Ошибка при записи сообщения: %s", err)
	}
	writer.Flush()
	fmt.Println("Сообщение записано в файл.")

}

func ReadMessageFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %s", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Ошибка при чтении сообщения: %s", err)
	}
	fmt.Printf("Прочитанное сообщение из файла: %s", message)
}
