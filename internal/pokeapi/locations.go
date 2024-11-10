package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := baseUrl + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationRes := Location{}
		err := json.Unmarshal(val, &locationRes)
		if err != nil {
			return Location{}, err
		}
		return locationRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	locationRes := Location{}
	err = json.Unmarshal(data, &locationRes)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)
	return locationRes, nil
}

func (c *Client) ListLocations(page *string) (LocationsList, error) {
	url := baseUrl + "/location-area"
	if page != nil {
		url = *page
	}

	if val, ok := c.cache.Get(url); ok {
		locationRes := LocationsList{}
		err := json.Unmarshal(val, &locationRes)
		if err != nil {
			return LocationsList{}, err
		}

		return locationRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationsList{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationsList{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsList{}, err
	}

	locationRes := LocationsList{}
	err = json.Unmarshal(data, &locationRes)
	if err != nil {
		return LocationsList{}, err
	}

	c.cache.Add(url, data)
	return locationRes, nil
}
