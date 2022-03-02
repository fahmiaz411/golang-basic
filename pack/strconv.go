package pack

import (
	"fmt"
	"strconv"
)

func StrConv() {

	valInt, _ := strconv.Atoi("255")
	fmt.Printf( "%T\n", uint8(valInt))
}