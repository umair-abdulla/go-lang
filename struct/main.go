package main

import "fmt"

type contactInfo struct {
	email string
	pin int
}

type person struct {
	firstName string
	lastName  string
	// contact contactInfo
	contactInfo
}

func main() {
	// alex := person{firstName: "alex", lastName: "son"}

	// var alex person

	// alex.firstName = "alex"
	// alex.lastName = "son"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "jim",
		lastName: "hobbs",
		contactInfo: contactInfo{
			email: "jim@r.c",
			pin: 673032,
		},
	}

	// jimPointer := &jim
	// jimPointer.updateName("jack")
	jim.updateName("jack")
	jim.print()
	
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print(){
	fmt.Printf("%+v", p)
}