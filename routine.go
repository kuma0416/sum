/*
package main

import (
	"fmt"
	"sync"
	"time"
)
func foo() {
	for i:=1;i<11;i++{
		fmt.Println("foo:",i)
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go foo()
	time.Sleep(5 * time.Millisecond)
	wg.Wait()
}
*/ //example 1
package main

import (
	"fmt"
	"sync"
	"time"
)

func withdraw() {
	mu.Lock()
	balance := money
	time.Sleep(3000 * time.Millisecond)
	balance -= 1000
	money = balance
	mu.Unlock()
	fmt.Println("after withdrawing $1000, balance: ", money)
	wg.Done()
}

var wg sync.WaitGroup
var money int = 1500
var mu sync.Mutex

func main() {
	fmt.Println("we have $1500 balance")
	wg.Add(2)
	go withdraw()
	go withdraw()
	wg.Wait()
}
