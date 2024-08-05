// Отсортировать срез целых чисел от 1 до 5 в порядке убывания.
// package main

// import "fmt"

// func bubbleSort(slice []int) {
// 	n := len(slice)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if slice[j] < slice[j+1] {
// 				slice[j], slice[j+1] = slice[j+1], slice[j]
// 			}
// 		}
// 	}
// }
// func main() {
// 	slice := []int{1, 2, 3, 4, 5}
// 	bubbleSort(slice)
// 	fmt.Println(slice)
// }

//------------------------------------------------------------------------------------

// Отсортировать срез строк, содержащих имена животных, в порядке возрастания.
// package main

// import "fmt"

// func bubbleSortStrings(slice []string) {
// 	n := len(slice)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-i-1; j++ {
// 			if slice[j] > slice[j+1] {
// 				slice[j], slice[j+1] = slice[j+1], slice[j]
// 			}
// 		}
// 	}
// }
// func main() {
// 	slice := []string{"horse", "fox", "lion"}
// 	bubbleSortStrings(slice)
// 	fmt.Println(slice)
// }

//------------------------------------------------------------------------------------

// Отсортировать срез вещественных чисел от 0.1 до 0.5 в порядке убывания.
// package main

// import "fmt"

// func main() {
// 	slice := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
// 	n := len(slice)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-1-i; j++ {
// 			if slice[j] <= slice[j+1] {
// 				slice[j], slice[j+1] = slice[j+1], slice[j]
// 			}
// 		}
// 	}
// 	fmt.Println(slice)
// }

//------------------------------------------------------------------------------------

// Реализовать пузырьковую сортировку для среза структур Person, отсортировать по возрасту в порядке убывания.

// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	slice := []Person{{"Abussabur", 25}, {"Shodmon", 23}, {"Valy", 35}}
// 	n := len(slice)
// 	for i := 0; i < n-1; i++ {
// 		for j := 0; j < n-1-i; j++ {
// 			if slice[j].Age <= slice[j+1].Age {
// 				slice[j], slice[j+1] = slice[j+1], slice[j]
// 			}
// 		}
// 	}
// 	fmt.Println(slice)
// }