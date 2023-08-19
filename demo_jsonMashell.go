package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID    int
	Name  string
	Tel   string
	Email string
}

func main() {

	data, _ := json.Marshal(&employee{101, "John", "123456789", "John@gg.com"})
	fmt.Println(string(data))
}
