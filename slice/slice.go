package main

import "fmt"

func main() {
	var courseName []string
	courseName = []string{"Golang", "Java", "Python"}
	fmt.Println(courseName)
	courseName = append(courseName, "C++", "Dart")
	fmt.Println(courseName)

	// Fliter
	courseweb := courseName[1:3]
	fmt.Println(courseweb)
	courseweb3 := courseName[:3]
	fmt.Println(courseweb3)

}
