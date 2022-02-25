package main

import "fmt"

func main() {
	var arr = []string{
		"hello",
		"hey",
		"ello",
	}

	for i := 0; i < len(arr); i++ {
		if i == 0 {
			continue
		}
		fmt.Println(arr[i])
		
		if i == 1 {
			break
		}
	}

	for _, data := range arr {
		fmt.Println(data)
	}
}