package main

import (
	"fmt"
	"strings"
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

func sliceUsingMake() {
	a := make([]int, 5)
	fmt.Println(len(a), cap(a), a)
	b := make([]int, 1, 7)
	fmt.Println(len(b), cap(b), b)
	c := b[3:6]
    fmt.Println(len(c), cap(c), c)
}

func sliceMore() {
	arr := [...]int{3,445, 452, 23,234, 32, 2,34,64,12,6,31}
	a := arr[2:8]
	fmt.Println(len(a), cap(a))
}

func sliceOfSlice() {
	board := [][]string {
		[]string{"_", "_", "_"}, 
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[2][2] = "O"
	board[0][1] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}

func appendUsage() {
	var s []int
	fmt.Println(len(s), cap(s), s)
    s = append(s, 1)
	fmt.Println(len(s), cap(s), s)
	s = append(s, 5)
	fmt.Println(len(s), cap(s), s)
	s = append(s, 4, 9,23, 46)
	fmt.Println(len(s), cap(s), s)
}

func rangeMethod() {
	pow := [...]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	for i, v := range pow {
        fmt.Printf("2**%d = %d \n", i, v)
	}
}

func powRange() {
	pow := make([]int, 11)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func main() {
	// var a = [4]int{4, 5,23, 32}
	// fmt.Println(a)
	// var s[]int = a[0:2]
	// last := []int(a[3:])
	// fmt.Println(s, last)
	//sliceAsReference()
	// sliceMore()
	// sliceUsingMake()
	// sliceOfSlice()
	// appendUsage()
	rangeMethod()
	powRange()
}