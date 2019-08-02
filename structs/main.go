package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

//go is a pass by value language
func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	//& is an operator: look at this variable give acces to memory where this is
	//jimPointer is memory address of jim
	//jimPointer := &jim
	jim.updateName("jimmy")
	jim.print()

	//*blah turn address into value (but remember type vs operator/pointer)
	//&blah turn value into address

}

//* is give me value of this memory address
//*pointerToPerson is memory access point where jim lives
//*person is a pointer that points to a person type
//*person = in place of type this is a descirption of type. * here is not an operator but a description
//jimPOinter and pointer to person is the same

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

//BIG GOTCHA IN GO slices and struts are very different for pointers vs place in a slice

func (p person) print() {
	fmt.Printf("%+v", p)
}
