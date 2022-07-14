package main

import (
	"fmt"
	"time"
)

func main() {

	go count("Jerry")
	count("Tom")

}

func count(str string) {

	for i := 1; true; i++ {
		fmt.Println(i, str)
		time.Sleep(time.Microsecond * 500)

	}

}
