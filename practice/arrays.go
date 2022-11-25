package main

import (
	"fmt"
)

func sliceAsReference() {
	var names = []string {
		"John",
		"Anshu",
		"Rahul",
		"Pragya",
		"Prajakta",
		"Shiva",
		"Subrahmanya",
		"Vishal",
	}

	a := names[0:2]
	b := names[1:]
	fmt.Println(a, b, names)

	b[0] = "NoName"

	fmt.Println(a,b,names)
}

func main() {
	var a = [4]int{4, 5,23, 32}
	fmt.Println(a)
	var s[]int = a[0:2]
	last := []int(a[3:])
	fmt.Println(s, last)
	sliceAsReference()
}