package main

import (
	"fmt"
)

type Student struct {
	name string
	rollno int
	id string
}

func pointers() {
   n1, n2 := 23, 44

   p := &n1
   fmt.Println(*p)
   *p = 28
   fmt.Println(*p)

   *p = (*p) * 2
   fmt.Println(n1)
   fmt.Println(&n2)
}

func main() {
   pointers()
   s1 := Student{"Shalini", 2010580, "201IT256"}
   p1 := &s1
   p1.name = "Shalu"
   fmt.Println(s1)
   s2 := Student{name: "Megha"}
   fmt.Println(s2)
}