package homework2

import (
	"fmt"
)

func YearsToJubilee() {
	var age int
	fmt.Println("Введите ваш возраст:")
	fmt.Scanln(&age)
	yearsToJubilee := 10 - age%10

	fmt.Printf("До следующего юбилея осталось: %d лет\n", yearsToJubilee)
}
