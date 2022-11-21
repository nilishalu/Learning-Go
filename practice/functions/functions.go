package main

import (
	"fmt"
	"math/rand"
)

func add(n1 int, n2 int) int {
	return n1 + n2
}

func multiply(n1, n2 int) int {
	return n1 * n2
}

func swapNames(name1, name2 string) (string, string) {
	return name2, name1
}

func splitValue(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
    fmt.Println("My favorite number is", rand.Intn(15))
	// fmt.Println(add(5, -2))
	fmt.Println(multiply(4, 8))
	fmt.Println(swapNames("Shalu", "Shiva"))
	fmt.Println(splitValue(17))
}