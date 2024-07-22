// ⦁ Учет накопительных счетов с ежемесячным пополнением
// ⦁ Начальный баланс накопительного счета равен 100000 рублей.
// ⦁ Реализуйте функции для пополнения счета каждый месяц на
// фиксированную сумму.
// ⦁ Выводите баланс после каждого пополнения.
// ⦁ Если баланс становится больше 500000 рублей, выведите сообщение
// "Достигнут лимит накоплений".
package main

import "fmt"

func main() {
	var nbalans int = 100000
	var nfs int = 100000
	var nbptr *int
	nbptr = &nbalans
	ps(nbptr, nfs)
}
func ps(nbptr *int, nfs int) {
	for *nbptr <= 500000 {
		*nbptr = *nbptr + nfs
		fmt.Println(*nbptr)
	}
	if *nbptr > 500000 {
		fmt.Println("Достигнут лимит накоплений")
	} else {
		fmt.Println(*nbptr)
	}
}