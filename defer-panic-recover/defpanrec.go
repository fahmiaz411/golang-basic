package main

import "fmt"

func log() {
	err := recover()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("selesai")
}

func run(error bool){
	defer log()
	if error {
		panic("ERROR")
	}
	fmt.Println("run")
}

func main() {
	run(false)
	fmt.Println("next")
}