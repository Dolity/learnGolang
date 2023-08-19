package main

import "fmt"

func hello() {
	fmt.Println("Hello World")
}

func plus(value1 int, value2 int) int {
	return value1 + value2
}

func plus3value(value1, value2, value3 int) int {
	return value1 + value2 + value3
}

func main() {
	hello()
	result := plus(5, 4)
	fmt.Println("result =", result)

	result2 := plus(9, 8)
	fmt.Println("result =", result2)
}
