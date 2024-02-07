package main

import (
	"fmt"
	"math"
)

/*
Author: M. Syauqi Frizman
The pseudocode
 1. user insert a number
 2. loop until the number
 3. when start looping
    3.1 get square root or cube root of i
    3.2 if the value can be modulo by 1 and the result is 0 (it means the number is integer, then do the print)
    3.3 otherwise, the number isn't a perfect square or cube (don't print, just print the number of i)
*/
func main() {

	fmt.Print("Input n: ")
	var number int
	fmt.Scan(&number)

	for i := 0; i < number; i++ {
		// it feels too strict that the Sqrt method from math lib only receive float64, so I can't do something like math.Sqrt(i+1)
		sqrtNum := math.Sqrt(float64(i + 1))
		cbrtNum := math.Cbrt(float64(i + 1))

		var kuadrat bool = (math.Mod(sqrtNum, 1) == 0)
		if kuadrat { // kuadrat sempurna "Square"
			fmt.Print("Square")
		}

		var kubus bool = (math.Mod(cbrtNum, 1) == 0)
		if kubus { // kubus sempurna "Cube"
			fmt.Print("Cube")
		}

		if !kuadrat && !kubus {
			fmt.Println(i + 1)
		} else {
			fmt.Println()
		}
	}
}
