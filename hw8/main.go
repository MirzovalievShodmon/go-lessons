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

// package main

// import "fmt"

// func main() {
// 	// Определяем исходные срезы
// 	slice1 := []float64{0, 1, 2, 1, 4, 5, 6, 7, 8, 9}
// 	slice2 := []float64{10, 2, 1, 13, 14, 15, 16, 17, 18, 19}

// 	// Объединяем срезы
// 	combined := append(slice1, slice2...)

// 	// Создаем новый срез для уникальных элементов
// 	result := []float64{}

// 	// Проходим по всем элементам объединенного среза
// 	for _, value := range combined {
// 		// Проверяем, есть ли уже этот элемент в результирующем срезе
// 		found := false
// 		for _, existing := range result {
// 			if value == existing {
// 				found = true
// 				break
// 			}
// 		}
// 		// Если элемент не найден в результирующем срезе, добавляем его
// 		if !found {
// 			result = append(result, value)
// 		}
// 	}

// 	// Выводим результат
// 	fmt.Println(result)
// }
// // if found ==>> if found==true
// // if !found ==>> if found==false

//-----------------------------------------------------------------------------------

// Поменять местами максимальный и минимальный элементы массива.
// package main

// import "fmt"

// func main() {
// 	slice := []float64{1, 5, 3, 2, 4}
// 	minind := 0
// 	maxind := 0
// 	for i := 1; i < len(slice); i++ {
// 		if slice[i] <= slice[minind] {
// 			minind = i
// 		}
// 		if slice[i] >= slice[maxind] {
// 			maxind = i
// 		}
// 	}
// 	slice[minind], slice[maxind] = slice[maxind], slice[minind]
// 	fmt.Println(slice)
// }

//-----------------------------------------------------------------------------------

// Проверить, является ли массив палиндромом.
// package main

// import "fmt"

// func main() {
// 	slice := []int{10, 23, 21, 6, 78, 78, 6, 21, 23, 10}
// 	mid := len(slice) / 2
// 	left := slice[:mid]
// 	right := slice[mid+(mid+1)%2:]
// 	c := 0
// 	for i := 0; i < mid; i++ {
// 		j := mid - i - 1
// 		if left[i] == right[j] {
// 			c++
// 		} else {
// 			break
// 		}
// 	}
// 	if c == mid {
// 		fmt.Println("Массив является палиндромом.")
// 	} else {
// 		fmt.Println("Массив не является палиндромом.")

// 	}
// }

//-----------------------------------------------------------------------------------
