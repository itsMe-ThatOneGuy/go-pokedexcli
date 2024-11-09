package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"Previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocations(page *string) (Locations, error) {
	url := baseUrl + "/location-area"
	if page != nil {
		url = *page
	}

	if val, ok := c.cache.Get(url); ok {
		locationRes := Locations{}
		err := json.Unmarshal(val, &locationRes)
		if err != nil {
			return Locations{}, nil
		}

		return locationRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, nil
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, nil
	}

	locationRes := Locations{}
	err = json.Unmarshal(data, &locationRes)
	if err != nil {
		return Locations{}, nil
	}

	c.cache.Add(url, data)
	return locationRes, nil
}
