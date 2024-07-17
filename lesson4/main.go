package main

import "fmt"

func main() {
num := 15
if num > 9 && num < 100 {
	fmt.Println("Двузначное")
}
n :=6
if n%2 ==1 {
	fmt.Println( "число нечетное")
} else {
	fmt.Println("число четное")
}
m :=78
if m>=0 && m <10  {
	fmt.Println( "число однозначное")
} else if m>=10 && m <100 {
	fmt.Println("число двузначное")
}else if m>=100 && m <1000 {
	fmt.Println("число трехзначное")
	}else  {
		fmt.Println("слишком большое")
}
a :=10 
b :=20
if a >=b {
	fmt.Println(a)
} else {
	fmt.Println(b)
}
 
}