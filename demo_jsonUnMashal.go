package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID    int
	Name  string
	Tel   string
	Email string
}

func main() {
	e := employee{}
	err := json.Unmarshal([]byte(`{"ID":101,"Name":"John","Tel":"123456789","Email":"John@gg.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e.Name)
}
