/*
 A tuple is a finite sorted list of elements. It is a data structure that groups data.
  Tuples are typically immutable sequential collections
*/

package main

// importing fmt package
import (
	"fmt"
)

//gets the power series of integer a and returns tuple of square of a and cube of a
func powerSeries(a int) (int, int) {

	return a * a, a * a * a

}

func main() {

	var square int
	var cube int
	square, cube = powerSeries(3)

	fmt.Println("Square ", square, "Cube", cube)

}
