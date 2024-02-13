package helpers

import "fmt"

func greet(p Person) {
	fmt.Println("Halo! ini private function")
	fmt.Printf("name: %s\nage: %d\naddress: %s\n", p.Name, p.Age, p.Address)
}

func Greet() {
	fmt.Println("Halo! ini public function")
}
