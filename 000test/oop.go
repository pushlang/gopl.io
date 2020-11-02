package main

import (
	"fmt"
)

func (p person) info() {
	fmt.Printf("First name: %s, Last name: %s, Address: %v\n", p.firstName, p.lastName, p.address)
}

func (c coder) info() {
	c.person.info()
	fmt.Printf("Position: %s, Salary:%d\n", c.position, c.salary)
}

func (s speciality) info() {
	s.coder.info()
	fmt.Printf("Speciality: %s\n", s.sName)
}

func (c company) info() {
	fmt.Printf("Company name: %s\n", c.cName)
	for _, pers := range c.personal {
		pers.person.info()
	}
}

type company struct {
	cName string
	address
	personal []speciality
}

type address struct {
	city   string
	street string
}

type person struct {
	firstName string
	lastName  string
	address
}

type coder struct {
	person
	position string
	salary   int
}

type speciality struct {
	coder
	sName string
}

func main() {
	//cm := make(map[speciality]coder)
	//pm := make(map[coder]person)
	//am := make(map[person]address)

	a := make([]address, 10)
	p := make([]person, 10)
	c := make([]coder, 10)
	s := make([]speciality, 10)
	co := make([]company, 10)

	id5 := 5
	a[id5] = address{city: "Los-Angeles", street: "Main str."}
	p[id5] = person{firstName: "Elon", lastName: "Musk", address: a[id5]}
	c[id5] = coder{person: p[id5], position: "Devop", salary: 5000}
	s[id5] = speciality{coder: c[id5], sName: "Frontend"}
	id8 := 8
	a[id8] = address{city: "Denver", street: "Central str."}
	p[id8] = person{firstName: "Bill", lastName: "Gates", address: a[id8]}
	c[id8] = coder{person: p[id8], position: "Boss", salary: 10000}
	s[id8] = speciality{coder: c[id8], sName: "Billionaire"}

	co[0] = company{
		cName:   "Tesla",
		address: a[id5],
		personal: []speciality{
			s[id5],
			s[id8],
		},
	}

	fmt.Println("----------person----------")
	//p[5].info()
	fmt.Println("----------coder-----------")
	//c[5].info()
	fmt.Println("--------speciality--------")
	//s[5].info()
	fmt.Println("--------company--------")
	co[0].info()
	fmt.Println()

	for _, spec := range s {
		if spec != (speciality{}) {
			fmt.Printf("%s, %s, %v\n", spec.lastName, spec.position, spec.address)
		}
	}
}

