package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"strings"
	"time"
)

var client http.Client

const (
	NSE_GAINER_URL = "https://www1.nseindia.com/live_market/dynaContent/live_analysis/gainers/niftyGainers1.json"
	NSE_LOSER_URL  = "https://www1.nseindia.com/live_market/dynaContent/live_analysis/losers/niftyLosers1.json"
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

type Symbolinfo struct {
	Symbol    string    `json:"symbol"`
	Series    string    `json:"series"`
	Openprice float64   `json:"openPrice"`
	Highprice float64   `json:"highPrice"`
	Lowprice  float64   `json:"lowPrice"`
	Lastprice float64   `json:"ltp"`
	Prevprice float64   `json:"previousPrice"`
	Percent   float64   `json:"netPrice"`
	Tradeqty  float64   `json:"tradedQuantity"`
	Turnover  float64   `json:"turnoverInLakhs"`
	Anndate   time.Time `json:"lastCorpAnnouncementDate"`
	Ann       string    `json:"lastCorpAnnouncement"`
}

type Perfstocks struct {
	Data    []Symbolinfo `json:"data"`
	Rpttime time.Time    `json:"time"`
}

func connNSEperf(url string) ([]byte, error) {
	method := "GET"

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

	if len(body) == 0 {
		err = errors.New("No Data")
		return nil, err
	}

	return body, nil
}
func (perf *Perfstocks) Init(url string) {

	body, err := connNSEperf(url)
	if err != nil {
		panic(err)
	}
	//fmt.Println(body)
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}
	strs := dat["data"].([]interface{})

	rptTime, err := ConvStrDate("Jan 02, 2006 15:04:05", dat["time"].(string))
	fmt.Println("Trading date : ", rptTime)
	perf.Rpttime = rptTime
	for _, value := range strs {
		//fmt.Println(value)
		var symbol Symbolinfo
		val := value.(map[string]interface{})

		symbol.Symbol = fmt.Sprint(val["symbol"])
		symbol.Series = fmt.Sprint(val["series"])
		symbol.Openprice, err = ConvStrCurrency(val["openPrice"].(string))
		symbol.Highprice, err = ConvStrCurrency(val["highPrice"].(string))
		symbol.Lowprice, err = ConvStrCurrency(val["lowPrice"].(string))
		symbol.Lastprice, err = ConvStrCurrency(val["ltp"].(string))
		symbol.Prevprice, err = ConvStrCurrency(val["previousPrice"].(string))
		symbol.Percent, err = ConvStrCurrency(val["netPrice"].(string))
		symbol.Tradeqty, err = ConvStrCurrency(val["tradedQuantity"].(string))
		symbol.Turnover, err = ConvStrCurrency(val["turnoverInLakhs"].(string))
		symbol.Turnover, err = ConvStrCurrency(val["turnoverInLakhs"].(string))
		symbol.Ann = fmt.Sprint(val["lastCorpAnnouncement"])
		symbol.Anndate, err = ConvStrDate("02-Jan-2006", val["lastCorpAnnouncementDate"].(string))
		perf.Data = append(perf.Data, symbol)
	}
}

func ConvStrCurrency(str string) (float64, error) {
	return strconv.ParseFloat(strings.ReplaceAll(str, ",", ""), 64)
}

func ConvStrDate(format string, str string) (time.Time, error) {
	return time.Parse(format, str)

}

func Getnsegainer() *Perfstocks {
	gainer := new(Perfstocks)
	gainer.Init(NSE_GAINER_URL)
	return gainer
}

func Getnseloser() *Perfstocks {
	loser := new(Perfstocks)
	loser.Init(NSE_LOSER_URL)
	return loser
}

func main() {
	gainer := Getnsegainer()
	fmt.Println(gainer)

	loser := Getnseloser()
	fmt.Println(loser)
}
