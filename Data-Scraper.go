package main

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	"encoding/json"
)

type Inner struct {
	Name		string		`json:"SC_FULLNM"`
	SCID		string		`json:"DISPID"`
	Price		FlexInt		`json:"pricecurrent"`
	PriceChange	FlexInt		`json:"pricechange"`
	Percentage	FlexInt		`json:"pricepercentchange"`
	High		FlexInt		`json:"HP,omitempty"`
	Low			FlexInt		`json:"LP,omitempty"`
	Volume		FlexInt		`json:"VOL"`
	LastUpdated	CustomTime	`json:"lastupd"`
	LCL			FlexInt		`json:"lower_circuit_limit,omitempty"`
	UCL			FlexInt		`json:"upper_circuit_limit,omitempty"`
}

type Collection struct {
	Status			FlexInt		`json:"code"`
	Data			Inner		`json:"data"`
	URL				string		`json:"url"`
	LastScraped		time.Time	`json:"lastscraped"`
}

type CustomTime struct {
	time.Time
}

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	ctLayout := "2006-01-02 15:04:05"
	timezone, _ := time.LoadLocation("Asia/Kolkata")

	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	ct.Time, err = time.ParseInLocation(ctLayout, string(b), timezone)
	return
}

type FlexInt float32

func (fi *FlexInt) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, err := strconv.ParseFloat(string(b), 32)
	*fi = FlexInt(i)
	return
}

func getData(url string) (data *Collection, err error) {
	
	client := &http.Client{
        Timeout: 30 * time.Second,
	}

    response, err := client.Get(url)
    if err != nil {
		err = fmt.Errorf("Error at response: %v", err)
        return
    }
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		err = fmt.Errorf("Error at decoding: %v", err)
        return
	}

	if data.Status != FlexInt(200) {
		err = fmt.Errorf("Response Error: Status Not 200!")
	}
	data.URL = url

	data.LastScraped = time.Now()
	
	return
}

/* Sample JSON from moneycontrol
{
	"code": "200",
	"data": {
		  "SC_FULLNM": "Purshottam Investofin",
		  "DISPID": "PI39",
		  "pricecurrent": "11.94",
		  "VOL": "143874",
		  "pricepercentchange": "-5.4632",
		  "pricechange": "-0.6900",
		  "lastupd": "2020-12-09 16:00:01",
		  "lower_circuit_limit": "11.37",
		  "upper_circuit_limit": "13.89",
		  "HP": "13.89",
		  "LP": "11.38",
	}
  }
*/