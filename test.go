package main

import (
	// "fmt"
	"fmt"
	_ "golang-basic/pack"
)


func main() {

	var array = []interface{}{
		"yo",
		1,
		true,
	}

	array2 := append(array, "hey")
	fmt.Println(array2, array)

}