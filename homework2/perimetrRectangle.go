package homework2

import "fmt"

// PerimeterRectangle находит периметр прямоугольника. Пользователь вводит длины сторон.
func PerimeterRectangle() {
	var length, width float64
	fmt.Println("Введите длину и ширину прямоугольника:")
	fmt.Scanln(&length, &width)
	perimeter := (length + width) * 2
	fmt.Printf("Периметр прямоугольника: %.2f\n", perimeter)
}
