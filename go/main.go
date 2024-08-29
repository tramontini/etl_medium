package main

import (
	"encoding/json"
	"fmt"
	"go-medium-etl/model"
	"io"
	"log"
	"net/http"
)

func main() {

	url := "https://api.coinpaprika.com/v1/tickers/btc-bitcoin?quotes=USD,BRL"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var ticker model.Ticker

	err = json.Unmarshal([]byte(body), &ticker)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %s, Name: %s, Symbol: %s\n", ticker.ID, ticker.Name, ticker.Symbol)
	for fiatID, quote := range ticker.Quotes {
		fmt.Printf("Fiat: %s, Price: %.2f, ATH Date: %s\n", fiatID, quote.Price, quote.LastUpdatedAt)
	}
}
