package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const binancePriceUrl = "https://api.binance.com/api/v3/ticker/price"

type Data struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func main() {
	response, err := http.Get(binancePriceUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bytesBuffer := bytes.Buffer{}
	buffer := make([]byte, 1024)
	for {
		bytesRead, err := response.Body.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		bytesBuffer.Write(buffer[:bytesRead])

		if err == io.EOF {
			break
		}
	}

	var arrData []Data
	err = json.Unmarshal(bytesBuffer.Bytes(), &arrData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(arrData, len(arrData))
}
