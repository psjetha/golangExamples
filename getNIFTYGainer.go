package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
)

var client http.Client

type ticker struct {
	symbol string `json:string symbol`
}

func init() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
	}

	client = http.Client{
		Jar: jar,
	}
}

func main() {

	url := "https://www1.nseindia.com/live_market/dynaContent/live_analysis/gainers/niftyGainers1.json"
	method := "GET"

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	//req.Header.Add("Host", "www1.nseindia.com")
	req.Header.Add("Accept", "*/*")
	//req.Header.Add("X-Requested-With", "X-Requested-With")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//	fmt.Println(string(body))
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}
	strs := dat["data"].([]interface{})
	for _, value := range strs {
		//fmt.Println(value)
		val := value.(map[string]interface{})
		fmt.Println(val["symbol"], val["openPrice"])
		for key, _ := range val {

			fmt.Println(key)
		}
	}

}
