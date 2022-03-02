package location

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get preferred outbound ip of this machine
func Location(apikey string, zip string, countryCode string) string {
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
