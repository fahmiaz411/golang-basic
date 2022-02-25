package main

import "fmt"

type Has interface {
	GetPerson(string) string
}

type Person struct {
	Name string
}

func (person Person) GetPerson(say string) string {
	return say + person.Name
}

func say(has Has) {

	fmt.Println(has.GetPerson("hello "))
}


func main() {
	var person Person
	person.Name = "fahmi"
	say(person)
}