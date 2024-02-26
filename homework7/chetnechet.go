package homework7

import "fmt"

func SplitSliceIntoEvenOdd(numbers []int) {
	var even, odd []int
	for _, num := range numbers {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	fmt.Println("Чётные числа:", even)
	fmt.Println("Нечётные числа:", odd)
}
