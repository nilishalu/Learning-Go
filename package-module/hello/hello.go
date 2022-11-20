package main

import (
	"fmt"
	"example/greetings"
)

func main() {
	message := greetings.Hello('Shalini')
	fmt.Println(message)
}