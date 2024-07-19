// Создайте функцию PrintGreeting, которая печатает "Доброе утро!", если dayType равен "утро";
//	"Добрый день!", если dayType равен "день"; и "Добрый вечер!", если dayType равен "вечер".
//	Переменную  dayType вводить с консоли внутри функции.
// package main

// import "fmt"

//	func main() {
//		PrintGreeting()
//	}
//	func PrintGreeting(){
//		var dayType string
//		fmt.Scanln(&dayType)
//		if dayType=="утро"{
//			fmt.Println("Доброе утро!")
//		}
//		if dayType=="день"{
//			fmt.Println("Добрый день!")
//		}
//		if dayType=="вечер"{
//			fmt.Println("Добрый вечер!")
//		}
//	}
//-------------------------------------------------------------------------------------------------
// Создайте функцию PrintWeather, которая печатает "Солнечно", если
// weatherType равен "солнечно"; "Облачно", если weatherType равен
// "облачно"; и "Дождливо", если weatherType равен "дождливо".
// Переменную weatherType вводить с консоли внутри функции.

// package main

// import "fmt"

//	func main() {
//		PrintWeather()
//	}
//
//	func PrintWeather() {
//		var weatherType string
//		fmt.Scanln(&weatherType)
//		if weatherType == "солнечно" {
//			fmt.Println("Солнечно")
//		}
//		if weatherType == "облачно" {
//			fmt.Println("Облачно")
//		}
//		if weatherType == "дождливо" {
//			fmt.Println("Дождливо")
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Создайте функцию PrintTrafficLight, которая печатает
// "Стоп", если lightColor равен "красный";
// "Внимание", если lightColor равен "желтый"; и
// "Идите", если lightColor равен "зеленый".
// Переменную lightColor вводить с консоли внутри функции.
// package main

// import "fmt"

//	func main() {
//		PrintTrafficLight()
//	}
//
//	func PrintTrafficLight() {
//		var lightColor string
//		fmt.Scanln(&lightColor)
//		if lightColor == "красный" {
//			fmt.Println("Стоп")
//		}
//		if lightColor == "желтый" {
//			fmt.Println("Внимание")
//		}
//		if lightColor == "зеленый" {
//			fmt.Println("Идите")
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Создайте функцию GetGrade, которая возвращает оценку "A", "B" или
// "C" в зависимости от значения переменной score. Переменную scope
// вводить с консоли внутри функции.
// package main

// import "fmt"

// func main() {
// fmt.Println(GetGrade())
// }
//
//	func GetGrade() string {
//	  var score int
//	  fmt.Scan(&score)
//	   var grade string
//	   switch {
//	   case score >= 90:
//	     grade = "A"
//	   case score >= 75 && score < 90:
//	     grade = "B"
//	   case score < 75:
//	     grade = "C"
//	   }
//	   return grade
//	 }
//
// // -------------------------------------------------------------------------------------------------
// Создайте функцию GetDiscount, которая возвращает скидку "10%" -
// при amount > 1000, "5%" - при 500 < amount <=100 или "0%" - при
// amount <= 500 в зависимости от суммы покупки amount. Переменную
// amount вводить с консоли внутри функции.
package main

import "fmt"

func main() {
	fmt.Println(GetDiscount())
}
func GetDiscount() string {
	var amount int
	fmt.Scan(&amount)
	var discount string
	switch {
	case amount >= 1000:
		discount = "10%"
	case amount > 500 && amount <= 1000:
		discount = "5%"
	case amount <= 500:
		discount = "0%"
	}
	return discount
}
