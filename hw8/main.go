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

//Найти второе наибольшее число в массиве.
// package main

// import "fmt"

// func main() {
// 	slice := []float64{1, 45, 3, 6, 8, 4, 5, 8, 25, 6}
// 	maxslice := slice[0]
// 	result := slice[0]
// 	for i := 1; i < len(slice); i++ {
// 		if slice[i] >= maxslice {
// 			result = maxslice
// 			maxslice = slice[i]
// 		} else if slice[i] <= maxslice && result <= slice[i] {
// 			result = slice[i]
// 		}
// 	}
// 	fmt.Println(result)
// }

//-----------------------------------------------------------------------------------

//Перевернуть массив.
// package main

// import "fmt"

// func main() {
// 	slice := []float64{26, 56, 58, 62, 26, 12, 36, 32, 12}
// 	n := len(slice)
// 	m := n / 2
// 	for i := 0; i < m; i++ {
// 		slice[i], slice[2*m-(n+1)%2-i] = slice[2*m-(n+1)%2-i], slice[i]
// 	}
// 	fmt.Println(slice)
// }

//-----------------------------------------------------------------------------------

//Удалить дубликаты из массива.

// package main

// import "fmt"

// func main() {
// 	slice := []float64{1, 2, 3, 4, 5, 6, 3, 4, 6}
// 	Theslice := []float64{}

// 	for i := 0; i < len(slice); i++ {
// 		duplicateFound := false
// 		for j := 0; j < len(Theslice); j++ {
// 			if slice[i] == Theslice[j] {
// 				duplicateFound = true
// 				break
// 			}
// 		}
// 		if !duplicateFound {
// 			Theslice = append(Theslice, slice[i])
// 		}
// 	}

// 	fmt.Println(Theslice)
// }

//-----------------------------------------------------------------------------------

// Переместить все нули в конце массива, сохраняя порядок
// ненулевых элементов.

// package main

// import "fmt"

// func main() {
// 	slice := []float64{32, 68, 0, 19, 26, 0, 37}
// 	s := []float64{}
// 	c := []float64{}
// 	for _, v := range slice {
// 		if v != 0 {
// 			s = append(s, v)
// 		} else {
// 			c = append(c, v)
// 		}
// 	}
// 	slice = append(s, c...)
// 	fmt.Println(slice)
// }

//-----------------------------------------------------------------------------------

//Найти пересечение двух массивов.

// package main

// import "fmt"

// func main() {
// 	slice1 := []float64{1, 2, 3, 4, 5, 6, 3, 4, 6, 0}
// 	slice2 := []float64{4, 0, 3, 34, 5, 6, 0, 4, 66, 45, 24, 9}
// 	slice := []float64{}
// 	Theslice1 := []float64{}
// 	Theslice2 := []float64{}
// 	for i := 0; i < len(slice1); i++ {
// 		duplicateFound := false
// 		for j := 0; j < len(Theslice1); j++ {
// 			if slice1[i] == Theslice1[j] {
// 				duplicateFound = true
// 				break
// 			}
// 		}
// 		if !duplicateFound {
// 			Theslice1 = append(Theslice1, slice1[i])
// 		}
// 	}
// 	for i := 0; i < len(slice2); i++ {
// 		duplicateFound := false
// 		for j := 0; j < len(Theslice2); j++ {
// 			if slice2[i] == Theslice2[j] {
// 				duplicateFound = true
// 				break
// 			}
// 		}
// 		if !duplicateFound {
// 			Theslice2 = append(Theslice2, slice2[i])
// 		}
// 	}
// 	for _, v := range Theslice1 {
// 		for i := 0; i < len(Theslice2); i++ {
// 			if v == Theslice2[i] {
// 				slice = append(slice, v)
// 				break
// 			}
// 		}
// 	}
// 	fmt.Println(slice)
// }

//-----------------------------------------------------------------------------------

// Проверить, является ли массив подмножеством другого массива.

package main

import "fmt"

func isSubset(slice1, slice2 []float64) bool {
	for _, v1 := range slice1 {
		found := false
		for i, v2 := range slice2 {
			if v1 == v2 {
				slice2[i] = -1 // Используем -1 или любой другой уникальный маркер, чтобы исключить этот элемент из последующих проверок
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
func main() {
	slice1 := []float64{60, 7, 3, 4, 2, 2, 0, 45}
	slice2 := []float64{0, 2, 4, 0, 2, 4, 0, 60, 7, 0, 3, 45, 0, 45}

	if isSubset(slice1, slice2) {
		fmt.Println("Да, является подмножеством")
	} else {
		fmt.Println("Нет, не является подмножеством")
	}
}
