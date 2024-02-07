package main

import "fmt"

func main() {
	/*
		// MACAM-MACAM PENULISAN VARIABLE
		// with data type
		var title string = "GO Programming Language"
		var totalMember int = 18

		var title1 string
		var totalMember1 int
		title1 = "GO Programming Language"
		totalMember1 = 18

		fmt.Println("Course title is: ", title)
		fmt.Println("The total member is: ", totalMember)
		fmt.Println("Course title is: ", title1)
		fmt.Println("The total member is: ", totalMember1)

		// without data type
		title := "GO Programming Language" // short declaration technique
		totalMember := 18

		fmt.Printf("%T, %T\n", title, totalMember) // untuk print jenis tipe datanya
		fmt.Println("Course title is: ", title)
		fmt.Println("The total member is: ", totalMember)

		// multiple declaration variable
		var firstName, surName, lastName string = "Muhammad", "Syauqi", "Frizman"

		var age1, age2, age3 int
		age1, age2, age3 = 20, 21, 22
		// age1, age2, age3 := 20, 21, 22

		fmt.Println(firstName, surName, lastName)
		fmt.Println(age1, age2, age3)

		fmt.Printf("Hello, my name is %s and my age is %d years old\n", firstName, age2)
	*/

	// first := 15.125

	// firstString := "2543"

	// fmt.Println(first)

	// status, _ := strconv.Atoi(firstString)

	// fmt.Println(status)

	// angkaString := "0234"

	// angkaInt, _ := strconv.Atoi(angkaString)
	// var angkaMurni int = 0234

	// fmt.Println(angkaString)
	// fmt.Println(angkaInt)
	// fmt.Println(angkaMurni)

	/*
		// TIPE-TIPE DATA
		// 1. Basic Type (number, string, boolean)
		// INTEGER
		first := 22
		second := -23
		fmt.Printf("First data type: %T\n", first)
		fmt.Printf("Second data type: %T\n", second)

		var third uint8 = 98
		var fourth int8 = -9
		fmt.Printf("Third data type: %T\n", third)
		fmt.Printf("Fourth data type: %T\n", fourth)

		// FLOAT
		decimalNumber := 3.45
		fmt.Printf("Decimal number: %f\n", decimalNumber)
		fmt.Printf("Decimal number: %.3f\n", decimalNumber)

		// BOOLEAN
		condition := false
		fmt.Printf("are you a women? %t\n", condition)

		// STRING
		var message string = "Hello!"
		fmt.Printf("%s\n", message)

		// 2. Aggregate Type (array, struct)

		// 3. Reference Type (slices, map, pointer, function, channel)

		// 4. Interface Type (interface)
	*/

	/*
		// CONSTANT & OPERATOR
		// Constant -> jenis variable yang nilainya tidak dapat berubah a.k.a final

		// IOTA -> representasi urutan pengingatkan variable (increment) konstan

		// ATOI -> method pada strconv untuk konversi string ke integer

		// Operator
		// 1. Operator aritmatika
		var value = (2 + 2) * 3
		fmt.Println(value)

		// 2. Operator relational/perbandingan
		var firstCondition bool = 2 < 3
		var secondCondition bool = "women" == "Women"
		var thirdCondition bool = 10 != 2.3
		var fourthCondition bool = 11 <= 11

		fmt.Println("first: ", firstCondition)
		fmt.Println("first: ", secondCondition)
		fmt.Println("first: ", thirdCondition)
		fmt.Println("first: ", fourthCondition)

		// 3. Operator logical/logika
		right := true
		wrong := false

		wrongAndRight := right && wrong
		fmt.Println("wrong && right: ", wrongAndRight)

		wrongOrRight := right || wrong
		fmt.Println("Wrong || right: ", wrongOrRight)

		wrongNegation := !wrong
		fmt.Println("!wrong: ", wrongNegation)
	*/

	/*
		// CONDITION
		// if-else
		currentYear := 2023
		age := currentYear - 2018

		if age < 17 {
			fmt.Println("Anda belum bisa membuat KTP")
		} else {
			fmt.Println("Anda sudah bisa membuat KTP")
		}

		// switch case
		score := 88
		switch {
		case score == 100:
			fmt.Println("Perfect!")
		case (score >= 81) && (score <= 99):
			fmt.Println("Good Job")
		case (score >= 60) && (score <= 80):
			fmt.Println("Not bad")
		default:
			fmt.Println("Please study harder!")
		}
	*/

	/*
		// LOOPING
		// 1. For Loop
		for i := 0; i < 3; i++ {
			fmt.Println("Angka: ", i)
		}

		var i int = 0
		for i < 3 {
			fmt.Println("Angka: ", i)
			i++
		}

		i = 0
		for {
			fmt.Println("Angka: ", i)
			i++
			if i == 3 {
				break
			}
		}

		for i := 0; i < 5; i++ {
			for j := i; j < 5; j++ {
				fmt.Println(j, " ")
			}
			fmt.Println()
		}

		// For range loop
		var data []string
		data = []string{"aaa", "bbb"}
		for increment, content := range data {
			fmt.Println("No ", increment+1, ":", content)
		}

		datas := []string{"a", "b", "c"}
		for _, content := range datas {
			fmt.Println(content)
		}
	*/

	// switch case
	score := 88
	switch {
	case score == 100:
		fmt.Println("Perfect!")
	case (score > 60) && (score < 100):
		fmt.Println("Not bad")
	default:
		fmt.Println("Please study harder!")
	}

	// For range loop
	// var data []string
	// data = []string{"aaa", "bbb"}
	// for _, content := range data {
	// 	fmt.Println(content)
	// }

}
