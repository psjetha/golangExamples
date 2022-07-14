package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	defer close(c1)
	// c2 := make(chan string)
	start := time.Now()

	go func() {
		time.Sleep(time.Millisecond * 500)
		c1 <- "Every 500ms"
	}()

	go func() {
		c1 <- "Every two sec"
		time.Sleep(time.Second * 2)
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		default:
			continue

		}
	}

	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)

}
