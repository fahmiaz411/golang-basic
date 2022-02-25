package main

import "fmt"

func run(val interface{}) interface{} {
	return val
}


func main() {
	fmt.Println(run(1))
}