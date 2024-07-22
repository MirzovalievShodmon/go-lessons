// Напишите функцию updateValue, которая принимает указатель на целое число и увеличивает его значение на 10.
// package main

// import (
// 	"fmt"
// )

//	func main() {
//		var x int = 5
//		fmt.Println(x)
//		updateValue(&x)
//		fmt.Println(x)
//	}
//
//	func updateValue(p *int) {
//		*p = *p + 10
//	}
//-------------------------------------------------------------------------------------------------
// Напишите функцию swap, которая меняет местами значения двух переменных с
// использованием указателей.
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a, b int = 5, 6
// 	fmt.Println(a, " ", b)
// 	swap(&a, &b)
// 	fmt.Println(a, " ", b)
// }

//	func swap(a, b *int) {
//		 c := *a
//		 *a = *b
//		 *b = c
//
// //	(или) *a, *b = *b, *a
// }
//-------------------------------------------------------------------------------------------------
// Напишите функцию printValue, которая принимает указатель на целое число и
// выводит его значение.
// Если указатель равен nil, выведите сообщение "Указатель пуст".
// package main

// import "fmt"

// func main() {
// 	var ptr *int
// 	printValue(ptr)
// 	a := 23
// 	printValue(ptr)
// }
// func printValue(ptr *int) {
// 	if ptr == nil {
// 		fmt.Println("Указатель пуст")
// 	} else {
// 		fmt.Println(*ptr)
// 	}
// }
//-------------------------------------------------------------------------------------------------
// Напишите функцию increment, которая принимает указатель на целое число и
// увеличивает его значение на 1.
// package main

// import "fmt"

//	func main() {
//		a := 10
//		increment(&a)
//		fmt.Println(a)
//	}
//
//	func increment(ptr *int) {
//		*ptr++
//	}
//-------------------------------------------------------------------------------------------------
// Напишите функцию isPositive, которая принимает указатель на целое число и
// возвращает true,
// если значение положительное, и false в противном случае.
// package main

// import "fmt"

// func main() {
// 	a := -5
// 	fmt.Println(isPositive(&a))
// }

//	func isPositive(ptr *int) (b bool) {
//		return *ptr > 0
//	}
//-------------------------------------------------------------------------------------------------
// Напишите функцию changeString, которая принимает указатель на строку и
// добавляет к ней слово "Go".
// package main

// import "fmt"

// func main() {
// 	a := "lets's "
// 	changeString(&a)
// 	fmt.Println(a)

// }
//
//	func changeString(ptr *string) {
//		*ptr = *ptr + "Go"
//	}
//-------------------------------------------------------------------------------------------------
// Напишите функцию double, которая принимает указатель на целое число и
// удваивает его значение.
// package main

// import "fmt"

//	func main() {
//		a := 10
//		double(&a)
//		fmt.Println(a)
//	}
//
//	func double(ptr *int) {
//		*ptr *= 2
//	}
//
// -------------------------------------------------------------------------------------------------
// Напишите функцию isEven, которая принимает указатель на целое число и
// возвращает true, если число четное,
// и false, если нечетное.
// package main

// import "fmt"

//	func main() {
//		a := 10
//		b:=11
//		fmt.Println(isEven(&a))
//		fmt.Println(isEven(&b))
//	}
//
//	func isEven(ptr *int) (b bool) {
//		return *ptr%2 == 0
//	}
//-------------------------------------------------------------------------------------------------
// Напишите функцию checkNil, которая принимает указатель на целое число и
// проверяет, является ли он нулевым (nil).
// package main

// import "fmt"

//	func main() {
//		var ptr *int
//		a := 10
//		fmt.Println(checkNil(ptr))
//		ptr = &a
//		fmt.Println(checkNil(ptr))
//	}
//
//	func checkNil(ptr *int) (b bool) {
//		return ptr == nil
//	}
//
// -------------------------------------------------------------------------------------------------
// Middle Level
// Реализуйте функции для хранения и вывода имени и возраста человека
// с использованием указателей.
// package main

// import "fmt"

// func main() {
// 	var name string
// 	var age int
// 	namePtr := &name
// 	agePtr := &age
// 	storeName(namePtr, "Шодмон")
// 	storeAge(agePtr, 23)
// 	printPerson(namePtr, agePtr)
// }
// func storeName(namePtr *string, name string) {
// 	*namePtr = name
// }

// func storeAge(agePtr *int, age int) {
// 	*agePtr = age
// }

//	func printPerson(namePtr *string, agePtr *int) {
//		fmt.Printf("Имя: %sец, Возраст: %d лет\n", *namePtr, *agePtr)
//	}
// -------------------------------------------------------------------------------------------------
// Напишите программу для учета посещаемости студентов на курсе.
// Реализуйте функции для отметки студента как посетившего курс
// и вывода списка студентов с их статусом посещения.
package main

import "fmt"

func main() {
	var name1, name2 string
	var attended1, attended2 bool
	namePtr1 := &name1
	namePtr2 := &name2
	attendedPtr1 := &attended1
	attendedPtr2 := &attended2
	markAttendance(namePtr1, attendedPtr1)
	printAttendance(namePtr1, attendedPtr1)
	printAttendance(namePtr2, attendedPtr2)
}
func markAttendance(namePtr *string, attendedPtr *bool) {
	*attendedPtr = true
}
func printAttendance(namePtr *string, attendedPtr *bool) {
	status := "не посетил"
	if *attendedPtr {
		status = "посетил"
	}
	fmt.Printf("Студент %s - %s\n", *namePtr, status)
}