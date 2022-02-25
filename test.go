package main

import "fmt"

type Data struct{
	Name string
}

func main() {
	var data = Data {
		Name: "fahmi",
	}
	fmt.Println(data)
}