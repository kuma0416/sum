package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	//map架構為:count[字元] = 個數 (字元型別為string, 個數型別為int)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	// n -> 個數, line -> 字元
	for line, n := range counts {
		if n > 1 { //如果該個數>1就代表有重複的字元
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
