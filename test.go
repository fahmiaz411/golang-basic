package main

import (
	"fmt"
	_ "golang-basic/pack"
	"os"
)

// type Data struct{
// 	Name string
// }

func main() {
	// var data = Data {
	// 	Name: "fahmi",
	// }
	// fmt.Println(data, helper.Other(), helper.Version)
	// fmt.Println(database.GetConnection(), database.Sample)

	fmt.Println(os.Args)
}