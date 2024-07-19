//--------------------BE CAREFUL !!!---------------------------------------------------------------
//
// package main

// import "fmt"

// func main() {
// 	fmt.Print("hello\n")
// 	fmt.Print("hello")
//fmt.Print("FizzBuzz", " ")
//-------------------------------------------------------------------------------------------------
// 	fmt.Printf("Наибольший общий делитель чисел %d и %d равен: %d\n", 1, 2, 3)
//-------------------------------------------------------------------------------------------------
// Бесконечный цыкл или процедура
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
//
//  НОД двух чисел
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
// package main

// import "fmt"

// func main() {
// 	for {
// 		fmt.Print("Hello", "      ")
// 	}
// }
//-------------------------------------------------------------------------------------------------
// НАХОЖДЕНИЕ ПРОСТЫХ ЧИСЕЛ
package main

import (
	"fmt"
)

func main() {
		var number int
		fmt.Scan(&number)
		fmt.Println(isPrime(number))
	}
func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}