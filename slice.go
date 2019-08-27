package main

import "fmt"

func main(){
	s := make([]string, 0, 5)
	fmt.Println(cap(s))
}