package homework2

import "fmt"

// PerimeterSquare находит периметр квадрата. Пользователь вводит длину стороны.
func PerimeterSquare() {
	var side float64
	fmt.Println("Введите длину стороны квадрата:")
	fmt.Scanln(&side)
	perimeter := side * 4
	fmt.Printf("Периметр квадрата: %.2f\n", perimeter)
}
