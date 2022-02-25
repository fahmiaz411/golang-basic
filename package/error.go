package main

import (
	"errors"
	"fmt"
)

func Divider(val int, div int ) (int, error) {
	if div == 0 {
		return 0, errors.New("Divider not valid")
	} else {
		return val / div, nil
	}
}

func main (){
	val, err:= Divider(10,2)
	if val == 0 {

		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
}