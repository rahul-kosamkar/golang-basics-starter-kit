package main

import (
	"fmt"
	"strconv"
)

// Define Struct

type Person struct {
	// firstName string
	// lastName  string
	// gender    string
	// age       int
	// city      string

	firstName, lastName, gender, city string
	age                               int
}

// Greeting Method (Value Reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " " + "and I am " + strconv.Itoa(p.age)
}

// hasBirthday Method (Pointer Reciever)
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Rahul", lastName: "Kosamkar", gender: "Male", age: 26, city: "Mumbai"}
	fmt.Println(person1)
	fmt.Println(person1.firstName)

	person1.hasBirthday()
	fmt.Println(person1.greet())
}
