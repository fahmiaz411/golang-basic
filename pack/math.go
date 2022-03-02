package pack

import (
	"fmt"
	"math"
)

func Maths() {
	fmt.Println("round", math.Round(1.7))
	fmt.Println("round", math.Round(1.3))
	fmt.Println("floor", math.Floor(1.7))
	fmt.Println("ceil", math.Ceil(1.1))

	fmt.Println("max", math.Max(1, 2))
	fmt.Println("min", math.Min(1,2))
}