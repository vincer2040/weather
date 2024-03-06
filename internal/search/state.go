package search

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type States map[string]string

func GetStatesMap() (States, error) {
	url := "https://gist.githubusercontent.com/mshafrir/2646763/raw/8b0dbb93521f5d6889502305335104218454c2bf/states_hash.json"
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
	var states States
	err = json.Unmarshal(data, &states)
	if err != nil {
		return nil, err
	}
	return states, nil
}
