/*
A list is a sequence of elements. Each element can be connected to another with a link in a forward or backward direction.
Lists have a variable length and developer can remove or add elements more easily than an array.
*/

package main

// importing fmt and container list packages
import (
	"container/list"
	"fmt"
)

// main method
func main() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
