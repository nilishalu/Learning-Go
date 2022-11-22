package main

import (
	"fmt"
	"math/rand"
	"math"
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

func squareRoot(num int) (rootFloat float64, rootInt int) {
    rootFloat = math.Sqrt(float64(num))
	rootInt = int(rootFloat)
	return 
}

func main() {
	var a, b int = 5, 6
	msg := "I can learn Go"
    const truth = true
	const pi = math.Pi

    fmt.Println("My favorite number is", rand.Intn(15))
	// fmt.Println(add(5, -2))
	fmt.Println(multiply(4, 8))
	fmt.Println(swapNames("Shalu", "Shiva"))
	fmt.Println(splitValue(17))
	fmt.Printf("%d %d \n", a, b)
	fmt.Println(msg)
	fmt.Println(squareRoot(6))
	fmt.Printf("%v %v", pi, truth)
}