package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
)

var client http.Client

const (
	URL = "https://www.nseindia.com/api/option-chain-indices?symbol=BANKNIFTY"
)

func init() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Got error while creating cookie jar %s", err.Error())
	}

	client = http.Client{
		Jar: jar,
	}
}

func getOptionChain() (map[string]interface{}, error) {
	url := "https://www.nseindia.com/api/option-chain-indices?symbol=BANKNIFTY"
	method := "GET"

	//client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var optionsLst map[string]interface{}
	if err := json.Unmarshal(body, &optionsLst); err != nil {
		return nil, err
	}
	return optionsLst, nil

}

func main() {
	var c []map[string]interface{}
	optionsLst, err := getOptionChain()
	if err != nil {
		panic(err)
	}
	parseMap(optionsLst, c)
	fmt.Println(c)
}

func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func parseMap(aMap map[string]interface{}, *c []map[string]interface{}) {
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			if key == "CE" || key == "PE" {
				c = append(c, val.(map[string]interface{}))
			}

			fmt.Println(key)
			parseMap(val.(map[string]interface{}), c)
		case []interface{}:
			fmt.Println(key)
			parseArray(val.([]interface{}), c)
		default:
			fmt.Println(key, ":", concreteVal)
		}
	}

}

func parseArray(anArray []interface{}, c []map[string]interface{}) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println("Index:", i)
			parseMap(val.(map[string]interface{}), c)
		case []interface{}:
			fmt.Println("Index:", i)
			parseArray(val.([]interface{}), c)
		default:
			fmt.Println("Index", i, ":", concreteVal)

		}
	}
}
