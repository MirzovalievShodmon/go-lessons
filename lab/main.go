//КАК ПЕЧАТАТЬ? !!!---------------------------------------------------------------
//
// package main

// import "fmt"

// func main() {
//fmt.Println("Hello Go!");fmt.Println("Hello Golang!");fmt.Println("Hello Go!")

// 	fmt.Print("hello\n")

// 	fmt.Print("hello")

//fmt.Print("FizzBuzz", " ")

// 	fmt.Printf("Наибольший общий делитель чисел %d и %d равен: %d\n", 1, 2, 3)

//fmt.Printf("Имя: %s, Возраст: %d лет\n", "Shodmon", 23)

//------------------------------------------------------------------------------------

// КОММЕНТАРИИ

// Однострочный комментарий располагается в одну строку после двойного слеша . Все, что идет после этих символов, воспринимается компилятором как комментарий. Многострочный комментарий заключается между символами /* и */ и может занимать несколько строк:

/*
   Первая программа
   на языке Go
*/
// package main    // определение пакета для текущего файла
// import "fmt"    // подключение пакета fmt

// // определение функции main
// func main() {
//     fmt.Println("Hello Go!")    // вывод строки на консоль
// }

//------------------------------------------------------------------------------------

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
// -----------------------------------------------------------------------------------

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
// -----------------------------------------------------------------------------------

//БЕСКОНЕЧНЫЙ FOR
//
// package main

// import "fmt"

// func main() {
// 	for {
// 		fmt.Print("Hello", "      ")
// 	}
// }
//------------------------------------------------------------------------------------

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

//------------------------------------------------------------------------------------

////НЕ ЗАБУДЬ: дефолтное значение ДЕФОЛТНОЕ ЗНАЧЕНИЕ ДЛЯ "BOOLEAN"- ЭТО "FALSE"

// -----------------------------------------------------------------------------------

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

//------------------------------------------------------------------------------------

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
// -----------------------------------------------------------------------------------

//РАБОТА С УКАЗАТЕЛЕМ
//
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
// -----------------------------------------------------------------------------------

////Реализация и Указатель
//
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
// -----------------------------------------------------------------------------------

//ЗАДАЧА НА ТЕМУ ТИПЫ
//
// Применение функции к числу
// Определите тип функции UnaryOperation, которая принимает int и возвращает int. Напишите функции для удвоения
// и утроения числа. Используйте тип UnaryOperation для вызова этих функций.
// package main

// import "fmt"

// type UnaryOperation func(int) int

// // Функция для удвоения числа
// func double(a int) int {
// 	return a * 2
// }

// // Функция для утроения числа
// func triple(a int) int {
// 	return a * 3
// }
// func main() {
// 	var op2 UnaryOperation
// 	op2 = double
// 	fmt.Println("Double:", op2(4))
// 	op2 = triple
// 	fmt.Println("Triple:", op2(3))
// }

//------------------------------------------------------------------------------------

// Тип+int ==>> не работает.     Но псевдоним+int==>> работает
//------------------------------------------------------------------------------------

//ЗАДАЧА НА ТЕМУ ПСЕВДОНИМ
//
// Применение скидки
// Создайте псевдоним для типа float64 под названием Price.
// Напишите функцию, которая принимает Price
// и возвращает новую цену с 10% скидкой.
// package main

// import "fmt"

// type Price = float64

// func newPrice(a Price) Price {
// 	return 9 * a / 10
// }
// func main() {
// 	var a Price = 1000
// 	fmt.Println(newPrice(a))
// }

//-----------------------------------------------------------------

//Задача на тему "структура"
//
// package main

// import "fmt"

// // Person
// // Определяем структуру
// type Person struct {
// 	Name string
// 	Age  int
// }

//	func main() {
//		// Создаем и инициализируем структуру
//		p := Person{Name: "Alice", Age: 30}
//		fmt.Println(p)
//	}
//
// -----------------------------------------------------------------------------------

//     Структуры можно создавать и инициализировать несколькими способами.
// - Использование литерала структуры:

// p := Person{Name: "Alice", Age: 30}

// - Без указания имен полей (порядок важен):

// p := Person{"Alice", 30}

// - Создание структуры с использованием ключевого слова new:

// p := new(Person)
// p.Name = "Alice"
// p.Age = 30

// - Создание структуры через указатель:

// p := &Person{Name: "Alice", Age: 30}
//Обращение к полям структуры
//    Обращение к полям структуры осуществляется с помощью оператора точки (.).
// package main

// import "fmt"

// type Person struct {
//     Name string
//     Age  int
// }

// func main() {
//     p := Person{Name: "Alice", Age: 30}
//     fmt.Println(p.Name)
//     fmt.Println(p.Age)
// }
// Указатели на структуры
// Указатели на структуры позволяют изменять значения полей структуры без необходимости создания копии структуры.
// package main

// import "fmt"

// type Person struct {
//     Name string
//     Age  int
// }

// func main() {
//     p := &Person{Name: "Alice", Age: 30}
//     p.Age = 31
//     fmt.Println(p)
// }
//Вложенные структуры
//Структуры могут содержать другие структуры в качестве полей, что позволяет создавать более сложные данные.
//  package main

//  import "fmt"

//  type Address struct {
// 	 City   string
// 	 Street string
//  }
//  type Student struct {
// 	 Name    string
// 	 Age     int
// 	 Address Address
//  }

//	 func Inf(a Student) {
//	 fmt.Printf("Name:%s, Age:%d\n",a.Name, a.Age )
//	 fmt.Printf("Address:%s, %s",a.Address.City,a.Address.Street)
//	 }
//	 func main() {
//		 a := Student{
//			 Name: "Shodmon",
//			 Age:  23,
//			 Address: Address{
//				 City:   "Dushanbe",
//				 Street: "Nusratullo Makhsum",
//			 },
//		 }
//		 Inf(a)
//	 }
package main
import "fmt"
type Employee2 struct {
	Name     string
	Position string
	Manager  *Employee2
}
// Функция для рекурсивного вывода информации о сотруднике и его менеджере
func printEmployeeHierarchy(e *Employee2) {
	if e == nil {
		return
	}
	fmt.Printf("Name: %s, Position: %s\n", e.Name, e.Position)
	if e.Manager != nil {
		fmt.Printf("Manager: %s\n", e.Manager.Name)
		printEmployeeHierarchy(e.Manager)
	}
}
func main() {
	manager := &Employee2{Name: "Alice", Position: "Senior Manager"}
	employee2 := &Employee2{Name: "Bob", Position: "Junior Developer", Manager: manager}
	printEmployeeHierarchy(employee2)
}