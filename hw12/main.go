// ⦁ Создание и вывод map
// Описание: Создайте map для хранения строк и их длин, добавьте несколько
// элементов и выведите содержимое.

// package main

// import "fmt"

// func main() {
// 	StringLengths := make(map[string]int)
// 	StringLengths["Shodmon"] = len("Shodmon")
// 	StringLengths["Valy"] = len("Valy")
// 	StringLengths["Abdussabur"] = len("Abdussabur")
// 	for key, value := range StringLengths {
// 		fmt.Printf("Строка: %s, Длина: %d\n", key, value)
// 	}
// }

//------------------------------------------------------------------------------------

// ⦁ Проверка наличия ключа
// Описание: Создайте map с несколькими элементами и напишите функцию,
// которая проверяет наличие заданного ключа.

// 1 СПОСОБ
// package main

// import "fmt"

// func Check(m map[string]int, key string) {
// 	if value, exists := m[key]; exists {
// 		fmt.Printf("Ключ '%s' существует в карте и связан с значением %d\n", key, value)
// 	} else {
// 		fmt.Printf("Ключ '%s'  не найден в карте\n", key)
// 	}
// }
// func main() {
// 	Map := make(map[string]int)
// 	Map["Shodmon"] = 23
// 	Map["Valy"] = 31
// 	Map["Abdusabur"] = 36
// 	Check(Map, "Shodmon")
// 	Check(Map, "Valy")
// 	Check(Map, "Said")
// }

// 2 СПОСОБ
// package main

// import "fmt"

// func Check(m map[string]int, key string) bool {
// 	_, exists := m[key]
// 	return exists
// }
// func main() {
// 	Map := make(map[string]int)
// 	Map["Shodmon"] = 23
// 	Map["Valy"] = 31
// 	Map["Abdussabur"] = 36
// 	Slice := []string{"Shodmon", "Valy", "Said"}
// 	for _, key := range Slice {
// 		if Check(Map, key) {
// 			fmt.Printf("Ключ '%s' существует в карте\n", key)
// 		} else {
// 			fmt.Printf("Ключ '%s'не найден в карте\n", key)
// 		}
// 	}
// }

//------------------------------------------------------------------------------------

// ⦁ Обновление значения по ключу
// Описание: Напишите функцию для обновления значения в map по заданному ключу.

// package main

// import "fmt"

// func Update(m map[string]int, key string, newValue int) {
// 	if _, exists := m[key]; exists {
// 		m[key] = newValue
// 	} else {
// 		fmt.Println("Ключ не найден в карте")
// 	}
// }
// func main() {
// 	Map := make(map[string]int)
// 	Map["Shodmon"] = 23
// 	Update(Map, "Shodmon", 25)
// 	Map["Valy"] = 31
// 	Update(Map, "Valy", 33)
// 	Map["Abdussabur"] = 36
// 	Update(Map, "Said", 33)
// 	fmt.Println(Map)
// }

//-----------------------------------------------------------------------------------

// ⦁ Удаление элемента из map
// Описание: Создайте функцию, которая удаляет элемент из map по заданному
// ключу.

// package main

// import "fmt"

// func Delete(m map[string]int, key string) {
// 	if _, exists := m[key]; exists {
// 		delete(m, key)
// 		fmt.Printf("Элемент с ключом '%s' удален\n", key)
// 	} else {
// 		fmt.Printf("Ключ '%s' не найден в карте\n", key)
// 	}
// }
// func main() {
// 	Map := make(map[string]int)
// 	Map["Shodmon"] = 23
// 	Map["Valy"] = 31
// 	Map["Abdussabur"] = 36
// 	Delete(Map, "Shodmon")
// 	fmt.Println(Map)
// }

//-----------------------------------------------------------------------------------

// ⦁ Подсчет частоты слов
// Описание: Напишите функцию, которая подсчитывает частоту слов в строке и
// возвращает map с результатами.

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // CountWordFrequency подсчитывает частоту каждого слова в строке
// func CountWordFrequency(text string) map[string]int {
// 	// Создаем карту для хранения частоты слов
// 	wordCount := make(map[string]int)

// 	// Разбиваем строку на слова, используя пробелы в качестве разделителей
// 	words := strings.Fields(text)

// 	// Подсчитываем частоту каждого слова
// 	for _, word := range words {
// 		wordCount[word]++
// 	}

// 	return wordCount
// }

// func main() {
// 	// Пример строки для подсчета частоты слов
// 	text := "hello world hello go world go hello"

// 	// Получаем частоту слов
// 	frequencies := CountWordFrequency(text)

// 	// Выводим частоту слов
// 	for word, count := range frequencies {
// 		fmt.Printf("Слово '%s' встречается %d раз\n", word, count)
// 	}
// }

//-----------------------------------------------------------------------------------

// ⦁ Сортировка ключей в map
// Описание: Напишите функцию, которая сортирует ключи в map и возвращает отсортированные ключи.

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// // Функция для получения отсортированных ключей карты
// func SortedKeys(m map[string]int) []string {
// 	// Создаем срез для хранения ключей
// 	keys := make([]string, 0, len(m))

// 	// Извлекаем ключи из карты
// 	for key := range m {
// 		keys = append(keys, key)
// 	}

// 	// Сортируем ключи
// 	sort.Strings(keys)

// 	return keys
// }

// func main() {
// 	// Создаем карту с начальными значениями
// 	myMap := map[string]int{
// 		"Shodmon":    23,
// 		"Valy":       31,
// 		"Abdussabur": 36,
// 	}

// 	// Получаем отсортированные ключи
// 	sortedKeys := SortedKeys(myMap)

// 	// Выводим отсортированные ключи и их значения
// 	fmt.Println("Отсортированные ключи и их значения:")
// 	for _, key := range sortedKeys {
// 		fmt.Printf("Ключ: '%s', Значение: %d\n", key, myMap[key])
// 	}
// }

//-----------------------------------------------------------------------------------

// ⦁ Проверка пустого map
// Описание: Напишите функцию, которая проверяет, пустой ли map.

// package main

// import (
// 	"fmt"
// )

// // Функция для проверки, пустой ли map
// func isMapEmpty(m map[string]int) bool {
// 	// Проверяем длину map. Если длина равна 0, то map пустой.
// 	return len(m) == 0
// }

// func main() {
// 	// Примеры использования функции isMapEmpty

// 	// Пустой map
// 	emptyMap := make(map[string]int)

// 	// Map с элементами
// 	nonEmptyMap := map[string]int{
// 		"apple":  1,
// 		"banana": 2,
// 	}

// 	// Проверка пустоты map
// 	fmt.Println("emptyMap пустой?", isMapEmpty(emptyMap))
// 	fmt.Println("nonEmptyMap пустой?", isMapEmpty(nonEmptyMap))
// }

//-----------------------------------------------------------------------------------

// ⦁ Инвертирование ключей и значений
// Описание: Напишите функцию, которая инвертирует ключи и значения в map.

// package main

// import (
// 	"fmt"
// )

// // Функция для инвертирования ключей и значений в map
// func invertMap(m map[string]int) map[int]string {
// 	inverted := make(map[int]string, len(m)) // Создаем новый map для хранения инвертированных данных

// 	// Проходим по каждому элементу исходного map
// 	for key, value := range m {
// 		inverted[value] = key // Меняем местами ключи и значения
// 	}

// 	return inverted
// }

// func main() {
// 	// Исходный map
// 	originalMap := map[string]int{
// 		"apple":  1,
// 		"banana": 2,
// 		"cherry": 3,
// 	}

// 	// Инвертируем map
// 	invertedMap := invertMap(originalMap)

// 	// Выводим результат
// 	fmt.Println("Оригинальный map:", originalMap)
// 	fmt.Println("Инвертированный map:", invertedMap)
// }

//-----------------------------------------------------------------------------------

// ⦁ Проверка дубликатов в map
// Описание: Напишите функцию, которая проверяет, есть ли дубликаты значений в map.

//1 СПОСОБ

// package main

// import "fmt"

// func Repeat(m map[string]int) bool {
// 	r := 0
// 	for _, v1 := range m {
// 		for _, v2 := range m {
// 			if v1 == v2 {
// 				r++
// 			}
// 			if r == 2 {
// 				return true
// 			}
// 		}
// 	r = 0
// }
// return false
// }
// func main() {
// 	Map := map[string]int{
// 		"apple":  1,
// 		"banana": 1,
// 		"cherry": 3,
// 	}
// 	fmt.Println(Repeat(Map))
// }

//2 СПОСОБ

// package main

// import "fmt"

// // Функция для проверки наличия дубликатов значений в map
// func hasDuplicateValues(m map[string]int) bool {
//     seen := make(map[int]bool) // Карта для отслеживания встреченных значений

//     for _, value := range m {
//         if _, exists := seen[value]; exists {
//             // Если значение уже встречалось, возвращаем true
//             return true
//         }
//         // Добавляем значение в карту как встреченное
//         seen[value] = true
//     }

//     // Если дубликаты не найдены, возвращаем false
//     return false
// }

// func main() {
//     m := map[string]int{
//         "apple":  1,
//         "banana": 2,
//         "cherry": 3,
//     }

//     fmt.Println("Есть ли дубликаты значений?", hasDuplicateValues(m)) // Output: false

//     // Пример с дубликатами
//     mWithDuplicates := map[string]int{
//         "apple":  1,
//         "banana": 2,
//         "cherry": 1,
//     }

//     fmt.Println("Есть ли дубликаты значений?", hasDuplicateValues(mWithDuplicates)) // Output: true
// }

//-----------------------------------------------------------------------------------
