package main

import "fmt"

// Person is a struct that represents a person
type Person struct {
	firstName string
	lastName string
	// contact contactInfo // This is a nested/embedded struct
	contactInfo // This is the same as above, but it's a shorthand
}

type contactInfo struct {
	email string
	zipCode int
}



func main() {
	// Create a new person

	// alex := Person{"Alex", "Anderson"} // This is also valid (tutorial uses this one)
	alex := Person{firstName: "Alex", lastName: "Anderson"} // This one is better because it's more readable
	// var alex Person -> another way of declaring a struct

	// Print the name
	fmt.Println(alex)
	fmt.Printf("%+v\n", alex) // %+v prints the field names as well

	// Update the name
	alex.firstName = "Alexander"
	fmt.Printf("%+v\n", alex)

	// Create a new person with contact info
	jim := Person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	fmt.Println("Before update:")
	jim.print()
	jim.updateName("Jimmy")
	fmt.Println("After update:")
	jim.print()

	//Using Pointers
	fmt.Println("Before update:")
	jim.print()
	jim.newUpdateName("Jimmy")
	fmt.Println("After update:")
	jim.print()

	// Getting the pointer
	jimPointer := &jim // The "&" operator gives the memory address of the value this variable is pointing at
	jimValue := *jimPointer // The "*" operator gives the value this memory address is pointing at

	fmt.Println("Pointer:")
	fmt.Printf("%+v\n", jimPointer)
	fmt.Println("Value:")
	fmt.Printf("%+v\n", jimValue)

}

// This is a receiver function. It's a function that can be called on a specific type
func (p Person) print() { 
	fmt.Printf("%+v\n", p)
}


// This function may not work as expected because Go is a pass by value language
// This means that the function will create a copy of the struct and modify that copy
// The original struct will not be modified
func (p Person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// To fix this, we can use pointers
func (pointer *Person) newUpdateName(newFirstName string) {
	(*pointer).firstName = newFirstName
}