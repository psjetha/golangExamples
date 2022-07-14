package main

import "fmt"

func main() {

    messages := make(chan string, 2)

    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
    
    messages <- "buffered-2"
    messages <- "channel-2"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
