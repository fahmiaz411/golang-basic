package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) say (str string){
	fmt.Println("hello " + str + " i'm " + customer.Name)
}

func main() {
	var person = Customer{
		Name:    "fahmi",
		Address: "kp",
		Age:     10,
	}

	person.say("ega")

	// fmt.Println(person.Name)
}