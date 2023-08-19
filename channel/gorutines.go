package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	go f("Hello Golang")
	go f("msg2")
	time.Sleep(5 * time.Second)
}
