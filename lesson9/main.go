// Отсортировать срез целых чисел от 1 до 5 в порядке убывания.
package main

import "fmt"

func bubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] < slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}
func main() {
	slice := []int{1, 2, 3, 4, 5}
	bubbleSort(slice)
	fmt.Println(slice) // [5 4 3 2 1]
}