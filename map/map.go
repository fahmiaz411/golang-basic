package main

import "fmt"

func main() {

	var person = map[string]string{
		"name":    "fahmi",
		"address": "bogor",
	}

	person["age"] = "10"
	fmt.Println(person["age"])

	book:= make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "fahmi"
	book["wrong"] = "salah"

	delete(book, "wrong")

	fmt.Println(len(book))
}