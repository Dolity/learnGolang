package main

import "fmt"

func plus(value1, value2 float64) {
	result := value1 + value2
	fmt.Println(result)
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}
}

func deferloop() {
	for j := 0; j < 10; j++ {
		defer fmt.Println("j = ", j)
	}
}

func main() {
	// fmt.Println("Hello World")
	// defer fmt.Println("End of Program")
	// defer plus(20, 10)
	// plus(10, 2)
	// plus(5, 2)

	loop()
	deferloop()

}
