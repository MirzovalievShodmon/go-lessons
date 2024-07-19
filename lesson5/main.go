package main

import "fmt"

func main() {
	// var i int
	// for i := 0; i <= 9; i++ {
	// 	fmt.Println(i)
	// }
	// fmt.Println(i)

	// var num = 546
	// for num > 0 {
	// 	a := num % 10
	// 	fmt.Println(a)
	// 	num = num / 10
	// }
	// for n := 546; n > 0; n /= 10 {
	// 	a := n % 10
	// 	fmt.Println(a)
	// }
	//    var n int
	//    n=4
	//    или
	//    n:=4

	// for{
	// 	time.Sleep(1*time.Second)
	// 	fmt.Println("HELLO")
	// }
	// for i:=1; i<=100; i++ {
	// 	if i%3 == 0 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	// еще есть continiue
	// for i := 1; i < 10; i++ {
	// 	for j := 1; j < 10; j++ {
	// 		fmt.Printf("%d * %d = %d\n", i, j, i*j)
	// 	}
	// 	fmt.Println(i)
	// }
	// var n int = 10
	// var s float64 = 0
	// for i := 1; i <= n; i++ {
	// 	s = s + 1/float64(i)
	// }
	// fmt.Println(s)
	// var n int = 3
	// var p float64 = 1
	// for i := 1; i <= n; i+=1 {
	// 	p = p * float64(i+10)/10
	// 	fmt.Println(p)
	// }
	// var p float64 = 1
	// var a float64
	// fmt.Scan(&a)
	// for i := 0; i < 10; i ++ {
	// 	fmt.Scan(&a)
	// 	p = p * a
	// }
	// fmt.Println(p)
	var p float64 = 1
	var n int
	var a float64 
	var b float64 
	fmt.Scan(&n)
	for i := 1; i <= n; i ++ {
		fmt.Scan(&b)
		a = float64((n) - int(n))
		fmt.Println(a)
		p = p *a
	}
	fmt.Println(p)
}
