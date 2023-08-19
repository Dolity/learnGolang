package main

import "fmt"

func zerovalue(ivalue int) {
	ivalue = 0
}

func zeroPointer(iPointer *int) {
	*iPointer = 0
}

func main() {
	i := 1
	fmt.Println("i =", i)

	zerovalue(i)
	fmt.Println("i from zerovalue =", i)

	zeroPointer(&i)
	fmt.Println("i value from zeroPointer =", i)
	fmt.Println("i address from zeroPointer =", &i)
}
