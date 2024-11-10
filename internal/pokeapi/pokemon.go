package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemonRes := Pokemon{}
		err := json.Unmarshal(val, &pokemonRes)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonRes := Pokemon{}
	err = json.Unmarshal(data, &pokemonRes)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return pokemonRes, nil
}