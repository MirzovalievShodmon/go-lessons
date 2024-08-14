package main

import (
	"fmt"
	"os"
)

func main() {
	// Открытие или создание файла с правами на чтение и запись
	file, err := os.OpenFile("example.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	fmt.Println("File opened successfully")
}
