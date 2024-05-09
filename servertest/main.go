package main

import "fmt"

type Person struct {
	name  string
	phone int
	job   string
}

func main() {
	p1 := Person{name: "manoj", phone: 8782719872, job: "GoLang Developer"}

	printPerson(&p1)

	fmt.Println("********************main Thread***************************")
	fmt.Println(p1.name)
	fmt.Println(p1.job)
	fmt.Println(p1.phone)

}

func printPerson(person *Person) {
	fmt.Println("**********************before changed****************")
	fmt.Println("Name :", person.name)
	fmt.Println("phone:", person.phone)
	fmt.Println("job:", person.job)

	x := 123412321321

	person.name = "salman khan"
	person.phone = x
	person.job = "Actor"

	fmt.Println("***************after changed****************")

	fmt.Println("Name :", person.name)
	fmt.Println("phone:", person.phone)
	fmt.Println("job:", person.job)

}
