//// 1_INTRODUCE

// package main

// import "fmt"

// type Speaker interface {
// 	Speak() string
// }

// type Person struct {
// 	Name string
// }

// func (p Person) Speak() string {
// 	return "Hello, my name is " + p.Name
// }

// func Greet(s Speaker) {
// 	fmt.Println(s.Speak())
// }

// func introduce() {
// 	p := Person{Name: "Alice"}
// 	Greet(p) // Output: Hello, my name is Alice
// }

// func main() {
// 	introduce()
// }

//-----------------------------------------------------------------------------------

//2_EMBEDDING

// package main

// import "fmt"

// // Starter Интерфейс для запуска
// type Starter interface {
// 	Start() string
// }

// // Mover Интерфейс для движения
// type Mover interface {
// 	Move() string
// }

// // Stopper Интерфейс для остановки
// type Stopper interface {
// 	Stop() string
// }

// // Vehicle Интерфейс для транспортного средства
// type Vehicle interface {
// 	Starter
// 	Mover
// 	Stopper
// }

// // Реализация автомобиля
// type Car struct {
// 	Brand string
// }

// func (c Car) Start() string {
// 	return c.Brand + " started"
// }

// func (c Car) Move() string {
// 	return c.Brand + " is moving"
// }

// func (c Car) Stop() string {
// 	return c.Brand + " stopped"
// }

// // Bicycle Реализация велосипеда
// type Bicycle struct {
// 	Brand string
// }

// func (b Bicycle) Start() string {
// 	return b.Brand + " started"
// }

// func (b Bicycle) Move() string {
// 	return b.Brand + " is moving"
// }

// func (b Bicycle) Stop() string {
// 	return b.Brand + " stopped"
// }

// // OperateVehicle Функция для работы с любым транспортным средством
// func OperateVehicle(v Vehicle) {
// 	fmt.Println(v.Start())
// 	fmt.Println(v.Move())
// 	fmt.Println(v.Stop())
// }

// func embedding() {
// 	car := Car{Brand: "Toyota"}
// 	bike := Bicycle{Brand: "Giant"}

// 	OperateVehicle(car)
// 	OperateVehicle(bike)
// }
// func main() {
// 	embedding()
// }

//-----------------------------------------------------------------------------------

//3_DUCK

// package main

// import (
// 	"fmt"
// )

// // Animal
// // Интерфейс Animal с методом Speak
// type Animal interface {
// 	Speak() string
// }

// // Dog
// // Тип Dog реализует метод Speak
// type Dog struct{}

// func (d Dog) Speak() string {
// 	return "Woof!"
// }

// // Cat
// // Тип Cat реализует метод Speak
// type Cat struct{}

// func (c Cat) Speak() string {
// 	return "Meow!"
// }

// // Cow Тип Cow реализует метод Speak
// type Cow struct{}

// func (c Cow) Speak() string {
// 	return "Moo!"
// }

// // Describe
// // Функция для работы с любым Animal
// func Describe(a Animal) {
// 	fmt.Println(a.Speak())
// }
// func main() {
// 	dog := Dog{}
// 	cat := Cat{}
// 	cow := Cow{}

// 	Describe(dog) // Output: Woof!
// 	Describe(cat) // Output: Meow!
// 	Describe(cow) // Output: Moo!// }

//-----------------------------------------------------------------------------------

// 4_INTERFACE_VALUE

// package main

// import (
// 	"fmt"
// )

// // Определяем интерфейс Animal
// type Animal interface {
// 	Speak() string
// }

// // Определяем структуру Dog
// type Dog struct{}

// // Реализуем метод Speak для Dog
// func (d Dog) Speak() string {
// 	return "Woof!"
// }

// // Определяем структуру Cat
// type Cat struct{}

// // Реализуем метод Speak для Cat
// func (c Cat) Speak() string {
// 	return "Meow!"
// }

// func interfaceValue() {
// 	var a Animal

// 	// Присваиваем значение типа Dog переменной интерфейса Animal
// 	a = Dog{}
// 	fmt.Printf("Value: %v, Type: %T\n", a, a) // Output: Value: {}, Type: main.Dog
// 	fmt.Println(a.Speak())                    // Output: Woof!

// 	// Присваиваем значение типа Cat переменной интерфейса Animal
// 	a = Cat{}
// 	fmt.Printf("Value: %v, Type: %T\n", a, a) // Output: Value: {}, Type: main.Cat
// 	fmt.Println(a.Speak())                    // Output: Meow!
// }

// func main() {
// 	interfaceValue()
// }

//-----------------------------------------------------------------------------------

// 5_NIL_INTERFACES

