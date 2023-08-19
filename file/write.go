package main

import "os"

func main() {
	data1 := []byte("hello\n golang")
	err := os.WriteFile("C:/WebDev/demoGolang/learnGolang/file/data.txt", data1, 0644)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("employeeName")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	data2 := []byte("Write\n from Golang")
	os.WriteFile("employee1", data2, 0644)
}
