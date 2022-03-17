package main

import (
	// "fmt"
	"errors"
	"fmt"
	"os"
)

// func zip(a ,b []int) ([]int, []int) {
// 	max := math.Max(len(a), len(b))
// }

func AnyMapStrStr(r []interface{}) ([]map[string]string, error) {

	var nr []map[string]string

	for _, sl := range r {

		m := map[string]string{}
		
		n_sl, s := sl.(map[string]string)

		if s == false {
			return nil, errors.New("Convert failed to map[string]string")
		}

		for k, v := range n_sl{
			m[k] = v
		}

		nr = append(nr, m)
	}

	return nr, nil
 
}

func main() {
	
	// arr1 := []interface{}{
	// 	1,2,3,
	// }
	// arr2 := []interface{}{
	// 	4,5,6,
	// }


	// for _, d := range zipper.Zip(arr1, arr2) {
	// 	a1, a2 := d[0], d[1]
	// 	fmt.Println(a1,a2)
	// }

	// array := []int{
	// 	1,2,3,4,
	// }

	fmt.Println(os.Getenv("os"))
	
}