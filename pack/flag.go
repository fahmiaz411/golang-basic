package pack

import (
	"flag"
	"fmt"
)

func init() {
	host := flag.String("host", "root", "Put host")
	flag.Parse()
	fmt.Println(*host)
}