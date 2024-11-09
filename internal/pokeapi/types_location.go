package pokeapi

type Locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"Previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}
