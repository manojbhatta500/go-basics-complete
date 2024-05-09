package main

import "fmt"

func main() {
	person1 := Person{firstName: "manoj", lastName: "Bhatta", age: 30, contact: Contact{
		email:   "helloworld@gmail.com",
		phone:   "43453435654",
		address: "himanchal pradesh",
	}}
	fmt.Println(person1.contact.phone)

}

type Contact struct {
	email   string
	phone   string
	address string
}

type Person struct {
	firstName string
	lastName  string
	age       int
	contact   Contact
}
