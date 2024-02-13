package main

import (
	"fmt"
	"session-3/helpers"
)

type Employee struct {
	name     string
	age      int
	position string
}

type Company struct {
	name     string
	location []string
	employee []Employee
}

func (e Employee) greet(message string) string {
	return fmt.Sprintf("msg: %s\nMy name is %s, and I am %d years old", message, e.name, e.age)
}

func main() {
	// var employee Employee
	// employee.name = "Syauqi"
	// employee.age = 21
	// employee.position = "Developer"

	// fmt.Printf("name: %s\nage: %d\nposition: %s\n", employee.name, employee.age, employee.position)

	// var employee2 = Employee{
	// 	name:     "uqisya",
	// 	age:      25,
	// 	position: "Backend",
	// }

	// fmt.Printf("name: %s\nage: %d\nposition: %s\n", employee2.name, employee2.age, employee2.position)

	// fmt.Println("Employee 1:", employee)
	// fmt.Printf("Employee 2: %+v\n", employee2)

	// var myCompany = Company{
	// 	name: "PT ABC",
	// }

	// myCompany.employee = []Employee{
	// 	{
	// 		name:     "syauqi",
	// 		age:      21,
	// 		position: "Backend",
	// 	},
	// 	{
	// 		name:     "uqisya",
	// 		age:      25,
	// 		position: "Frontend",
	// 	},
	// }

	// myCompany.location = []string{
	// 	"Jakarta",
	// 	"Bandung",
	// 	"Purwakarta",
	// }

	// for _, value := range myCompany.location {
	// 	fmt.Println(value)
	// }

	// secondFunc("syauqi", "jl. raya utama")

	// employee5 := Employee{
	// 	name: "uqi",
	// 	age:  25,
	// }

	// myCompany := Company{
	// 	name: "PT XYZ",
	// }
	// myCompany.employee = []Employee{
	// 	{
	// 		name:     "syauqi",
	// 		age:      21,
	// 		position: "Backend",
	// 	},
	// 	{
	// 		name:     "uqisya",
	// 		age:      25,
	// 		position: "Frontend",
	// 	},
	// }

	// fmt.Println(employee5.greet("hai"))

	// for _, e := range myCompany.employee {
	// 	fmt.Println(e.greet("hai!"))
	// }

	// msg, num := secondStringInt("syauqi", 25)
	// fmt.Printf("msg: %s\nnum: %d\n", msg, num)

	// numberList := []int{1, 2, 3, 4, 5}

	// fmt.Println(sumNumberList(numberList...))

	// var angka int = 12
	// value := reflect.ValueOf(angka)
	// typeD := reflect.TypeOf(angka)

	// fmt.Println(value)
	// fmt.Println(typeD)

	helpers.Greet()

	var person = helpers.Person{
		Name:    "syauqi",
		Age:     25,
		Address: "Jl. raya utama",
	}

	helpers.Person.InvokeGreet(person)
}

func secondFunc(name string, address string) {
	fmt.Println("Hello there, my name is ", name)
	fmt.Println("I live at ", address)
}

func secondStringInt(message string, number int) (string, int) {
	return message, number
}

func sumNumberList(numberList ...int) int {
	total := 0
	for _, value := range numberList {
		total += value
	}

	return total
}
