package main

import "fmt"

func main(){
	var p *int
	a := 15

	p = &a

	fmt.Println(p)
	fmt.Println(*p)

	a = 10

	fmt.Println(p)
	fmt.Println(*p)
}