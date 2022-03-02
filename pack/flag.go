package pack

import (
	"flag"
	"fmt"
)

func Flags() {
	host := flag.String("host", "root", "Put host")
	flag.Parse()
	fmt.Println(*host)
}