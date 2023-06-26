package main

import "fmt"

type contactinfo struct {
	email    string
	postCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactinfo // Can add structs in structs
}

func main() {
	// bob := person{"Bob", "Bobson"}                        // Option #1 - Not recommended
	// dug := person{firstName: "Dug", lastName: "Stanhope"} // Option #2

	// bob.firstName = "bobby" // Update

	// fmt.Println(bob)
	// fmt.Println(dug)
	// fmt.Println(dug.firstName)
	// fmt.Printf("%+v", dug) // Print types

	timmy := person{
		firstName: "Timmy",
		lastName:  "Tim",
		contact: contactinfo{
			email:    "jim@southpark.com",
			postCode: "ST16 2MY",
		},
	}
	fmt.Printf("%+v", timmy)

}
