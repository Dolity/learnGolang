package main

import "fmt"

func main() {
	input := 5
	switch input {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("another number")
	}

}
