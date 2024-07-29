// Найти максимальный элемент в массиве.
// package main

// import "fmt"

// func main() {
// 	var arr [10]float64
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Scan(&arr[i], " ")
// 	}
// 	max := arr[0]
// 	for i := 1; i < len(arr); i++ {
// 		if arr[i] > max {
// 			max = arr[i]
// 		}
// 	}
// 	fmt.Println(max)
// }

//-----------------------------------------------------------------------------------

// Найти минимальный элемент в массиве.
// package main

// import "fmt"

// func main() {
// 	var arr [10]float64
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Scan(&arr[i], " ")
// 	}
// 	min := arr[0]
// 	for i := 1; i < len(arr); i++ {
// 		if arr[i] < min {
// 			min = arr[i]
// 		}
// 	}
// 	fmt.Println(min)
// }

//-----------------------------------------------------------------------------------

// Подсчитать количество положительных чисел в массиве.
// package main

// import "fmt"

// func main() {
// 	var arr [10]float64
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Scan(&arr[i], " ")
// 	}
// 	k := 0
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] > 0 {
// 			k++
// 		}
// 	}
// 	fmt.Println(k)
// }

//-----------------------------------------------------------------------------------

// Найти сумму всех элементов массива.
// package main

// import "fmt"

// func main() {
// 	var arr [10]float64
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Scan(&arr[i], " ")
// 	}
// 	s := float64(0)
// 	for i := 0; i < len(arr); i++ {
// 		s = s + arr[i]
// 	}
// 	fmt.Println(s)
// }

//-----------------------------------------------------------------------------------

// Найти среднее значение всех элементов массива.
// package main

// import "fmt"

// func main() {
// 	var arr [10]float64
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Scan(&arr[i], " ")
// 	}
// 	s := float64(0)
// 	for i := 0; i < len(arr); i++ {
// 		s = s + arr[i]
// 	}
// 	s = s / float64(len(arr))
// 	fmt.Println(s)
// }

//-----------------------------------------------------------------------------------

// Удалить все вхождения заданного числа из массива.
// package main

// import "fmt"

// func main() {
// 	var arr [10]float64
// 	var n float64
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Scan(&arr[i], " ")
// 	}
// 	fmt.Scan(&n)
// 	EmptyArr := []float64{}
// 	for i := 0; i < len(arr); i++ {
// 		if arr[i] != n {
// 			EmptyArr = append(EmptyArr, arr[i])
// 		}
// 	}
// 	fmt.Println(EmptyArr)
// }

//-----------------------------------------------------------------------------------

// Умножить все элементы массива на заданное число.
// package main

// import "fmt"

// func main() {
// 	var arr [10]float64
// 	var n float64
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Scan(&arr[i], " ")
// 	}
// 	fmt.Scan(&n)
// 	for i := 0; i < len(arr); i++ {
// 		n = n * arr[i]
// 	}
// 	fmt.Println(n)
// }

//-----------------------------------------------------------------------------------

// Найти все индексы заданного числа в массиве.
// package main

// import "fmt"

// func main() {
// 	var arr [10]float64
// 	var n float64
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Scan(&arr[i], " ")
// 	}
// 	fmt.Scan(&n)
// 	for i := 0; i < len(arr); i++ {
// 		if n == arr[i] {
// 			fmt.Println(i, " ")
// 		}
// 	}
// }

//-----------------------------------------------------------------------------------

// Создать копию массива.
// package main

// import "fmt"

// func main() {
// 	var arr [10]float64
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Scan(&arr[i], " ")
// 	}
// 	EmptyArr := []float64{}
// 	for i := 0; i < len(arr); i++ {
// 		EmptyArr = append(EmptyArr, arr[i])
// 	}
// 	fmt.Println(EmptyArr)
// }

//-----------------------------------------------------------------------------------

// Объединить два массива.
package main

import "fmt"

func main() {
	slice1 := make([]float64, 10)
	slice2 := make([]float64, 10)
	for i := 0; i < len(slice1); i++ {
		fmt.Scan(&slice1[i], " ")
	}
	for i := 0; i < len(slice2); i++ {
		fmt.Scan(&slice1[2], " ")
	}
	slice := slice1
	for i := 0; i < len(slice1); i++ {
		for j := 1; j < len(slice2); j++ {
			if slice1[i] == slice2[j] {
				slice = append(slice, slice2[j])
			}
		}
	}
	fmt.Println(slice)
}
