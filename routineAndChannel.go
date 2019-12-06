package main

import "fmt"

func main() {
	s := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("sender hello", i)
			s <- fmt.Sprintf("receiver hello %d", i)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-s
		fmt.Println(val)
	}
}

/*define length of channel
//and let it operate dead-lock
package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
  	c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
}
*/
