package main

import "fmt"

func random() interface{} {
	if true {
		return 1
	} else {

		return "Ok"
	}

}

func init() {
	result := random()
	var res interface{}
	switch result.(type){
		case string:
			res = result.(string)
		case int:
			res = result.(int)
	}
	fmt.Println(res)
}