package greetings

import "fmt" 

func Hello(name string) string {
	message := fmt.sprintf("Hello, %v. Welcome to the Go World!", name)
	return message
}