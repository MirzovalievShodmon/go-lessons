// Напишите программу, которая используя конструкцию циклов выводит
// числа от 1 до 10.
// package main

// import "fmt"

//	func main() {
//		var i int
//		for i = 1; i <= 9; i++ {
//			fmt.Print(i, " ")
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Напишите программу, которая используя конструкцию циклов выводит
// квадраты чисел от 1 до 5.
// package main

// import "fmt"

//	func main() {
//		var i int
//		for i = 1; i <= 5; i++ {
//			fmt.Print(i * i, " ")
//		}
//	}
//-------------------------------------------------------------------------------------------------
// Напишите программу, которая выводит таблицу умножения на 3 от 1
// до 10.
// // package main

// import "fmt"

// func main() {
// 	var i int
// 	for i = 1; i <= 10; i++ {
// 		fmt.Println("3 х", i, "=", 3*i)
// 	}
// }
//-------------------------------------------------------------------------------------------------
// Напишите программу, которая выводит первые 10 чисел
// последовательности Фибоначчи.
// package main

// import "fmt"

//	func main() {
//		var i, f1, f2, f3 int
//		f1, f2 = 1, 1
//		fmt.Print(f1, " ", f2, " ")
//		for i = 3; i <= 10; i++ {
//			f3 = f1 + f2
//			f1 = f2
//			f2 = f3
//			fmt.Print(f3, " ")
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Напишите программу, которая находит наибольший общий делитель
// (НОД) двух чисел, используя алгоритм Евклида.
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
//		fmt.Println(a)
//	}
//
// -------------------------------------------------------------------------------------------------
// Напишите программу, которая выводит числа от 1 до 100, заменяя
// числа, кратные 3, на "Fizz", числа, кратные 5, на "Buzz", а числа,
// кратные 3 и 5, на "FizzBuzz".
// package main

// import "fmt"

// var i int

//	func main() {
//		for i = 1; i <= 100; i++ {
//			if i%3 == 0 && i%5 != 0 {
//				fmt.Print("Fizz", " ")
//			} else if i%5 == 0 && i%3 != 0 {
//				fmt.Print("Buzz", " ")
//			} else if i%15 == 0 {
//				fmt.Print("FizzBuzz", " ")
//			} else {
//				fmt.Print(i, " ")
//			}
//		}
//	}
//-------------------------------------------------------------------------------------------------
// Напишите программу, которая проверяет, является ли число простым.
// package main

// import (
// 	"fmt"
// )

// func main() {
// 		var number int
// 		fmt.Scan(&number)
// 		fmt.Println(isPrime(number))
// 	}
// func isPrime(n int) bool {
// 	if n == 1 {
// 		return false
// 	}
// 	for i := 2; i*i <= n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
//-------------------------------------------------------------------------------------------------
// Напишите программу, которая выводит все делители числа.
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var number int
// 	fmt.Print("Введите число: ")
// 	fmt.Scan(&number)

//		fmt.Printf("Делители числа %d: ", number)
//		for i := 1; i <= number; i++ {
//			if number%i == 0 {
//				fmt.Printf("%d ", i)
//			}
//		}
//	}
//-------------------------------------------------------------------------------------------------
// Напишите программу, которая находит сумму цифр числа.
// package main

// import "fmt"

//	func main() {
//		var a, b, S int
//		fmt.Scan(&a)
//		for a != 0 {
//			b = a % 10
//			a = a / 10
//			S = S + b
//		}
//		fmt.Println(S)
//	}
//-------------------------------------------------------------------------------------------------
// Напишите программу, которая запрашивает у пользователя ввод
// положительного числа и повторяет запрос, пока не будет введено
// положительное число.
// package main

// import "fmt"

//	func main() {
//		var a int
//		fmt.Scan(&a)
//		for a <= 0 {
//			fmt.Print("Введите положительное число", " ")
//			fmt.Scan(&a)
//		}
//	}
//-------------------------------------------------------------------------------------------------
// Напишите программу, которая вычисляет произведение всех чисел от 1
// до введённого числа n, но прекращает выполнение, если результат
// превышает 1000.
// package main

// import "fmt"

//	func main() {
//		var n, a int
//		p := 1
//		fmt.Scan(&n)
//		for i := 1; i <= n; i++ {
//			if p > 1000 {
//				break
//			}
//			a = p
//			p = p * i
//		}
//		if p > 1000 {
//			fmt.Println(a)
//		} else {
//			fmt.Println(p)
//		}
//	}
//-------------------------------------------------------------------------------------------------
// Напишите программу, которая запрашивает у пользователя ввод числа
// и проверяет, является ли оно палиндромом.
package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Scan(&a)

	fmt.Println(Palindrome(a))
}

func Palindrome(a int) bool {
	var b, c int
	p := 0
	c = a
	for a != 0 {
		b = a % 10
		a = a / 10
		p = p*10+b
	}

	return c == p
}
