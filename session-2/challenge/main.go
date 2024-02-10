package main

import (
	"fmt"
	"strconv"
)

/*
Author: M. Syauqi Frizman
The pseudocode
 1. initialize kalimat
 2. initialize resMap
 3. loop until the end of kalimat
 4. when start looping
    4.1 assign value from the key to the stringAsInt
    4.1.1 define the key for resMap, resMap[string(kalimat[i])]
    4.1.2 convert string to int using strconv.Atoi(), it will return the int and boolean (error purpose)
    4.2 if the stringAsInt more than 0 (means the key has been calculated)
    4.2.1 increment the value of the key, use strconv.Itoa() to convert int to string
    4.2.2 then assign the value to the key
    4.3 otherwise, assign the value of key to 1 (because this is the first time the key is counted)
    4.4 print the key
 5. print all keys and values in the resMap
*/
func main() {

	// // un-comment if the system want to read input from user
	// fmt.Print("Masukkan kalimat: ")
	// // because in the package "fmt" there is no method to read full sentences (multiple words), we need to use bufio and os
	// // similar: new Scanner(System.in) in java
	// scanner := bufio.NewReader(os.Stdin)
	// kalimat, _ := scanner.ReadString('\n')
	// // to remove new line
	// kalimat = strings.TrimSpace(kalimat)

	var kalimat string = "selamat malam"

	var resMap map[string]string
	resMap = map[string]string{}

	for i := 0; i < len(kalimat); i++ {
		// use string() to convert (byte?/rune? a.k.a character) into string
		stringAsInt, _ := strconv.Atoi(resMap[string(kalimat[i])])
		if stringAsInt > 0 {
			resMap[string(kalimat[i])] = strconv.Itoa(stringAsInt + 1)
		} else {
			resMap[string(kalimat[i])] = strconv.Itoa(1)
		}
		fmt.Printf("%c\n", kalimat[i])
	}

	fmt.Print("map[")
	for key := range resMap {
		fmt.Printf("%s:%s ", key, resMap[key])
	}
	fmt.Print("]")
}

/*
Note: confused
every time we print at "fmt.Printf("%s:%s ", key, resMap[key])",
the order of the key and resMap[key] isn't fix (changes every time)

idk why??
*/
