package homework2

import "fmt"

func AreaRectangle() {
	var length, width float64
	fmt.Println("Введите длину и ширину прямоугольника:")
	fmt.Scanln(&length, &width)
	area := length * width
	fmt.Printf("Площадь прямоугольника: %.2f\n", area)
}
