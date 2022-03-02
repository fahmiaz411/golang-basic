package pack

import (
	"container/ring"
	"fmt"
	"strconv"
)

func CircularList() {
	var data *ring.Ring = ring.New(5)

	for i:= 0; i < data.Len(); i++{
		 data.Value = "value " + strconv.FormatInt(int64(i), 10)
		 data = data.Next()
	}

	data.Do(func(i interface{}) {
		fmt.Println(i)
	})
}