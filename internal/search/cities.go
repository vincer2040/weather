package search

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type City struct {
	Zip   int    `json:"zip_code"`
	City  string `json:"city"`
	State string `json:"state"`
}

type Cities []City

func GetCities() (Cities, error) {
	url := "https://raw.githubusercontent.com/millbj92/US-Zip-Codes-JSON/master/USCities.json"
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body := req.Body
	defer body.Close()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	var cities Cities
	err = json.Unmarshal(data, &cities)
	if err != nil {
		return nil, err
	}
	return cities, nil
}
