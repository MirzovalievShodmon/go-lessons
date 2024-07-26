// Проверка температуры
// Определите тип Temperature на основе float64. Напишите функцию,которая принимает
// температуру и возвращает сообщение о том, является ли она ниже, выше или равной нулю.
// package main

// import "fmt"

// type Temperature float64

//	func ChekTemp(a Temperature) string {
//		if a < 0 {
//			return "Ниже нуля"
//		}
//		if a == 0 {
//			return "Равно нулю"
//		}
//		return "Больше нуля"
//	}
//
//	func main() {
//		fmt.Println(ChekTemp(25))
//	}

//-----------------------------------------------------------------------------------

// Проверка возраста
// Определите тип Age на основе int. Напишите функцию, которая
// принимает возраст и возвращает сообщение о том, является ли
// человек подростком (возраст от 13 до 19 лет включительно) или нет.
// package main

// import "fmt"

// type Age int

// func Teenager(a Age) string {
// 	if a >= 13 && a <= 19 {
// 		return "Yes, is teenager"
// 	}
// 	return "No, isn't teenager"
// }
// func main() {
// 	fmt.Println(Teenager(20))
// }

//-----------------------------------------------------------------------------------

// Проверка скорости
// Определите тип Speed на основе float64. Напишите функцию, которая принимает скорость
// и возвращает сообщение о том, является ли она допустимой (от 0 до 120 км/ч включительно) или нет.
// package main

// import "fmt"

// type Speed float64

// func AbSpeed(a Speed) string {
// 	if a >= 0 && a <= 120 {
// 		return "Да, является допустимой"
// 	}
// 	return "Нет, не яаляется допустимой"
// }
// func main() {
// 	fmt.Println(AbSpeed(60))
// }

//-----------------------------------------------------------------------------------

// Проверка счета
// Определите тип Score на основе int. Напишите функцию, которая принимает счет и
// возвращает сообщение о том, является ли он положительным, отрицательным или нулевым.
// package main

// import "fmt"

// type Score int

// func GetCount(a Score) string {
// 	if a < 0 {
// 		return "Отрицательное"
// 	}
// 	if a == 0 {
// 		return "Нулевое"
// 	}
// 	return "Положительное"
// }
// func main() {
// 	fmt.Println(GetCount(23))
// }

//-----------------------------------------------------------------------------------

// Классификация BMI
// Определите тип BMI на основе float64. Напишите функцию, которая принимает
// значение BMI и возвращает сообщение о том, в какой категории оно находится:
// "Underweight", "Normal weight", "Overweight", или "Obesity".
// package main

// import "fmt"

// type BMI float64

// func significance(a BMI) string {
// 	if a < 18.5 {
// 		return "Underweight"
// 	}
// 	if a >= 18.5 && a <= 24.9 {
// 		return "Normal weight"
// 	}
// 	if a >= 25 && a <= 29.9 {
// 		return "Overweight"
// 	}
// 	return "Obesity"
// }
// func main() {
// 	fmt.Println(significance(23))
// }

//-----------------------------------------------------------------------------------

// Возведение в квадрат
// Определите тип функции UnaryOperation, которая принимает int и возвращает int. Напишите
// функцию для возведения числа в квадрат и используйте тип UnaryOperation для вызова этой функции.
// package main

// import "fmt"

// type UnaryOperation func(int) int

// func Square(a int) int {
// 	return a * a
// }
// func main() {
// 	var mySquare UnaryOperation = Square
// 	fmt.Println(mySquare(23))
// }

//-----------------------------------------------------------------------------------

// Удвоение числа
// Определите тип функции UnaryOperation, которая принимает int и возвращает int.
// Напишите функцию для удвоения числа и используйтетип UnaryOperation для вызова этой функции.
// package main

// import "fmt"

// type UnaryOperation func(int) int

// func Twise(a int) int {
// 	return 2 * a
// }
// func main() {
// 	var myTwise UnaryOperation = Twise
// 	fmt.Println(myTwise(23))
// }

//-----------------------------------------------------------------------------------

// Проверка четности
// Определите тип функции Check, которая принимает int и возвращает
// bool. Напишите функцию для проверки четности числа и используйте
// тип Check для вызова этой функции.
// package main

// import "fmt"

// type Check func(int) bool

// func Chetnost(a int) bool {
// 	if a%2 == 0 {
// 		return true
// 	}
// 	return false
// }
// func main() {
// 	var myChetnost Check = Chetnost
// 	if myChetnost(23) == true {
// 		fmt.Println("Число является четным")
// 	} else {
// 		fmt.Println("Число является нечетным")
// 	}
// }

//-----------------------------------------------------------------------------------

// Проверка положительности
// Определите тип функции Check, которая принимает int и возвращает bool. Напишите функцию
//для проверки, является ли число положительным, и используйте тип Check для вызова этой функции.

// package main

// import "fmt"

// type Check func(int) bool

// func NumberIs(a int) bool {
// 	if a > 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// func main() {
// 	var TheNumber Check = NumberIs
// 	if TheNumber(23) == true {
// 		fmt.Println("Число является положительным")
// 	} else {
// 		fmt.Println("Число не является положительным")
// 	}
// }

//-----------------------------------------------------------------------------------

// Обратный отсчет
// Создайте псевдоним для типа int под названием Count. Напишите функцию, которая
// принимает Count и выводит обратный отсчет до нуля.этих функций.
// package main

// import "fmt"

// type Count = int

// func MyCount(a Count) {
// 	for i := a; i >= 0; i-- {
// 		fmt.Print(i, " ")
// 	}
// }
// func main() {
// 	MyCount(10)
// }

//-----------------------------------------------------------------------------------

// Проверка уровня батареи
// Создайте псевдоним для типа int под названием BatteryLevel. Напишите функцию, которая
// принимает BatteryLevel и возвращает сообщение о том, низкий, средний или высокий уровень заряда.
// package main

// import "fmt"

// type BatteryLevel = int

// func Level(a BatteryLevel) string {
// 	if a < 30 {
// 		return "Низкий уровень заряда"
// 	}
// 	if a >= 30 && a < 70 {
// 		return "Средний уровень заряда"
// 	} else {
// 		return "Высокий уровень заряда"
// 	}
// }
// func main() {
// 	var a BatteryLevel = 80
// 	fmt.Println(Level(a))
// }

//-----------------------------------------------------------------------------------

// Определение категории веса
// Создайте псевдоним для типа float64 под названием Weight. Напишите функцию, которая
// принимает Weight и возвращает сообщение о том, в какой категории веса находится объект: "Light", "Medium" или "Heavy".
// package main

// import "fmt"

// type Weight = float64

// func PrintWeight(a Weight) string {
// 	if a >= 50 && a < 65 {
// 		return "Light"
// 	}
// 	if a >= 65 && a < 80 {
// 		return "Medium"
// 	}
// 	if a >= 80 {
// 		return "Heavy"
// 	}
// 	return ""
// }
// func main() {
// 	var a Weight
// 	a = 65
// 	fmt.Println(PrintWeight(a))
// }

//-----------------------------------------------------------------------------------

// Оценка успеваемости
// Создайте псевдоним для типа int под названием Grade. Напишите функцию, которая принимает
// Grade и возвращает сообщение о том, является ли оценка удовлетворительной (50 и выше) или нет.
// package main

// import "fmt"

// type Grade = int

// func Mark(a Grade) string {
// 	if a >= 50 {
// 		return "Оценка удовлетворительная"
// 	}
// 	return "Оценка неудовлетворительная"
// }
// func main() {
// 	var a Grade = 63
// 	fmt.Println(Mark(a))
// }

//-----------------------------------------------------------------------------------

// Оценка состояния здоровья
// Создайте псевдонимы для типов float64 под названиями BMI и
// BloodPressure. Напишите функцию, которая принимает BMI и
// BloodPressure, и возвращает сообщение о состоянии здоровья:
// "Healthy", "At risk" или "Unhealthy".
// package main
// type BMI=float64
// type BloodPressure=float64
// func How( BMI, BloodPressure) string{
// 	if
// }
// func main() {

// }

//-----------------------------------------------------------------------------------

// Описание продукта
// Создайте структуру Product с полями Name и Price. Напишите функцию, которая принимает Product и выводит его описание.
// package main

// import "fmt"

// type Product struct {
// 	Name  string
// 	Price float64
// }

// func MyProduct(p Product) {
// 	fmt.Printf("Name: %s, Price: %.2f", p.Name, p.Price)
// }
// func main() {
// 	p := Product{Name: "Cake", Price: 1000}
// 	MyProduct(p)
// }

//-----------------------------------------------------------------------------------

// Вывод информации о человеке
// Создайте структуру Person с полями FirstName, LastName и Age. Напишите функцию,
// которая принимает Person и выводит его полное имя и возраст.
// package main

// import "fmt"

// type Person struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

// func Describing(p Person) {
// 	fmt.Printf("FirstName: %s\nLastName: %s\nAge: %d", p.FirstName, p.LastName, p.Age)
// }
// func main() {
// 	p := Person{
// 		FirstName: "Shodmon",
// 		LastName:  "Mirzovaliev",
// 		Age:       23,
// 	}
// 	Describing(p)
// }

//-----------------------------------------------------------------------------------

// Сравнение продуктов
// Создайте структуру Product с полями Name и Price. Напишите
// функцию, которая принимает два Product и возвращает более дорогой продукт.
// package main

// import "fmt"

// type Product struct {
// 	Name  string
// 	Price float64
// }

// func Expensive(p1, p2 Product) {
// 	if p1.Price > p2.Price {
// 		fmt.Printf("Name: %s\nPrice: %.2f", p1.Name, p1.Price)
// 	} else if p1.Price < p2.Price {
// 		fmt.Printf("Name: %s\nPrice: %.2f", p2.Name, p2.Price)
// 	} else {
// 		fmt.Println("У обоих продуктов одинаковая цена")
// 	}
// }
// func main() {
// 	p1 := Product{
// 		Name:  "Оби Зулол",
// 		Price: 10,
// 	}
// 	p2 := Product{
// 		Name:  "Шохамбари",
// 		Price: 9,
// 	}
// 	Expensive(p1, p2)
// }

//-----------------------------------------------------------------------------------

// Обновление цены продукта
// Создайте структуру Product с полями Name и Price. Напишите функцию, которая
// принимает указатель на Product и обновляет его цену.
// package main

// import "fmt"

// type Product struct {
// 	Name  string
// 	Price float64
// }

// func NewPrice(a *Product) {
// 	a.Price=a.Price-100
// 	fmt.Printf("Name:%s\nPrice:%.2f", a.Name, a.Price)
// }
// func main() {
// 	var a = &Product{
// 		Name:  "Laptop",
// 		Price: 6500,
// 	}
// 	NewPrice(a)
// }

//-----------------------------------------------------------------------------------

// Увеличение возраста
// Создайте структуру Person с полями Name и Age. Напишите функцию,
// которая принимает указатель на Person и увеличивает его возраст на 1.
// package main

// import "fmt"

// type Person struct {
// 	Name string
// 	Age  int
// }

// func OtherAge(a *Person) {
// 	a.Age = a.Age - 1
// 	fmt.Printf("Name:%s\nAge:%d", a.Name, a.Age)
// }
// func main() {
// 	var a = &Person{
// 		Name: "Shodmon",
// 		Age:  24,
// 	}
// 	OtherAge(a)
// }