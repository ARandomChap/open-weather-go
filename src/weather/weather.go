package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/AR4ndomCh4p/open-weather-go/src/location"
)

func Weather(apiKey string, zip string, countryCode string, when string, hour int) string {
	var coords map[string]interface{}
	loc := location.Location(apiKey, zip, countryCode)

	json.Unmarshal([]byte(loc), &coords)
	lat := fmt.Sprint(coords["lat"])
	lon := fmt.Sprint(coords["lon"])

	req, err := http.Get("https://api.openweathermap.org/data/2.5/onecall?lat=" + lat + "&lon=" + lon + "&appid=" + apiKey)
	if err != nil {
		fmt.Printf("err: %v", err.Error())
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("err: %v", err.Error())
	}

	var weather map[string]interface{}
	json.Unmarshal(body, &weather)

	var res string
	switch when {
	case "current":
		res = current(weather)
	case "hourly":
		res = hourly(weather, hour)
	}

	return res
}

func current(current map[string]interface{}) string {
	return fmt.Sprint(current["current"].(map[string]interface{})["weather"].([]interface{})[0].(map[string]interface{})["description"])
}

func hourly(hourly map[string]interface{}, hour int) string {
	return fmt.Sprint(hourly["hourly"].([]interface{})[hour].(map[string]interface{})["weather"].([]interface{})[0].(map[string]interface{})["description"])
}
