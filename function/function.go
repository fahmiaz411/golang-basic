package main

import (
	"fmt"
)

type Filter func (string) string

func main() {
	say()
	param("fahmi", "aziz")

	var _, l = returnData()
	var _, d = returnVar()

	slice := []int{
		10,20,
	}

	var args = varArgs
	var num = args(slice...)
	fmt.Println(l,d, num)

	fmt.Println(functionParams("anj", func (name string) string {
		if name == "anj" {
			return "..."
		} else {
			return name
		}
	}))
	
}

func say() {
	fmt.Println("hello")
	return
}

func param(f string, l string){
	fmt.Println("hello", f, l)
}

func returnData() (string, string) {
	return "hello", "you"
}

func returnVar() (f string,l string) {
	f = "hey"
	l = "u"

	return
}

func varArgs (numbers... int)int{
	var result int = 0

	for _, num := range numbers {
		result+= num
	}

	return result
}

func functionParams (name string, filter Filter) string {

	return "hello" + filter(name) 

}
