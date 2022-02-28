package main

import "fmt"

func run(val interface{}) interface{} {
	return val
}


func init() {
	fmt.Println(run(1))
}