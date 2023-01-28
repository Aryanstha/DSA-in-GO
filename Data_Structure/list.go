package main

// importing the list and fmt packages

import (
	"container/list"
	"fmt"
)

// main function
func main() {
	var intList list.List
	intList.PushBack(1)
	intList.PushBack(2)
	intList.PushBack(3)
	intList.PushBack(4)
	intList.PushBack(5)

	// iterating over the list
	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
