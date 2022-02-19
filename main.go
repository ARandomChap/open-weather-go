package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var lol map[string]interface{}

	json.Unmarshal([]byte(loc), &lol)

	fmt.Println(lol["lat"].(map[string]interface{}))

	req, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=53.480759&lon=-2.242631&appid=5e59c3037bf76fb39a38b0257d2dfbad")
	if err != nil {
		fmt.Printf("err: %v", err.Error())
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("err: %v", err.Error())
	}

	fmt.Println(string(body))
}

// Get preferred outbound ip of this machine
func location() string {
	loc, err := http.Get("http://api.openweathermap.org/geo/1.0/zip?zip=WF9,GB&appid=5e59c3037bf76fb39a38b0257d2dfbad")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(loc.Body)
	if err != nil {
		fmt.Printf("err: %v", err.Error())
	}
	
	return string(body)
}
