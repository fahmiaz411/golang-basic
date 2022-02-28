package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Married(val bool) {
	if val == true {
		p.Name = "Mr. " + p.Name
	}
}

func init() {

	var person = Person{
		Name: "fahmi",
	}

	person.Married(true)

	fmt.Println(person.Name)

}