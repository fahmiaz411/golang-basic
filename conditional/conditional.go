package main

import "fmt"

func main() {
	var name = "ega"

	if name == "fahmi" {
		fmt.Println("hello")
	} else if name == "ega" {
		fmt.Println("hi ega")
	} else {
		fmt.Println("oh no")
	}

	if l:=1;l == 1{
		fmt.Println(l)
	}

	switch name {
		case "fahmi":
			fmt.Println("hello, fahmi")
		case "ega":
			fmt.Println("hello, ega")
		case "soni":
			fmt.Println("hello, soni")
		default: fmt.Println("nothing")
	}

	switch l:=1; l > 0 {
		case true: fmt.Println("true")
		case false: fmt.Println("false")
	}

	switch {
		case name == "ega":
			fmt.Println("ega")
		case name == "fahmi":
			fmt.Println("fahmi")
		default: fmt.Println("nama tidak terdaftar")
	}
}