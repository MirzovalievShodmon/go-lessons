// Определение возраста совершеннолетия
// Определите тип Age на основе int. Напишите функцию, которая принимает возраст и возвращает сообщение о том,
// является ли человек совершеннолетним (возраст 18 лет и старше) или нет.
// package main

// type Age int

//	func main() {
//		var myAge Age = 17
//		vozrast(myAge)
//	}
//
//	func vozrast(myAge Age) string {
//		if myAge >= 18 {
//			return "Да является совершеннолетним"
//		} else {
//			return "Нет не является совершеннолетним"
//		}
//	}
//-------------------------------------------------------------------------------------------------
// Проверка на четность
// Определите тип Number на основе int. Напишите функцию, которая принимает число и возвращает сообщение о том,
// является ли оно четным или нечетным.
// package main

// type Number int

//	func main() {
//		var chislo Number = 56
//		chetnost(chislo)
//	}
//
//	func chetnost(chislo Number) string {
//		if chislo%2 == 0 {
//			return "Да является четным"
//		} else {
//			return "Нет является четным"
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Проверка допустимого диапазона
// Определите тип Score на основе int. Напишите функцию, которая принимает оценку и возвращает сообщение,
// находится ли она в допустимом диапазоне от 0 до 100
// package main

// type Score int

// func main() {
// 	var num Score = 101
// 	dip(num)
// }

//	func dip(num Score) string {
//		if num >= 0 && num <= 100 {
//			return "Да находится в этом диапазоне"
//		} else {
//			return "Нет не находится в этом диапазоне"
//		}
//	}
//
// -------------------------------------------------------------------------------------------------
// Арифметические операции
// Определите тип функции Operation, которая принимает два int и возвращает int.
// Напишите функции для сложения,
// вычитания и умножения. Используйте тип Operation для вызова этих функций
// package main

// import "fmt"

// type Operation func(int, int) int

// func add(a int, b int) int {
// 	return a + b
// }
// func subtruct(a int, b int) int {
// 	return a - b
// }
// func multiple(a int, b int) int {
// 	return a * b
// }

//	func main() {
//		var myAdder Operation = add
//		result1 := myAdder(2, 3)
//		fmt.Println(result1)
//		var mySubstruct Operation = subtruct
//		result2 := mySubstruct(2, 3)
//		fmt.Println(result2)
//		var myMultyple Operation = multiple
//		result3 := myMultyple(2, 3)
//		fmt.Println(result3)
//	}

// -------------------------------------------------------------------------------------------------

// Сравнение чисел
// Определите тип функции Comparator, которая принимает два int и возвращает bool.
// Напишите функции для проверки равенства и сравнения больше/меньше.
// Используйте тип Comparator для вызова этих функций.
// package main

// import "fmt"

// type Comparator func(int, int) bool

// 	func equality(a int, b int) bool {
// 		return a == b
// 	}

// 	func greater(a int, b int) bool {
// 		return a > b
// 	}

// 	func smaller(a int, b int) bool {
// 		return a < b
// 	}

//	func main() {
//		var myEquality Comparator = equality
//		result1 := myEquality(2, 3)
//		fmt.Println(result1)
//		var myGreater Comparator = greater
//		result2 := myGreater(2, 3)
//		fmt.Println(result2)
//		var mySmaller Comparator = smaller
//		result3 := mySmaller(2, 3)
//		fmt.Println(result3)
//	}

//-------------------------------------------------------------------------------------------------

// Применение функции к числу
// Определите тип функции UnaryOperation, которая принимает int и возвращает int. Напишите функции для удвоения
// и утроения числа. Используйте тип UnaryOperation для вызова этих функций.
// package main

// import "fmt"

// type UnaryOperation func(int) int

//	func twise(a int) int {
//		return a*2
//	}
//
//	func three(a int) int {
//		return a*3
//	}
//
//	func main() {
//		var myTwise UnaryOperation = twise
//		result1 := myTwise(2)
//		fmt.Println(result1)
//		var myThree UnaryOperation = three
//		result2 := myThree(2)
//		fmt.Println(result2)
//	}
//
// -------------------------------------------------------------------------------------------------

// Обратный отсчет
// Создайте псевдоним для типа int под названием Count.
// Напишите функцию, которая принимает Count и выводит
// обратный отсчет до нуля.
// package main

// import "fmt"

// type Count = int

//	func myCount(a Count) {
//		for i := a; i >= 0; i-- {
//			fmt.Print(i," ")
//		}
//	}
//
//	func main() {
//		var a Count = 10
//		myCount(a)
//	}

//-------------------------------------------------------------------------------------------------

// Проверка температуры
// Создайте псевдоним для типа float64 под названием Temperature.
// Напишите функцию, которая принимает
// Temperature и выводит сообщение о состоянии (ниже нуля, выше нуля или ноль).
// package main

// import "fmt"

// type Temperature = float64

//	func thisTemp(a Temperature) {
//		if a < 0 {
//			fmt.Println("ниже нуля")
//		}else if a == 0 {
//			fmt.Println("ноль")
//		} else {
//			fmt.Println("выше нуля")
//		}
//	}
//
//	func main() {
//		var a Temperature = 10
//		thisTemp(a)
//	}
//
// -------------------------------------------------------------------------------------------------
// Применение скидки
// Создайте псевдоним для типа float64 под названием Price.
// Напишите функцию, которая принимает Price
// и возвращает новую цену с 10% скидкой.
// package main

// import "fmt"

// type Price = float64

//	func newPrice(a Price) Price {
//		return 9 * a / 10
//	}
//
//	func main() {
//		var a Price = 1000
//		fmt.Println(newPrice(a))
//	}

// ------------------------------------------------------------------------------------

// Информация о пользователе
// Создайте структуру User с полями Name (строка) и Age (целое число).
// Напишите функцию, которая принимает User и выводит информацию о нем.
// package main

// import "fmt"
//
//	type User struct{
//		Name string
//		Age int
//	}
//
//	func main() {
//		p := User{Name: "Shodmon", Age: 23}
//		fmt.Println(p)
//	}

// ------------------------------------------------------------------------------------

// Продукт и его стоимость
// Создайте структуру Product с полями Name (строка) и Price (тип Price).
// Напишите функцию, которая принимает Product и возвращает сообщение о его стоимости.
// package main

// import "fmt"

// type Price float64
// type Product struct {
// 	Name  string
// 	Price Price
// }

// func Buy(p Product) {
// 	fmt.Printf("Product:%s, Price: %.2f", p.Name, p.Price)
// }
// func main() {
// 	p := Product{Name: "Tasty Juice", Price: 1000}
// 	Buy(p)
// }

//------------------------------------------------------------------------------------

// Информация о книге
// Создайте структуру Book с полями Title (строка), Author (строка) и Pages (целое число).
// Напишите функцию, которая принимает Book и выводит информацию о книге.
// package main

// import "fmt"

// type Book struct {
// 	Title  string
// 	Author string
// 	Pages  int
// }

// func InfBook(b Book) {
// 	fmt.Printf("Title: %s,  Author: %s,  Pages: %d\n", b.Title, b.Author, b.Pages)
// }
// func main() {
// 	b := Book{Title:"Gulzory Hikmat", Author: "Shodmon", Pages: 1234}
// 	InfBook(b)
// }

//------------------------------------------------------------------------------------

// Изменение данных через указатель
// Создайте структуру Person с полями Name и Age. Напишите функцию, которая принимает
// указатель на Person и обновляет его возраст. Выведите информацию до и после обновления.
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func ptrPerson(p *Person) {

}
func main() {
	p := &Person{"Shodmon", 23}
	fmt.Println(*p)
	p.Name = "Abdussabur"
	p.Age = 25
	ptrPerson(p)
	fmt.Println(*p)
}

//-----------------------------------------------------------------------------------

// Создание и изменение структуры через указатель
// Создайте структуру Rectangle с полями Width и Height. Напишите функцию,
// которая принимает указательна Rectangle, вычисляет и обновляет его площадь,
// а затем выводит обновленную площадь.
// package main

// import "fmt"

// type Rectangle struct {
// 	Width  float64
// 	Height float64
// }

// func prtArea(a *Rectangle) {
// 	area := a.Width * a.Height
// 	fmt.Printf("Area:  %.2f\n", area)
// }

// func main() {
// 	a := &Rectangle{Width: 10, Height: 20}
// 	prtArea(a)
// 	a.Width = 15
// 	a.Height = 20
// 	prtArea(a)
// }

//-----------------------------------------------------------------------------------

// Сравнение двух структур через указатели
// Создайте структуру Coordinate с полями X и Y. Напишите функцию, которая принимает
// указатели на две Coordinate и возвращает сообщение о том, равны ли они или нет.
// package main

// import "fmt"

// type Coordinate struct {
// 	X float64
// 	Y float64
// }

//	func MCoor(c1, c2 *Coordinate) string {
//		if c1.X == c2.X && c1.Y == c2.Y {
//			return "Они равны"
//		}
//		return "Они не равны"
//	}
//
//	func main() {
//		c1 := &Coordinate{X: 10, Y: 15}
//		c2 := &Coordinate{X: 15, Y: 20}
//		fmt.Println(MCoor(c1, c2))
//	}

//-----------------------------------------------------------------------------------

// Программа для управления студентом
// Создайте структуры Student и Address. Student должен содержать поля Name,
// Age и Address (вложенная структура Address). Напишите функцию,
// которая выводит информацию о студенте, включая адрес.
// package main

// import "fmt"

// type Address struct {
// 	City   string
// 	Street string
// }
// type Student struct {
// 	Name    string
// 	Age     int
// 	Address Address
// }

// func Inf(a Student) {
// fmt.Printf("Name:%s, Age:%d\n",a.Name, a.Age )
// fmt.Printf("Address:%s, %s",a.Address.City,a.Address.Street)
// }
// func main() {
// 	a := Student{
// 		Name: "Shodmon",
// 		Age:  23,
// 		Address: Address{
// 			City:   "Dushanbe",
// 			Street: "Nusratullo Makhsum",
// 		},
// 	}
// 	Inf(a)
// }

//-----------------------------------------------------------------------------------

// Работа с сотрудниками
// Создайте структуры Employee и Department. Employee должен содержать поля Name,
// ID и Department (вложенная структура Department). Напишите функцию,
// которая выводит информацию о сотруднике и его отделе.
// package main

// import "fmt"

// type Department struct {
// 	Name string
// }
// type Employee struct {
// 	Name       string
// 	ID         string
// 	Department Department
// }

// func InfE(a Employee) {
// 	fmt.Printf("Name: %s,  ID: %s\n", a.Name, a.ID)
// 	fmt.Printf("Department: %s", a.Department.Name)
// }
// func main() {
// 	a := Employee{
// 		Name: "Shodmon",
// 		ID:   "A0B1C2D4E5",
// 		Department: Department{
// 			Name: "Golang Developer",
// 		},
// 	}
// 	InfE(a)
// }

//-----------------------------------------------------------------------------------

// Дерево сотрудников
// Создайте структуру Employee с полем Manager, которое указывает на другого
// сотрудника (вложенная структура). Напишите функцию, которая рекурсивно
// выводит информацию о сотруднике и его менеджере.
// package main

// import (
// 	"fmt"
// )

// type Employee struct {
// 	Name     string
// 	Position string
// 	Manager  *Employee
// }

// func RecInfo(a *Employee) {
// 	if a == nil {
// 		return
// 	}
// 	fmt.Printf("Name:%s, Position:%s\n", a.Name, a.Position)
// 	if a.Manager != nil {
// 		fmt.Printf("Manager:%s\n", a.Manager.Name)
// 		RecInfo(a.Manager)
// 	}

// }
//
//	func main() {
//		manager := &Employee{Name: "Abdulhalim", Position: "Senior Developer"}
//		employee := &Employee{Name: "Shodmon", Position: "Junior Developer", Manager: manager}
//		RecInfo(employee)
//	}

// -----------------------------------------------------------------------------------

// Временная структура для хранения информации о книге
// Создайте анонимную структуру для хранения информации о книге с полями Title и Author.
// Напишите функцию, которая принимает эту анонимную структуру и выводит информацию о книге.

// package main

// import "fmt"

// func printBookInfo(book struct {
// 	Title  string
// 	Author string
// }) {
// 	fmt.Printf("Title: %s\nAuthor: %s", book.Title, book.Author)
// }
// func main() {
// 	book := struct {
// 		Title  string
// 		Author string
// 	}{
// 		Title:  "Не грусти! Рецепты счастья и лекарство от груст",
// 		Author: "Аид аль-Карни",
// 	}
// 	printBookInfo(book)
// }

//-----------------------------------------------------------------------------------

// Временная структура для хранения информации о продукте
// Создайте анонимную структуру для хранения информации о продукте с полями Name и Price.
//Напишите функцию, которая принимает эту анонимную структуру и возвращает строку с описанием продукта.

// package main

// import "fmt"

// func PrintProductInfo(Product struct {
// 	Name  string
// 	Price float64
// }) string {
// 	return fmt.Sprintf("Name:%s\nPrice:%.2f", Product.Name, Product.Price)
// }
// func main() {
// 	Product := struct {
// 		Name  string
// 		Price float64
// 	}{
// 		Name:  "Laptop",
// 		Price: 6500,
// 	}
// 	fmt.Println(PrintProductInfo(Product))
// }

//-----------------------------------------------------------------------------------

// Временная структура для хранения информации о событии
// Создайте анонимную структуру для хранения информации о событии с полями EventName и Date.
// Напишите функцию, которая принимает эту анонимную структуру и возвращает строку с описанием события.
// package main

// import "fmt"

// func TheEvent(Event struct {
// 	EventName string
// 	Date      string
// }) string {
// 	return fmt.Sprintf("EventName: %s\nDate: %s", Event.EventName, Event.Date)
// }
// func main() {
// 	Event := struct {
// 		EventName string
// 		Date      string
// 	}{
// 		EventName: "The HumoLab's jubilee",
// 		Date:      "27.07.2024",
// 	}
// 	fmt.Println(TheEvent(Event))
// }