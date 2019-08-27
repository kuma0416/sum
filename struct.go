package main

import "fmt"

type person struct {
	name string
	height int
}

func (p person) greeting() {
	fmt.Println("hi!")
}
func main(){
	p := person{"tony", 194} //給struct值
	p.greeting()
}