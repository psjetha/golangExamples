package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("Jerry", c)

	for msg := range c {
		fmt.Println(msg)
	}

	/*	for
		{
		    msg, open := <-c
		    if !open {
		      break
		    }
		    fmt.Println(msg)
		  }
	*/
}

func count(str string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- str
		time.Sleep(time.Microsecond * 500)
	}
	close(c)
}
