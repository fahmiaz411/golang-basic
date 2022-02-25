package main

import "fmt"

func New(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := New("fahmi")

	fmt.Println(person == nil)
}