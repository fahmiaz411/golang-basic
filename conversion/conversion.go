package main

import "fmt"

func main() {
	var nilai32 int = 100000
	var nilai64 int64 = int64(nilai32)

	// var name = "Fahmi"
	fmt.Println(nilai32, uint32(nilai64), string("a"[0]))
}