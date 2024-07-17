// package main

// import "fmt"

// func main() {

// 	PrintGreeting()
// 	DisplaySeparator()
// 	NotifyStart()
// 	result := GetWelcomeMessage()
// 	fmt.Println(result)
// 	result1 := GetPiValue()
// 	fmt.Println(result1)
// 	result2 := IsServerActive()
// 	fmt.Println(result2)
// 	PrintNumber(1000000)
// 	GreetUser("Muhammad ibn Musa al-Khwarizmi")
// 	ToggleLight(true)
// 	result3 := Add(1, 2)
// 	fmt.Println(result3)
// 	result4 := Concat("О", "Б")
// 	fmt.Println(result4)
// 	result5 := IsEven(7)
// 	fmt.Println(result5)
// 	result6 := adder(500, 400)
// 	fmt.Println(result6)
// 	result7 := concatenator("MO", "DAR")
// 	fmt.Println(result7)
// 	result8 := isPositive(12)
// 	fmt.Println(result8)
// 	resutl9 := Calculate(1, 2, func(a, b int) int { return a + b })
// 	fmt.Println(resutl9)
// 	Execute(false, func(b bool) {
// 		if b {
// 			fmt.Println("It's day")
// 		} else {
// 			fmt.Println("It's night")
// 		}
// 	})
// }

// // 2.	В функции не передаются аргументы и ничего не возвращает функция
// // Создайте функцию PrintGreeting, которая печатает "Hello, World!" на экран.
// func PrintGreeting() {
// 	fmt.Println("Hello, World!")
// }

// // Создайте функцию DisplaySeparator, которая печатает строку из 20 символов - для разделения текста.
// func DisplaySeparator() {
// 	fmt.Println("--------------------")
// }

// // Создайте функцию NotifyStart, которая выводит на экран сообщение "Program started".
// func NotifyStart() {
// 	fmt.Println("Program started")
// }

// // В функции не передаются аргументы, но функция возвращает значение (значения)
// // Создайте функцию GetWelcomeMessage, которая возвращает строку "Welcome!".
// func GetWelcomeMessage() string {
// 	return "Welcome!"
// }

// // Создайте функцию GetPiValue, которая возвращает значение числа π с точностью до двух знаков после запятой (3.14).
// func GetPiValue() float32 {
// 	return 3.14
// }

// // Создайте функцию IsServerActive, которая возвращает булево значение true.
// func IsServerActive() bool {
// 	return true
// }

// // В функции передаются аргументы, но функция ничего не возвращает
// // Cоздайте функцию PrintNumber, которая принимает целое число и выводит его на экран.
// func PrintNumber(a int) {
// 	fmt.Println(a)
// }

// // Создайте функцию GreetUser, которая принимает строку с именем пользователя и выводит приветствие.
// func GreetUser(s string) {
// 	fmt.Println(s)
// }

// // Создайте функцию ToggleLight, которая принимает булевое значение и выводит "Light on" или "Light off" в зависимости от значения аргумента.
// func ToggleLight(b bool) {
// 	if b {
// 		fmt.Println("Light on")
// 	} else {
// 		fmt.Println("Light off")
// 	}
// }

// // В функции передаются аргументы, но функция возвращает значение(значения)
// // Создайте функцию Add, которая принимает два целых числа и возвращает их сумму.
// func Add(a, b int) int {
// 	return a + b
// }

// // Создайте функцию Concat, которая принимает две строки и возвращает их конкатенацию.
// func Concat(a, b string) string {
// 	return a + b
// }

// // Создайте функцию IsEven, которая принимает целое число и возвращает true, если число четное, и false в противном случае.
// func IsEven(a int) bool {
// 	return a%2 == 0
// }

// // Функция - как переменная
// // Создайте переменную adder, которая является функцией,
// // принимающей два целых числа и возвращающей их сумму.
// // Определение типа функции AdderFunc
// var adder func(a, b int) int = func(a, b int) int { return a + b }

// // // Создайте переменную concatenator, которая является функцией, принимающей две строки и возвращающей их конкатенацию.
// var concatenator func(a, b string) string = func(a, b string) string { return a + b }

// // Создайте переменную isPositive, которая является функцией, принимающей целое число и возвращающей true,
// // если число положительное.
// var isPositive func(n int) bool = func(n int) bool {return n > 0}

// var str string = "hello"

// // Функция - как аргумент функции
// // Создайте функцию Calculate, которая принимает два целых числа и функцию для их обработки.
// // Примените переданную функцию к этим числам и верните результат.
// func Calculate(a, b int, oldFunc func(a, b int) int) int {
// 	result:=oldFunc(a, b)
// 	return result
// }

// func newFunction()  {

// }

// // Создайте функцию Execute, которая принимает булевое значение и функцию,которая принимает
// // булевое значение и ничего не возвращает.Выполните переданную функцию с переданным булевым значением.
// func Execute(a bool, dayNight func(bool)) {
// 	dayNight(a)
// }

// -------------------------------------------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	Apply(99, func(n int) int { return n + 1 })
// }

// // Создайте функцию Apply, которая принимает целое число и функцию, которая принимает целое число
// // и возвращает целое число. Примените переданную функцию к переданному числу и верните результат.
// func Apply(n int, plusOne func(n int) int) {
// 	result:=plusOne(n)
// 	fmt.Println(result)
// }

//-------------------------------------------------------------------------------------------------
// Функция – как значение функции (callback)
// Создайте функцию Multiplier, которая принимает целое число factor и
// возвращает функцию, умножающую переданное ей число на factor.
// package main
// package main

// import "fmt"

// func main() {
// 	result := Multiplier(5)
// 	result1:=result(2)
// 	fmt.Println(result1)

// }

//	func Multiplier(factor int) func(n int) int {
//		fun := func(n int) int {
//			return factor * n
//		}
//		return fun
//	}
//
// -------------------------------------------------------------------------------------------------
// Создайте функцию StringRepeater, которая принимает целое число n и возвращает функцию,
// повторяющую переданную ей строку n раз.

// package main

// import "fmt"

// func main() {
// result:=StringRepeater(5)
//   result("McDonalds")
// }

// func StringRepeater(n int) func (s string){
//   return func(s string) {
//     for i := 0; i < n; i++ {
//       fmt.Println(s)

//	    }
//	  }
//	}
//
// -------------------------------------------------------------------------------------------------
// Создайте функцию ConditionalPrinter, которая принимает булево значение
// condition и возвращает функцию, печатающую строку, если condition истинно.
package main

import "fmt"

func main() {
	result := ConditionalPrinter(true)
	result("kfc")
}
func ConditionalPrinter(condition bool) func(s string) {
	return func(s string) {
		if condition {
			fmt.Println(s)
		}
	}
}
