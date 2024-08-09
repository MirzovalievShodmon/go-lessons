// Конкатенация строк
// Напишите функцию, которая принимает две строки и возвращает их конкатенацию.

// package main

// import "fmt"

// func concatStrings(s1, s2 string) string {
// 	return s1 + s2
// }

// func main() {
// 	s1 := "Hello "
// 	s2 := "World"
// 	s3 := concatStrings(s1, s2)
// 	fmt.Println(s3)
// }

//------------------------------------------------------------------------------------

// Длина строки
// Напишите функцию, которая принимает строку и возвращает её длину.

// package main

// import "fmt"

// func S(s string) int {
// 	return len(s)
// }
// func main() {
// 	s := "Aky Shohin"
// 	fmt.Println(S(s))
// }

//-----------------------------------------------------------------------------------

// Проверка подстроки
// Напишите функцию, которая проверяет, содержит ли строка заданную подстроку.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Check(Str, str string) bool {
// 	return strings.Contains(Str, str)
// }
// func main() {
// 	fmt.Println(Check("Hello World", "World"))
// }

//-----------------------------------------------------------------------------------

// Поиск подстроки
// Напишите функцию, которая находит индекс первого вхождения подстроки в строке.
// Верните -1, если подстрока не найдена.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func FindInd(Str, str string) int {
// 	return strings.Index(Str, str)
// }
// func main() {
// 	fmt.Println(FindInd("Hello World", "World"))
// }

//-----------------------------------------------------------------------------------

// Замена подстроки
// Напишите функцию, которая заменяет все вхождения подстроки в строке на другую подстроку.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Changing(S, s, c string) string {
// 	return strings.ReplaceAll(S, s, c)
// }
// func main() {
// 	fmt.Println(Changing("Hello World", "Hello", "Asalom"))
// }

//-----------------------------------------------------------------------------------

// // Удаление пробелов
// // Напишите функцию, которая удаляет пробелы в начале и в конце строки.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Delete(s string) string {
// 	return strings.TrimSpace(s)
// }
// func main() {
// 	fmt.Println(Delete("  Asalom  Shodmon  "))
// }

//-----------------------------------------------------------------------------------

// Преобразование регистра
// Напишите функцию, которая преобразует строку к верхнему и нижнему регистру.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Aa(s string) (string, string) {
// 	s1 := strings.ToLower(s)
// 	s2 := strings.ToUpper(s)
// 	return s1, s2
// }

// func main() {
// 	fmt.Println(Aa("ShOdMon"))
// }

//-----------------------------------------------------------------------------------

// Повторение строки
// Напишите функцию, которая повторяет строку заданное количество раз.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Repeating(S string, n int) string {
// 	return strings.Repeat(S, n)
// }
// func main() {
// 	fmt.Println(Repeating("Bashir", 3))
// }

//-----------------------------------------------------------------------------------

// Разбиение строки
// Напишите функцию, которая разбивает строку по заданному разделителю и возвращает срез строк.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func DivSlice(S string, n int) []string {
// 	return strings.SplitN(S, "", n)
// }
// func main() {
// 	fmt.Println(DivSlice("Abdussabur", 6))
// }

//-----------------------------------------------------------------------------------

// // Сравнение строк
// // Напишите функцию, которая сравнивает две строки без учета регистра.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Equality(s1, s2 string) bool {
// 	return strings.EqualFold(s1, s2)
// }
// func main() {
// 	fmt.Println(Equality("ShOdMoN", "Shodmon"))
// }

//-----------------------------------------------------------------------------------

// Поиск и замена первой подстроки
// Напишите функцию, которая заменяет первое вхождение подстроки в строке на другую подстроку.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func ReplaceFirst(s, old, new string) string {
// 	return strings.Replace(s, old, new, 1)
// }

// func main() {
// 	s := "Asalom Abdusabur Shodmon"
// 	old := "Shodmon"
// 	new := "Bashir"
// 	fmt.Println(ReplaceFirst(s, old, new))
// }

//-----------------------------------------------------------------------------------

//Инверсия строки
//Напишите функцию, которая переворачивает строку.

package main

import "fmt"

func Inversion(S string) string {
	len(S) := n
	k := n / 2
	for i := 0; i < n; i++ {
		for k >= 0 {
			k--
			string(S[i]), string(S[n-i]) = string(S[n-i]), string(S[i])
		}
	}
	return S
}
func main() {
	fmt.Println(Inversion("Shodmon"))
}

//-----------------------------------------------------------------------------------

// Подсчет символов
// Напишите функцию, которая подсчитывает количество каждого символа в строке и возвращает карту с результатами.
