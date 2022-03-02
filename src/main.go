package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/AR4ndomCh4p/open-weather-go/src/weather"
)

func main() {
	apiKey := flag.String("api-key", "", "The open-weather API Key")
	zip := flag.String("zip", "", "The first part of you zip code, e.g. ('WF9')")
	countryCode := flag.String("country-code", "", "The country code related to your zip code, e.g ('GB, US, ...')")
	when := flag.String("when", "", "What forecast would you like? e.g. ('current, hourly, ...')")

	hour := flag.NewFlagSet("hour", flag.ExitOnError)
	hour1 := hour.Int("time", -1, "test hour")

	fmt.Print(*hour1)

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
	if *when == "" {
		log.Fatalln("when was not set")
	}

	resp := weather.Weather(*apiKey, *zip, *countryCode, *when, *hour1)

	fmt.Println(resp)
}
