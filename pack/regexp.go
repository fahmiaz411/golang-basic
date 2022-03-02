package pack

import (
	"fmt"
	"regexp"
)

func Regex() {
	regex := regexp.MustCompile(`f[a-z]i`)

	fmt.Println(regex.MatchString("fAi"))
	fmt.Println(regex.MatchString("fai"))
	fmt.Println(regex.MatchString("fui"))

	search := regex.FindAllString("fai fAi fui", 1)
	
	for _, data := range search{
		fmt.Println(data)
	}
}