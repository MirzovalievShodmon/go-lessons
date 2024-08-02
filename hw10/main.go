// Книга и Автор
// Описание: Реализуйте структуру Book с полями title и author,
// а также метод GetDetails, который возвращает строку с деталями книги.
// Реализуйте структуру Library с массивом книг и метод AddBook,
// добавляющий книгу в библиотеку.
// Методы:
// GetDetails() string для структуры Book
// AddBook(book Book) для структуры Library

// package main

// import "fmt"

// type Book struct {
// 	title  string
// 	author string
// }

// func (b Book) GetDetails() string {
// 	return fmt.Sprintf("title: %s, author: %s", b.title, b.author)
// }

// type Library struct {
// 	books [10]Book
// }

// func (l *Library) AddBook(book Book) {
// 	l.books[0] = book
// }
// func main() {
// 	var kitob Book = Book{
// 		title:  "Gulzori hikmat",
// 		author: "Shodmon",
// 	}
// 	fmt.Println(kitob.GetDetails())
// 	var kitobkhona Library
// 	kitobkhona.AddBook(kitob)
// 	fmt.Println(kitobkhona.books[0])
// }

//------------------------------------------------------------------------------------

// Оценки студента
// Описание: Реализуйте структуру Student с полем grades (срез оценок). Реализуйте
// метод AddGrade, добавляющий оценку, и метод AverageGrade, возвращающий среднее значение оценок.
// Методы:
// AddGrade(grade int)
// AverageGrade() float64

// package main

// import (
// 	"fmt"
// )

// // Student представляет студента с оценками.
// type Student struct {
// 	grades []int
// }

// // AddGrade добавляет новую оценку в срез оценок.
// func (s *Student) AddGrade(grade int) {
// 	s.grades = append(s.grades, grade)
// }

// // AverageGrade вычисляет среднее значение оценок.
// func (s *Student) AverageGrade() float64 {
// 	if len(s.grades) == 0 {
// 		return 0
// 	}
// 	sum := 0
// 	for _, grade := range s.grades {
// 		sum += grade
// 	}
// 	return float64(sum) / float64(len(s.grades))
// }

// func main() {
// 	student := Student{}
// 	student.AddGrade(90)
// 	student.AddGrade(80)
// 	student.AddGrade(70)

// 	fmt.Printf("Средняя оценка: %.2f\n", student.AverageGrade())
// }

//------------------------------------------------------------------------------------

// Круг и Площадь
// ⦁ Описание: Реализуйте структуру Circle с полем radius. Реализуйте
// методы Area и Circumference для вычисления площади и периметра
// круга.
// ⦁ Методы:
// ⦁ Area() float64
// ⦁ Circumference() float64

// package main

// import "fmt"

// type Circle struct {
// 	radius float64
// }

// func (c Circle) Area() float64 {
// 	TheArea := 3.14 * c.radius * c.radius
// 	return TheArea
// }
// func (c Circle) Circumference() float64 {
// 	TheCircumference := 2*3.14 * c.radius
// 	return TheCircumference
// }
// func main() {
// 	TheCircle := Circle{100}
// 	fmt.Println(TheCircle.Area())
// 	fmt.Println(TheCircle.Circumference())
// }