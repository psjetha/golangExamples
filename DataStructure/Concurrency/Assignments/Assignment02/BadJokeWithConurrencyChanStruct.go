package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Response struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func main() {
	start := time.Now()
	c := make(chan Response)
	defer close(c)

	for i := 1; i <= 50; i++ {
		go func() {
			getBadJoke(c)
		}()
	}

	for i := 1; i <= 50; i++ {
		msg := <-c
		fmt.Println("\n JOKE ", msg.Joke, msg.ID, msg.Status)
	}

	elapsed := time.Since(start)
	fmt.Printf("Processes took %s", elapsed)
}
func getBadJoke(c chan Response) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject Response
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		return
	}

	c <- responseObject

}
