package main

import "fmt"

func main() {
	type KTP string
	type NIK string

	var ktp KTP = "fahmi"
	fmt.Println(ktp)
}