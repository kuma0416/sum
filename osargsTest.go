package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(len(os.Args))
	for idx, args := range os.Args {
		fmt.Println("参数"+strconv.Itoa(idx)+":", args)
	}
}
