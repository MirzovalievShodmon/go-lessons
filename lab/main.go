//--------------------BE CAREFUL !!!---------------------------------------------------------------
//
// package main

// import "fmt"

// func main() {
// 	fmt.Print("hello\n")
// 	fmt.Print("hello")
//fmt.Print("FizzBuzz", " ")
// 	fmt.Printf("Наибольший общий делитель чисел %d и %d равен: %d\n", 1, 2, 3)
//fmt.Printf("Имя: %s, Возраст: %d лет\n", "Shodmon", 23)
//-------------------------------------------------------------------------------------------------
// Бесконечный цыкл или процедура
//
// package main

// import "fmt"

//	func main() {
//		newFunc()
//	}
//
//	func newFunc()  {
//		fmt.Print("hello\n")
//		newFunc()
//	}
//
// ------------------------------------------------------------------------------------------------
// НОД двух чисел
//
// package main

// import "fmt"

// var a, b, c int

// func main() {
// 	fmt.Scan(&a, &b)

//		for b != 0 {
//			c = a
//			a = b
//			b = c % b
//		}
//		fmt.Println(c)
//	}
//
// -------------------------------------------------------------------------------------------------
//БЕСКОНЕЧНЫЙ FOR
//
// package main

// import "fmt"

// func main() {
// 	for {
// 		fmt.Print("Hello", "      ")
// 	}
// }
//-------------------------------------------------------------------------------------------------
// НАХОЖДЕНИЕ ПРОСТЫХ ЧИСЕЛ
//
// package main

// import (
// 	"fmt"
// )

//	func main() {
//			var number int
//			fmt.Scan(&number)
//			fmt.Println(isPrime(number))
//		}
//
//	func isPrime(n int) bool {
//		if n == 1 {
//			return false
//		}
//		for i := 2; i*i <= n; i++ {
//			if n%i == 0 {
//				return false
//			}
//		}
//		return true
//	}
//-------------------------------------------------------------------------------------------------
////НЕ ЗАБУДЬ: дефолтное значение ДЕФОЛТНОЕ ЗНАЧЕНИЕ ДЛЯ "BOOLEAN"- ЭТО "FALSE"
// ------------------------------------------------------------------------------------------------
//ТАБЛИЦА УМНОЖЕНИЯ
//
// package main

// import "fmt"

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	fmt.Println("ТАБЛИЦА УМНОЖЕНИЯ\n")
// 	for i := 0; i <= 10; i++ {
// 		for j := 0; j <= n; j++ {
// 			fmt.Print(j, "x", i, "=", i*j)
// 			if i != 10 {
// 				if i*j > 9 {
// 					fmt.Print("     ")
// 				}
// 				if i*j <= 9 {
// 					fmt.Print("      ")
// 				}
// 			} else {
// 				if i*j > 9 {
// 					fmt.Print("    ")
// 				}
// 				if i*j <= 9 {
// 					fmt.Print("     ")
// 				}
// 			}
// 		}
// 		fmt.Println()
// 	}
// }
//-------------------------------------------------------------------------------------------------
// СЛОЖЕНИЕ В  ==>> sring
//
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
//
// ------------------------------------------------------------------------------------------------
//РАБОТА С УКАЗАТЕЛЕМ
// package main

// import "fmt"

// func main() {
// 	num := 10
// 	fmt.Println(num)

//		ptrNum := &num
//		sum(ptrNum)
//		fmt.Println(num)
//	}
//
//	func sum(ptrNum *int) {
//		*ptrNum = *ptrNum - 1
//	}
//
// -------------------------------------------------------------------------------------------------
////Реализация и Указатель
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
//
// -------------------------------------------------------------------------------------------------
