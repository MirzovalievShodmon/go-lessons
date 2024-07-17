// package main

// import "fmt"

// func main() {
// 	task1()
// 	task2()
// 	task3()
// 	task4()
// 	task5()
// }

// // -----------------------------------------------

// // В функции не передаются аргументы и ничего не возвращает функция
// func task1() {
// 	PrintGreeting()
// 	DisplaySeparator()
// 	NotifyStart()
// }

// // Создайте функцию PrintGreeting, которая печатает "Hello, World!" на экран.
// func PrintGreeting() {
// 	fmt.Println("Hello, World!")
// }

// // Создайте функцию DisplaySeparator, которая печатает строку из 20 символов - для разделения текста.
// func DisplaySeparator() {
// 	fmt.Println("------------------")
// }

// // Создайте функцию NotifyStart, которая выводит на экран сообщение "Program started".
// func NotifyStart() {
// 	fmt.Println("Program started")
// }

// // -----------------------------------------------

// // В функции не передаются аргументы, но функция возвращает значение (значения)
// func task2() {
// 	result1 := GetWelcomeMessage()
// 	fmt.Println(result1)

// 	result2 := GetPiValue()
// 	fmt.Println(result2)

// 	result3 := IsServerActive()
// 	fmt.Println(result3)
// }

// // Создайте функцию GetWelcomeMessage, которая возвращает строку "Welcome!".
// func GetWelcomeMessage() string {
// 	return "Welcome!"
// }

// // Создайте функцию GetPiValue, которая возвращает значение числа π с точностью
// // до двух знаков после запятой (3.14).
// func GetPiValue() float32 {
// 	return 3.14
// }

// // Создайте функцию IsServerActive, которая возвращает булево значение true.
// func IsServerActive() bool {
// 	return true
// }

// // -----------------------------------------------

// // В функции передаются аргументы, но функция ничего не возвращает
// func task3() {
// 	PrintNumber(123)
// 	GreetUser("Анвар")
// 	ToggleLight(true)
// }

// // Cоздайте функцию PrintNumber, которая принимает целое число и выводит его на экран.
// func PrintNumber(a int) {
// 	fmt.Println(a)
// }

// // Создайте функцию GreetUser, которая принимает строку с именем пользователя и выводит приветствие.
// func GreetUser(x string) {
// 	fmt.Println("Здравствуйте " + x)
// }

// // Создайте функцию ToggleLight, которая принимает булевое значение и выводит
// // "Light on" или "Light off" в зависимости от значения аргумента.
// func ToggleLight(x bool) {
// 	if x {
// 		fmt.Println("Light on")
// 	} else {
// 		fmt.Println("Light off")
// 	}
// }

// // -----------------------------------------------

// // В функции передаются аргументы, но функция возвращает значение(значения)
// func task4() {
// 	a := Add(1, 2)
// 	b := Concat("hello", "world")
// 	c := IsEven(4)

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// //Создайте функцию Add, которая принимает два целых числа и возвращает их сумму.
// func Add(a, b int) (sum int) {
// 	sum = a + b
// 	return sum
// }

// // Создайте функцию Concat, которая принимает две строки и возвращает их конкатенацию.
// func Concat(a, b string) (c string) {
// 	c = a + " " + b
// 	return c
// }

// // Создайте функцию IsEven, которая принимает целое число и возвращает true, если
// // число четное, и false в противном случае
// func IsEven(a int) (b bool) {
// 	return a%2 == 0
// }

// // -----------------------------------------------

// // Функция - как переменная
// func task5() {
// 	sum := adder(1, 3)
// 	fmt.Println(sum)

// 	b := concatenator("Hello", "Friend")
// 	fmt.Println(b)

// 	c := isPositive(-23)
// 	fmt.Println(c)

// }

// // Создайте переменную adder, которая является функцией, принимающей два целых
// // числа и возвращающей их сумму.
// func adder(a, b int) int {
// 	return a + b
// }

// // Создайте переменную concatenator, которая является функцией, принимающей две
// // строки и возвращающей их конкатенацию.
// func concatenator(a, b string) string {
// 	return a + " " + b
// }

// // Создайте переменную isPositive, которая является функцией, принимающей целое
// // число и возвращающей true, если число положительное.
// func isPositive(a int) bool {
// 	return a > 0
// }

// -----------------------------------------------
package main
func main() {
	
}