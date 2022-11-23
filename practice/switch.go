package main

import (
	"fmt"
	"runtime"
	"time"
)

func workingOS() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s \n", os)
	}
}

func nearestSunday() {
	today := time.Now().Weekday()
	switch time.Wednesday {
	case today + 0:
		fmt.Println("It's today, yey!!")
	case today + 1:
		fmt.Println("It's tomorrow, yey!")
	case today + 2:
		fmt.Println("It's in 2 days!")
	default:
		fmt.Println("It's too far away :(")
	}
}

func greetMe() {
	currTime := time.Now()
	switch {
	case currTime.Hour() < 12:
		fmt.Println("Good morning!")
	case currTime.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

func deferExecution() {
	defer fmt.Println("It's me")
	fmt.Println("Heyyy")
}

func main() {
	workingOS()
	nearestSunday()
	greetMe()
	deferExecution()
}