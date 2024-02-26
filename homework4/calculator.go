package homework4

import "fmt"

func PrintMultiplicationTable() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}
