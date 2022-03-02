package pack

import (
	"fmt"
	"time"
)

func Time() {
	now := time.Now()
	
	fmt.Println(now)
	fmt.Println(now.Date())
	fmt.Println(now.Clock())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	// utc:= time.Date(2022, 10, 10, 10,10,10, 10, time.UTC)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2022-03-02")

	fmt.Print(parse)
}