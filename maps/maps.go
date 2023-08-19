package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("Product =", product)

	//add
	product["milk"] = 22.3
	product["water"] = 12.3
	product["coke"] = 25
	fmt.Println("Product =", product)

	//delete
	delete(product, "water")
	fmt.Println("Product =", product)

	//update
	product["milk"] = 20
	fmt.Println("Product =", product)

	// access map
	value1 := product["milk"]
	fmt.Println("Value =", value1)

	courseName := map[string]string{"101": "Golang", "102": "C#"}
	fmt.Println("Course =", courseName)
}
