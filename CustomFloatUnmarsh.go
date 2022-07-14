package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Foo struct {
	Price CustomFloat64 `json : "price" `
}
type CustomFloat64 float64

func (cf *CustomFloat64) UnmarshalJSON(b []byte) error {

	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	ft, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	*cf = CustomFloat64(ft)

	return nil
}

func main() {
	b := []byte(`{ "price" : "10,0000.00" }`)
	f := Foo{}

	json.Unmarshal(b, &f)
	fmt.Println(f)

}
