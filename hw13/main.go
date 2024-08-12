// 1. Длина строки и Количество слов

// ⦁ Описание: Реализуйте интерфейс StringProcessor, который будетсодержать методы
//   Length() и WordCount(). Реализуйте структуру MyString, которая будет работать с
//   текстом и реализует этот интерфейс.
// ⦁ Методы:
// ⦁ Length() int для получения длины строки.
// ⦁ WordCount() int для подсчета количества слов.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type StringProcessor interface {
// 	Length() int
// 	WordCount() int
// }
// type MyString struct {
// 	String string
// }

// func (m MyString) Length() int {
// 	return len(m.String)
// }
// func (m MyString) WordCount() int {
// 	return len(strings.Fields(m.String))
// }
// func A(s StringProcessor) {
// 	fmt.Println(s.Length())
// 	fmt.Println(s.WordCount())
// }
// func B() {
// 	m := MyString{String: "Golang is a very interesting language"}
// 	A(m)
// }
// func main() {
// 	B()
// }

//------------------------------------------------------------------------------------

// 2. Форматирование строки

// ⦁ Описание: Реализуйте интерфейс Formatter с методами ToUpper() и ToLower().
//   Реализуйте структуру MyFormatter, которая будет работать со строкой и реализует этот интерфейс.
// ⦁ Методы:
// ⦁ ToUpper() string для преобразования строки в верхний регистр.
// ⦁ ToLower() string для преобразования строки в нижний регистр.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type Formatter interface {
// 	ToUpper() string
// 	ToLower() string
// }
// type MyFormatter struct {
// 	String string
// }

// func (m MyFormatter) ToUpper() string {
// 	return strings.ToUpper(m.String)
// }
// func (m MyFormatter) ToLower() string {
// 	return strings.ToLower(m.String)
// }
// func A(f Formatter) {
// 	fmt.Println(f.ToUpper())
// 	fmt.Println(f.ToLower())
// }
// func B() {
// 	m := MyFormatter{String: "Bashir studies at the HumoLab to learn Golang"}
// 	A(m)
// }
// func main() {
// 	B()
// }

//------------------------------------------------------------------------------------

// 3. Работа с указателями на числа

// ⦁ Описание: Реализуйте интерфейс PointerOperations с методами Increment() и Decrement().
//   Реализуйте структуру IntPointer с указателем на число, которая реализует этот интерфейс.
// ⦁ Методы:
// ⦁ Increment() для увеличения значения числа на 1.
// ⦁ Decrement() для уменьшения значения числа на 1.

// package main

// import "fmt"

// type PointerOperations interface {
// 	Increment()
// 	Decrement()
// }
// type IntPointer struct {
// 	intt *int
// }

// func (i IntPointer) Increment() {
// 	fmt.Println(*i.intt + 1)
// }
// func (i IntPointer) Decrement() {
// 	fmt.Println(*i.intt - 1)
// }
// func main() {
// 	number := 1
// 	v := IntPointer{
// 		intt: &number,
// 	}
// 	v.Increment()
// 	v.Decrement()
// }

//------------------------------------------------------------------------------------

// 4. Удаление пробелов в строке

// ⦁ Описание: Реализуйте интерфейс StringCleaner с методами TrimSpaces() и RemoveSpaces().
//Реализуйте структуру TextCleaner, которая будет работать со строками и реализует этот интерфейс.
// ⦁ Методы:
// ⦁ TrimSpaces() string для удаления пробелов с начала и конца строки.
// ⦁ RemoveSpaces() string для удаления всех пробелов из строки.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type StringCleaner interface {
// 	TrimSpaces() string
// 	RemoveSpaces() string
// }
// type TextCleaner struct {
// 	String string
// }

// func (t TextCleaner) TrimSpaces() string {
// 	return strings.TrimSpace(t.String)
// }
// func (t TextCleaner) RemoveSpaces() string {
// 	return strings.ReplaceAll(t.String, " ", "")
// }
// func A(s StringCleaner) {
// 	fmt.Println(s.TrimSpaces())
// 	fmt.Println(s.RemoveSpaces())
// }
// func B() {
// 	t := TextCleaner{String: "  Osh is a very tasty food with black or green tea   "}
// 	A(t)
// }
// func main() {
// 	B()
// }

//-----------------------------------------------------------------------------------

// 5. Конкатенация строк

// ⦁ Описание: Реализуйте интерфейс StringConcatenator с методами Concat() и Join().
//   Реализуйте структуру StringJoiner, которая будет работать с массивами строк и реализует этот интерфейс.
// ⦁ Методы:
// ⦁ Concat() string для конкатенации всех строк в массиве.
// ⦁ Join(separator string) string для объединения всех строк с
// заданным разделителем.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type StringConcatenator interface {
// 	Concat() string
// 	Join(separator string) string
// }
// type StringJoiner struct {
// 	slices []string
// }

// func (sj StringJoiner) Concat() string {
// 	a := ""
// 	for _, v := range sj.slices {
// 		a = a + v
// 	}
// 	return a
// }
// func (sj StringJoiner) Join(separator string) string {
// 	return strings.Join(sj.slices, separator)
// }
// func A(s StringConcatenator) {
// 	fmt.Println(s.Concat())
// 	fmt.Println(s.Join(", "))
// }
// func B() {
// 	sj := StringJoiner{slices: []string{"I", "want", "to", "learn", "Golang"}}
// 	A(sj)
// }
// func main() {
// 	B()
// }
