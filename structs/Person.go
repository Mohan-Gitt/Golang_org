package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	Mohan := person{"Mohan", "Kumar G M", 26}
	Rohit := person{firstName: "Rohit", lastName: "Sharma", age: 35}

	var John person

	fmt.Println(Mohan)
	fmt.Println(Rohit)

	John.firstName = "John"
	John.lastName = "Robertt"
	John.age = 29

	fmt.Println(John)
	fmt.Printf("%+v", John)
}
