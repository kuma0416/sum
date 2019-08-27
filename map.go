package main

import "fmt"

func main(){
	h := map[string]int{"Sam": 188,"tom": 194,"yen": 160}

	h["gogo"] = 165

	fmt.Println(len(h))

	delete(h, "gogo")
	if key, exists := h["tom"]; exists {
		fmt.Println(key,exists)
	}
}