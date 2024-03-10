package homework8_9

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func CreateUser(name string, age int, country string) User {
	return User{Name: name, Age: age, Country: country}
}

func AddUser(id string, user User) {
	UsersMap[id] = user
}

func GetUser(id string) User {
	return UsersMap[id]
}

func UpdateUser(id string, user User) {
	UsersMap[id] = user
}

func DeleteUser(id string) {
	delete(UsersMap, id)
}

func ExecuteCRUDOperations(reader *bufio.Reader) {
	fmt.Println("Выберите операцию:\n1 - Create\n2 - Read\n3 - Update\n4 - Delete")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	switch op {
	case "1":
		fmt.Println("Введите имя, возраст и страну через запятую (например, John,30,USA):")
		input, _ := reader.ReadString('\n')
		parts := strings.Split(strings.TrimSpace(input), ",")
		age, _ := strconv.Atoi(parts[1])
		user := CreateUser(parts[0], age, parts[2])
		userId := parts[0]
		AddUser(userId, user)
		fmt.Println("Пользователь создан:", userId)
	case "2":
		fmt.Println("Введите ID пользователя для чтения:")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)
		user := GetUser(id)
		fmt.Printf("Пользователь: %+v\n", user)
	case "3":
		fmt.Println("Введите ID пользователя для обновления и новые данные через запятую (например, John,JohnDoe,31,Canada):")
		input, _ := reader.ReadString('\n')
		parts := strings.Split(strings.TrimSpace(input), ",")
		id := parts[0]
		age, _ := strconv.Atoi(parts[2])
		user := CreateUser(parts[1], age, parts[3])
		UpdateUser(id, user)
		fmt.Println("Пользователь обновлен:", id)
	case "4":
		fmt.Println("Введите ID пользователя для удаления:")
		id, _ := reader.ReadString('\n')
		id = strings.TrimSpace(id)
		DeleteUser(id)
		fmt.Println("Пользователь удален:", id)
	default:
		fmt.Println("Неверная операция")
	}
}
