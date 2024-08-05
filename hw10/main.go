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
//  Описание: Реализуйте структуру Circle с полем radius. Реализуйте
// методы Area и Circumference для вычисления площади и периметра
// круга.
//  Методы:
//  Area() float64
//  Circumference() float64

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

//------------------------------------------------------------------------------------

// Контейнер для товаров
// Описание: Реализуйте структуру Item с полями name и price.
//Реализуйте структуру Inventory с срезом товаров и методы для добавления товара
//и получения общего значения товаров в инвентаре.
// Методы:
// AddItem(item Item)
// TotalValue() float64

// package main

// import "fmt"

// // Item представляет товар с полями name и price.
// type Item struct {
// 	Name  string
// 	Price float64
// }

// // Inventory представляет инвентарь с срезом товаров.
// type Inventory struct {
// 	Items []Item
// }

// // AddItem добавляет товар в инвентарь.
// func (inv *Inventory) AddItem(item Item) {
// 	inv.Items = append(inv.Items, item)
// }
// func (inv *Inventory) TotalValue() float64 {
// 	total := 0.0
// 	for _, item := range inv.Items {
// 		total += item.Price
// 	}
// 	return total
// }

// func main() {
// 	// Создание инвентаря.
// 	inventory := Inventory{}
// 	// Добавление товаров.
// 	inventory.AddItem(Item{Name: "Laptop", Price: 999.99})
// 	inventory.AddItem(Item{Name: "Smartphone", Price: 499.99})
// 	inventory.AddItem(Item{Name: "Keyboard", Price: 29.99})
// 	// Вывод общего значения товаров в инвентаре.
// 	fmt.Printf("Общее значение товаров в инвентаре: %.2f\n", inventory.TotalValue())
// }

//------------------------------------------------------------------------------------

// Обработка транзакций
// ⦁ Описание: Реализуйте структуру Transaction с полями amount и
// description. Реализуйте метод Process, который выводит информацию
// о транзакции. Реализуйте структуру Account с методом AddTransaction.
// ⦁ Методы:
// ⦁ Process()
// ⦁ AddTransaction(amount float64, description string)

// package main

// import "fmt"

// type Transaction struct {
// 	amount      float64
// 	description string
// }

// type Account struct {
// 	transactions []Transaction
// }

// // Метод для вывода информации о транзакции
// func (t Transaction) Process() {
// 	fmt.Printf("amount:%.2f\n description: %s", t.amount, t.description)
// }

// // Метод для добавления транзакции в аккаунт
// func (a *Account) AddTransaction(amount float64, description string) {
// 	transaction := Transaction{amount: amount, description: description}
// 	a.transactions = append(a.transactions, transaction)
// }

// func main() {
// 	// Создание и обработка транзакций
// 	tx1 := Transaction{amount: 1000, description: "Покупка в магазине HUMO"}
// 	tx2 := Transaction{amount: 1500, description: "Покупка в магазине FENIX"}
// 	tx1.Process()
// 	tx2.Process()
// 	// Создание аккаунта и добавление транзакций
// 	TheAccount := Account{}
// 	TheAccount.AddTransaction(tx1.amount, tx1.description)
// 	TheAccount.AddTransaction(tx2.amount, tx2.description)

// 	transctions := TheAccount.transactions

// 	// Вывод всех транзакций аккаунта
// 	// for _, t := range TheAccount.transactions {
// 	// 	t.Process()
// 	// }

// 	for _, t := range transctions {
// 		t.Process()
// 	}
// }

//------------------------------------------------------------------------------------

// Управление задачами
// ⦁ Описание: Реализуйте структуру Task с полями title и completed.
// Реализуйте структуру TaskList с срезом задач и методы для добавления
// задачи и получения количества завершённых задач.
// ⦁ Методы:
// ⦁ AddTask(title string)
// ⦁ CompletedTasks() int

// package main

// import "fmt"

// type Task struct {
// 	title     string
// 	completed bool
// }
// type TaskList struct {
// 	tasks []Task
// }

// func (t *TaskList) AddTasks(task []Task) {
// 	t.tasks = append(t.tasks, task...)
// }

// func (t *TaskList) CompletedTasks() int {
// 	k := 0
// 	for _, v := range t.tasks {
// 		if v.completed {
// 			k++
// 		}
// 	}
// 	return k
// }
// func main() {
// 	task := []Task{
// 		{
// 			title:     "Покупка 1продуктов",
// 			completed: true,
// 		},
// 		{
// 			title:     "Покупка 2продуктов",
// 			completed: true,
// 		},
// 		{
// 			title:     "Покупка 3продуктов",
// 			completed: false,
// 		},
// 	}
// 	a := TaskList{}
// 	a.AddTasks(task)
// 	fmt.Println(a.tasks)
// 	a.CompletedTasks()
// 	fmt.Println(a.CompletedTasks())
// }

//------------------------------------------------------------------------------------

// Работа с температурой
// ⦁ Описание: Реализуйте структуру Temperature с полем celsius.
// Реализуйте методы для преобразования температуры в Фаренгейт и
// Кельвин.
// ⦁ Методы:
// ⦁ ToFahrenheit() float64
// ⦁ ToKelvin() float64

// package main

// import "fmt"

// type Temperature struct {
// 	celsius float64
// }

// func (t Temperature) ToFahrenheit() float64 {
// 	return t.celsius*9/5 + 32
// }
// func (t Temperature) ToKelvin() float64 {
// 	return t.celsius + 273.5
// }
// func main() {
// 	Temp := Temperature{
// 		celsius: 10,
// 	}
// 	fmt.Println("Temperature in Fahrenheit:", Temp.ToFahrenheit())
// 	fmt.Println("Temperature in Kelvin:", Temp.ToKelvin())
// }

//------------------------------------------------------------------------------------

// Управление списком студентов
// ⦁ Описание: Реализуйте структуру Student с полем name и age.
// Реализуйте структуру Classroom с срезом студентов и методы для
// добавления студента и получения средней возрастной группы.
// ⦁ Методы:
// ⦁ AddStudent(name string, age int)
// ⦁ AverageAge() float64

// package main

// import "fmt"

// type Student struct {
// 	Name string
// 	Age  int
// }
// type Classroom struct {
// 	Students []Student
// }

// func (c *Classroom) AddStudent(name string, age int) {
// 	student := Student{Name: name, Age: age}
// 	c.Students = append(c.Students, student)
// }
// func (c *Classroom) AverageAge() float64 {
// 	if len(c.Students) == 0 {
// 		return 0
// 	}
// 	totalAge := 0
// 	for _, v := range c.Students {
// 		totalAge = totalAge + v.Age
// 	}
// 	return float64(totalAge) / float64(len(c.Students))
// }
// func main() {
// 	classroom := Classroom{}
// 	classroom.AddStudent("Shodmon", 23)
// 	classroom.AddStudent("Abdussabur", 25)
// 	classroom.AddStudent("Valy", 36)
// 	fmt.Printf("Average Age: %.2f\n", classroom.AverageAge())
// 	fmt.Println(classroom.Students)
// }

//------------------------------------------------------------------------------------

// ⦁ Склад товаров
// ⦁ Описание: Реализуйте структуру Product с полями name и quantity.
// Реализуйте структуру Warehouse с срезом продуктов и методы для
// добавления продукта и получения общего количества товаров на
// складе.
// ⦁ Методы:
// ⦁ AddProduct(name string, quantity int)
// ⦁ TotalQuantity() int

// package main

// import "fmt"

// type Product struct {
// 	name     string
// 	quantity int
// }
// type Warehouse struct {
// 	Products []Product
// }

// func (w *Warehouse) AddProduct(product Product) {
// 	w.Products = append(w.Products, product)
// }
// func (w *Warehouse) TotalQuantity() int {
// 	if len(w.Products) == 0 {
// 		return 0
// 	}
// 	totalquantity := 0
// 	for _, v := range w.Products {
// 		totalquantity = totalquantity + v.quantity
// 	}
// 	return totalquantity
// }
// func main() {
// 	warehouse := Warehouse{}
// 	A := Product{name: "apple", quantity: 200}
// 	B := Product{name: "banana", quantity: 300}
// 	C := Product{name: "coconut", quantity: 150}
// 	warehouse.AddProduct(A)
// 	warehouse.AddProduct(B)
// 	warehouse.AddProduct(C)
// 	warehouse.TotalQuantity()
// 	fmt.Println(warehouse.Products)
// 	fmt.Printf("TotalQuantity: %.d\n", warehouse.TotalQuantity())
// }

//------------------------------------------------------------------------------------

// Заметки и метки
// ⦁ Описание: Реализуйте структуру Note с полем content и tags.
// Реализуйте структуру Notebook с срезом заметок и методы для
// добавления заметки и получения всех заметок с указанной меткой.
// ⦁ Методы:
// ⦁ AddNote(content string, tags []string)
// ⦁ NotesWithTag(tag string) []Note

package main

import (
	"fmt"
)

type Note struct {
	content string
	tags    []string
}

type Notebook struct {
	notes []Note
}

func (n *Notebook) AddNote(content string, tags []string) {
	n.notes = append(n.notes, Note{content: content, tags: tags})
}

func (n *Notebook) NotesWithTag(tag string) []Note {
	var result []Note
	for _, note := range n.notes {
		for _, t := range note.tags {
			if t == tag {
				result = append(result, note)
				break
			}
		}
	}
	return result
}

func main() {
	nb := &Notebook{}

	nb.AddNote("Заметка 1", []string{"работа", "важное"})
	nb.AddNote("Заметка 2", []string{"личное", "записка"})
	nb.AddNote("Заметка 3", []string{"работа", "проект"})

	notes := nb.NotesWithTag("работа")

	for _, note := range notes {
		fmt.Println("Заметка: ", note.content)
	}
}
