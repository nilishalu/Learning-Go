package main

import (
	"fmt"
	"math"
)

func loopVariants() {
    for i := 0; i < 10; i++ {
		fmt.Printf("%v Hello\n", i)
	}
	//Init and Post statements can be skipped
	sum := 1
	for ; sum < 100; {
		sum += sum
	}
	fmt.Println(sum)
	//For is used as Go's while as well

	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	//If the condition is also skipped, then it becomes a infinite loop
}

func squareRoot(num float64) string {
     if (num < 0) {
        return squareRoot(-num) + "i"
	 }
	 return fmt.Sprint(math.Sqrt(num))
}

func power(num, p, lim float64) float64 {
	if res := math.Pow(num, p); res < lim {
       return res
	}
	return lim
}

func conditionalVariants() {
    fmt.Println(squareRoot(8), squareRoot(-7))
	fmt.Println(power(2, 4, 50), power(3, 3, 20))
}

func main() {
	loopVariants()
    conditionalVariants()
}