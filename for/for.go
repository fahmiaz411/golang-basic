package main

import "fmt"

func BreakContinueFor (array []string) {
	defer fmt.Println("")
	for i := 0; i < len(array); i++ {
		if i == 0 {
			continue
		}
		fmt.Println(array[i])
		
		if i == 1 {
			break
		}
	}
}

func RangeFor(array []string){
	defer fmt.Println("")
	for _, data := range array {
		fmt.Println(data)
	}
}

func WhileFor(array []string){
	defer fmt.Println("")
	loop := 0
	for {
		if loop < len(array){
			fmt.Println(array[loop])
			loop++
		} else {
			break
		}
	}
}

func main() {
	var arr = []string{
		"hello",
		"hey",
		"l",
	}

	BreakContinueFor(arr)
	RangeFor(arr)
	WhileFor(arr)
	
}