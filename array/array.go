package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "fahmi"
	names[1] = "soni"
	names[2] = "ega"

	fmt.Println(names[0])

	var values = [...]int8{
		10,
		20,
		30,
	}
	values[2] = 10
	var slice = []int{1,2,3}

	fmt.Println(cap(slice))

}