package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func change(address *Address, val string){
	address.Country = val
}

func main() {
	var address1 Address = Address{
		"Bogor",
		"Jawa Barat",
		"Indonesia",
	}
	var address2 *Address = &address1
	var address3 *Address = &address1

	*address2 = Address{
		"Bandung",
		"Jawa Barat",
		"Indonesia",
	}

	fmt.Println(address2, address1, address3)

	// new 

	var address4 = new(Address)
	fmt.Println(address4,)
	
	// function pointer 
	
	change(address2, "Europe")
	fmt.Println(address3)
	change(&address1, "Asia")
	fmt.Println(address2)

}