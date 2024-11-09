package pokeapi

import (
	"net/http"
	"time"

	"github.com/itsMe-ThatOneGuy/go-pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		cache: pokecache.NewCache(),
		httpClient: http.Client{
			Timeout: 5 * time.Second,
		},
	}
}
