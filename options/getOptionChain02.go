package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/http/cookiejar"
	"time"
)

var client http.Client

type OptionChain struct {
	Records  Record   `json:"records"`
	Filtered Filtered `json:"filtered"`
}

type Filtered struct {
	Data     []Data `json:"data"`
	CETotVol Volume `json:"CE"`
	PETotVol Volume `json:"PE"`
}

type Record struct {
	ExpiryDates     []CustomTime `json:"expiryDates"`
	Data            []Data       `json:"data"`
	Timestamp       string       `json:"timestamp"`
	UnderlyingValue float64      `json:"underlyingValue"`
	StrikePrices    []float64    `json:"strikePrices"`
	Index           Index        `json:"index"`
}
type Data struct {
	StrikePrice int        `json:"strikePrice"`
	ExpiryDate  CustomTime `json:"expiryDate"`
	PE          *Option    `json:"PE,omitempty"`
	CE          *Option    `json:"CE,omitempty"`
}

type Volume struct {
	TotOI  int `json:"totOI"`
	TotVol int `json:"totVol"`
}

type Option struct {
	StrikePrice           int        `json:"strikePrice"`
	ExpiryDate            CustomTime `json:"expiryDate"`
	Underlying            string     `json:"underlying"`
	Identifier            string     `json:"identifier"`
	OpenInterest          float64    `json:"openInterest"`
	ChangeinOpenInterest  float64    `json:"changeinOpenInterest"`
	PchangeinOpenInterest float64    `json:"pchangeinOpenInterest"`
	TotalTradedVolume     int        `json:"totalTradedVolume"`
	ImpliedVolatility     float64    `json:"impliedVolatility"`
	LastPrice             float64    `json:"lastPrice"`
	Change                float64    `json:"change"`
	PChange               float64    `json:"pChange"`
	TotalBuyQuantity      int        `json:"totalBuyQuantity"`
	TotalSellQuantity     int        `json:"totalSellQuantity"`
	BidQty                int        `json:"bidQty"`
	Bidprice              float64    `json:"bidprice"`
	AskQty                int        `json:"askQty"`
	AskPrice              float64    `json:"askPrice"`
	UnderlyingValue       float64    `json:"underlyingValue"`
}

type Index struct {
	Key           string  `json:"key"`
	Index         string  `json:"index"`
	IndexSymbol   string  `json:"indexSymbol"`
	Last          float64 `json:"last"`
	Variation     float64 `json:"variation"`
	PercentChange float64 `json:"percentChange"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	PreviousClose float64 `json:"previousClose"`
	YearHigh      float64 `json:"yearHigh"`
	YearLow       float64 `json:"yearLow"`
	Pe            string  `json:"pe"`
	Pb            string  `json:"pb"`
	Dy            string  `json:"dy"`
	Declines      string  `json:"declines"`
	Advances      string  `json:"advances"`
	Unchanged     string  `json:"unchanged"`
}

type CustomTime struct {
	time.Time
}

const ctLayout = "02-Jan-2006"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	ct.Time, err = time.ParseInLocation(ctLayout, string(b), loc)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(ctLayout)), nil
}

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

func getOptionChain() (OptionChain, error) {
	method := "GET"

	req, err := http.NewRequest(method, URL, nil)

	if err != nil {
		fmt.Println(err)
		return OptionChain{}, err
	}
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return OptionChain{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return OptionChain{}, err
	}
	var optionsLst OptionChain
	if err := json.Unmarshal(body, &optionsLst); err != nil {
		return OptionChain{}, err
	}
	return optionsLst, nil

}

func main() {
	optionsLst, err := getOptionChain()
	if err != nil {
		panic(err)
	}
	for {
		time.Sleep(time.Second * 20)
		cemax := optionsLst.Records.Data[0].CE
		pemax := optionsLst.Records.Data[0].PE

		num := 100
		step := 10
		nearest := nearest_option(optionsLst.Records.Index.Last, num)
		var optionChainList []Data
		for _, data := range optionsLst.Records.Data {

			if data.ExpiryDate.String() == optionsLst.Records.ExpiryDates[0].String() {
				if (nearest-(step*num)) <= data.StrikePrice && data.StrikePrice <= (nearest+(step*num)) {

					optionChainList = append(optionChainList, data)
				}
				if data.CE != nil {
					if data.CE.OpenInterest > cemax.OpenInterest {
						cemax = data.CE
					}

				}
				if data.PE != nil {
					if data.PE.OpenInterest > pemax.OpenInterest {
						pemax = data.PE
					}
				}

			}
		}

		fmt.Println("CE Stricke Price Support", cemax, cemax.OpenInterest)
		fmt.Println("PE Stricke Price Resistance", pemax, pemax.OpenInterest)
		/*	for _, item := range optionChainList {
				fmt.Println("CE :", item.StrikePrice, item.CE.OpenInterest, item.CE.TotalBuyQuantity, item.CE.TotalSellQuantity, item.CE.TotalTradedVolume)
				fmt.Println("PE :", item.StrikePrice, item.PE.OpenInterest, item.PE.TotalBuyQuantity, item.PE.TotalSellQuantity, item.PE.TotalTradedVolume)

			}
		*/
	}
}

func nearest_option(ind float64, step int) int {
	return int((math.Ceil(ind/float64(step)) * float64(step)))
}
