package pack

import (
	"container/list"
	"fmt"
)

func LinkedList() {
	data := list.New()
	data.PushBack("fahmi")
	data.PushBack("soni")
	data.PushBack("ega")
	data.PushFront("faisal")

	for e:= data.Front(); e != nil; e = e.Next(){
		fmt.Println(e.Value)
	}
}