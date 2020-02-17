//example 1
/*package main

import "fmt"

func main(){
	c := make(chan int)

	go func() {
		for i:=0; i<10; i++{
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}*/
//example 2
package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
