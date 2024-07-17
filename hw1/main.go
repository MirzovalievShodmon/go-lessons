package main

import "fmt"

func main() {
	var a int = 1
	fmt.Println(a)

	var b int8=2
	fmt.Println(b)

	var c int16=2593
	fmt.Println(c)

	var d int32=123456789
	fmt.Println(d)

	var e int64=9080706050403020100
	fmt.Println(e)

	var f float32=26.256
	fmt.Println(f)	

	var h float64=96526.965262
	fmt.Println(h)	
	
	var m,P float32
	m=5
	P=4*m
	fmt.Println(P)	

	var a2,S float32
	a2=10
	S=a2*a2
	fmt.Println(S)	

	var a3,b3,P3,S3 float32
	a3=12
	b3=20
	P3=2*(a3+b3)
	S3=a3*b3
	fmt.Println(P3, S3)	

	const P4=3.14
	var D4,L4 float32
	D4=100
    L4=P4*D4
	fmt.Println(L4)		

	var a5,v5,s5 float32
	a5=1.5
	v5=a5*a5*a5
	s5=6*a5*a5
	fmt.Println(v5,s5)

	var a6,b6,c6,v6,s6 float32
	a6=1
	b6=2
	c6=3
	v6=a6*b6*c6
	s6=2*(a6*b6+b6*c6+c6*a6)
	fmt.Println(v6,s6)


}