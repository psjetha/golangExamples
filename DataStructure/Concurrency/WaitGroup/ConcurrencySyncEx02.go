package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		count("Jerry")
		wg.Done()
	}()

	wg.Wait()
}

func count(str string) {

	for i := 1; i <= 5; i++ {
		fmt.Println(i, str)
		time.Sleep(time.Microsecond * 500)

	}

}
