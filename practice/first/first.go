package main

import "fmt"

func main() {
	fmt.Println("Hello all")
	fmt.Println(2 + 2)

	var firstName = "Shalini"
	var secondName = "C E"

	fmt.Println(firstName + secondName)

	var firstNames = [...] string {"Shalini", "Shiva", "Shilpa", "Sagari", "Meghana"}
	var lastNames = [...] string {"C E", "Dasari", "N S", "K S", "G S"}

	fmt.Println(firstNames)
	fmt.Println(lastNames)
	
    num := 57

	fmt.Printf("%v %d %x %o %b %T %T\n", num, num, num, num, num, num, &num)

	char := "H"
	str := "Hello"

	fmt.Printf("%v %s %T\n", char, char, char)
	fmt.Printf("%v %s %T\n", str, str, str)
}