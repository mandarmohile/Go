package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
	likes []string
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 523401,
		},
		likes: []string{
			"Hip-hop",
			"Jazz",
			"Metal",
		},
	}

	jim.updateName("Jimmy")
	updateLikes(jim)
	jim.print()

	updateLikesUsingPointerToPerson(&jim)
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
	(*pointerToPerson).likes[2] = "Heavy Metal"
}

func updateLikes(p person) {
	p.lastName = "Kong"     // does not update
	p.likes[2] = "Electric" // updates
}

func updateLikesUsingPointerToPerson(p *person) {
	(*p).lastName = "Kong"             // updates
	(*p).likes[2] = "Electric & Metal" // updates
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
