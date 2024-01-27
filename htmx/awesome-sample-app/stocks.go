package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const PolygonPath = "https://api.polygon.io"
const ApiKey = "apiKey=KpXGkgim6HyuF5IWvp0d2dy7s8EpUWuu"

type Stock struct {
	Ticker string  `json:"ticker"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
}

type Values struct {
	Open float64 `json:"open"`
}

func SearchTicker(ticker string) []Stock {
	resp, err := http.Get(
		PolygonPath + "/v3/reference/tickers?" + ApiKey + "&ticker=" + strings.ToUpper(ticker),
	)

	if err != nil {
		fmt.Printf("error getting the response; %s", err.Error())
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)

	data := struct {
		Results []Stock `json:"results"`
	}{}

	json.Unmarshal(body, &data)
	fmt.Printf("data returned by search; %T", data.Results)
	return data.Results
}

func GetDailyValues(ticker string) Values {
	resp, err := http.Get(
		PolygonPath + "/v1/open-close/" + strings.ToUpper(ticker) + "/2023-10-10/?" + ApiKey,
	)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)

	data := Values{}

	json.Unmarshal(body, &data)

	return data

}
