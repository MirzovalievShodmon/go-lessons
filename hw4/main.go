package main

import "fmt"

func main() {
	PrintGreeting()
}
// Создайте функцию PrintGreeting, которая печатает "Доброе утро!", если dayType равен "утро";
//  "Добрый день!", если dayType равен "день"; и "Добрый вечер!", если dayType равен "вечер". 
//  Переменную  dayType вводить с консоли внутри функции.
func PrintGreeting(){
	var dayType string
	fmt.Scanln(&dayType)
	if dayType=="утро"{
		fmt.Println("Доброе утро!")
	}
	if dayType=="день"{
		fmt.Println("Добрый день!")
	}
	if dayType=="вечер"{
		fmt.Println("Добрый вечер!")
	}
}