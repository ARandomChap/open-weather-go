package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	apiKey := flag.String("api-key", "", "The open-weather API Key")
	zip := flag.String("zip", "", "The first part of you zip code, e.g. ('WF9')")
	countryCode := flag.String("country-code", "", "The country code related to your zip code, e.g ('GB, US, ...')")

	flag.Parse()

	if *apiKey == "" {
		log.Fatalln("api-key was not set")
	}
	if *zip == "" {
		log.Fatalln("zip was not set")
	}
	if *countryCode == "" {
		log.Fatalln("country-code was not set")
	}

	resp := weather(*apiKey, *zip, *countryCode)

	fmt.Println(resp)
}

func weather(apiKey string, zip string, countryCode string) string {
	var coords map[string]interface{}
	loc := location(apiKey, zip, countryCode)

	json.Unmarshal([]byte(loc), &coords)
	lat := fmt.Sprint(coords["lat"])
	lon := fmt.Sprint(coords["lon"])

	req, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lon + "&appid=" + apiKey)
	if err != nil {
		fmt.Printf("err: %v", err.Error())
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("err: %v", err.Error())
	}

	var weather map[string]interface{}
	json.Unmarshal(body, &weather)

	main := fmt.Sprint(weather["weather"])

	return main
}

// Get preferred outbound ip of this machine
func location(apikey string, zip string, countryCode string) string {
	loc, err := http.Get("http://api.openweathermap.org/geo/1.0/zip?zip=" + zip + "," + countryCode + "&appid=" + apikey)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(loc.Body)
	if err != nil {
		fmt.Printf("err: %v", err.Error())
	}

	return string(body)
}
