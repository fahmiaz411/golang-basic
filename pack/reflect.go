package pack

import (
	"fmt"
	"reflect"
	"strconv"
)

type Sample struct {
	Name string `required:"true" max: "10"`
	Age string `required:"true" max: "10"`
}

func isValid (data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			val := reflect.ValueOf(data).Field(i).Interface()
			if val == "" {
				return false
			}
		}
	}

	return true
}


func Reflect() {
	sample := Sample{
		Name: "fahmi",
	}

	var sampleType reflect.Type = reflect.TypeOf(sample)
	structField := sampleType.Field(0)
	
	parse, _ := strconv.ParseBool(structField.Tag.Get("required"))

	fmt.Println(sampleType.NumField())
	fmt.Println(parse)

	fmt.Println("validate", isValid(sample))
}