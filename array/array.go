package main

import "fmt"

var productName [4]string
var price [4]float32

func main() {
	productName[0] = "Milk"
	price[0] = 22.3
	fmt.Println(productName)
	fmt.Println(price)
}
