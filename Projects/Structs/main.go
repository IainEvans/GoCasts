package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	bob := person{"Bob", "Bobson"} // Not recommended
	dug := person{firstName: "Dug", lastName: "Stanhope"}

	fmt.Println(bob)
	fmt.Println(dug)
	fmt.Println(dug.firstName)

}
