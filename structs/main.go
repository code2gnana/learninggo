package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	// Step1 on to assign value to Struct Person
	// alex := person{"Alex", "Anderson"}

	// Step2 on to assign value to Struct Person
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// Step3 on to assign value to Struct Person

	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// Printing inside a main function
	// fmt.Printf("%+v", p)

	// Update First Name with Pointer and & sign
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")

	// Update First Name without address or normal assigment
	jim.updateName("jimmy")

	jim.print()
}

// Function to update first name with a pointer - updates the first name
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// Function to update first name without a pointer - doesn't updates the first name
// func (p person) updateName(newFirstName string) {
//	p.firstName = newFirstName
//}

//Struct with Receiver function
func (p person) print() {
	fmt.Printf("%+v", p)
}
