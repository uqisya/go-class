package main

import (
	"fmt"
)

func main() {
	// balances := [3][5]int{
	// 	{5, 6},
	// 	{7, 2, 3, 4},
	// }

	// fmt.Println(balances)

	// var animals = make([]string, 3)
	// animals[0] = "kucing"
	// animals[1] = "anjing"

	// animals = append(animals[:1], "harimau", "gajah", "ayam")
	// fmt.Println(animals)

	// var newFruits = []string{
	// 	"apple",
	// 	"manggo",
	// 	"guava",
	// }

	// var newFruits1 = newFruits[:2]

	// fmt.Println("Fruits capacity: ", cap(newFruits1))
	// fmt.Println("Fruits length: ", len(newFruits1))

	// var people = []map[string]string{
	// 	{"name": "Jisoo", "age": "27"},
	// 	{"name": "test", "age": "30"},
	// }

	// for _, person := range people {
	// 	fmt.Printf("name: %s, age: %s\n", person["name"], person["age"])
	// }

	fmt.Println("----- mainFunc -----")
	var firstNumber int = 4
	fmt.Println("firstNumber&: ", &firstNumber) // address firstNumber

	secondFunc(&firstNumber)

	fmt.Println("----- mainFunc -----")
	fmt.Print("\n[check value dari firstNumber]: ", firstNumber)
}

func secondFunc(secondNumber *int) {
	fmt.Println("\n----- secondFunc -----")
	fmt.Println("secondNumber: ", secondNumber)   // merujuk ke address firstNumber
	fmt.Println("secondNumber&: ", &secondNumber) // address secondNumber
	fmt.Println("value: ", *secondNumber)         // value dari firstNumber

	thirdFunc(&secondNumber)
}

func thirdFunc(thirdNumber **int) {
	fmt.Println("\n----- thirdFunc -----")
	**thirdNumber = 21
	fmt.Println("----- update value -----")
	fmt.Println("thirdNumber: ", thirdNumber)   // merujuk ke address secondNumber
	fmt.Println("thirdNumber&: ", &thirdNumber) // address thirdNumber
	fmt.Println("value: ", **thirdNumber)       // value dari firstNumber
}
