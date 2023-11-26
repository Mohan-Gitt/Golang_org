package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	age       int
	// contact contactInfo
	contactInfo
}

func main() {
	jim := Person{
		firstName: "Jim",
		lastName:  "Party",
		age:       19,
		// contact : contactInfo
		contactInfo: contactInfo{
			email:   "jim19@gmail.com",
			zipCode: 570321,
		},
	}

	// jimPointer := &jim // give us memory address jim is pointing at

	jim.updateFirstName("Jimmy")
	jim.print()

}

// *Person -> pointer we are working with the Person type
// *pointerToPerson -> Actual pointer - give the value jim is pointing at the memory address
func (pointerToPerson *Person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
