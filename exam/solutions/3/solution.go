package main

import "fmt"

func main() {
	a := 0
	if a != 1 || a != 2 {
		a++
	}

	fmt.Printf("a = %s\n", a)
}
