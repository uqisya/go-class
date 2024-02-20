package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	id      int
	name    string
	address string
	job     string
	message string
}

func main() {
	var person = []Person{
		{
			id:      0,
			name:    "Syauqi",
			address: "Jl. Utama",
			job:     "Backend",
			message: "lorem dummy msg syauqi",
		},
		{
			id:      1,
			name:    "Uqi",
			address: "Jl. Selatan",
			job:     "Mobile Engineer",
			message: "lorem dummy msg uqi",
		},
		{
			id:      2,
			name:    "Gilang",
			address: "Jl. Timur",
			job:     "Data Engineer",
			message: "lorem dummy msg gilang",
		},
		{
			id:      3,
			name:    "Fitri",
			address: "Jl. Selatan",
			job:     "Fullstack",
			message: "lorem dummy msg fitri",
		},
	}

	// validation the cli input must have two argument,
	// main.go as first argument, <name> as second argument
	if len(os.Args) < 2 {
		// exit the program immediately
		fmt.Println("try add one argument, for ex. go run main.go <name>")
		return
	}

	// get second argument (index 1, because index 0 is main.go)
	inputData := os.Args[1] // for ex. go run main.go <syauqi>, inputData = 'syauqi'

	for _, value := range person {
		if inputData == value.name || inputData == strconv.Itoa(value.id) {
			dataPerson := Person{
				id:      value.id,
				name:    value.name,
				address: value.address,
				job:     value.job,
				message: value.message,
			}
			showAttendanceList(&dataPerson)
			return
		}
	}

	// if the name or absen doesn't exist
	fmt.Printf("Data dengan nama/absen <%s> tidak tersedia", inputData)
}

func showAttendanceList(p *Person) {
	fmt.Println("ID : ", p.id)
	fmt.Println("nama : ", p.name)
	fmt.Println("alamat : ", p.address)
	fmt.Println("pekerjaan : ", p.job)
	fmt.Println("alasan : ", p.message)
}
