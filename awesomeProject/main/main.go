package main

import "fmt"

func main() {
	for i := 1060; i < 3183; i++ {
		fmt.Printf("%d - %b - %#x, %q \n", i, i, i, i)
	}
}
