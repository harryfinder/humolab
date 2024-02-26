package homework7

import "fmt"

func ReverseSliceStrings(slice []string) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Println("Отзеркаленный массив строк:", slice)
}
