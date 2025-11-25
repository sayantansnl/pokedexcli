package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sayantansnl/pokedexcli/internal/pokecache"
)

type Client struct {
	cache   *pokecache.Cache
	baseUrl string
}

func NewClient(baseUrl string, interval time.Duration) *Client {
	return &Client{
		cache: pokecache.NewCache(interval),
		baseUrl: baseUrl,
	}
}

func (c *Client) FetchLocationAreas(url string) (LocationList, error) {
	if data, ok := c.cache.Get(url); ok {
		var locations LocationList
		if err := json.Unmarshal(data, &locations); err != nil {
			return LocationList{}, fmt.Errorf("unreadable: %w", err)
		}
		return locations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationList{}, fmt.Errorf("couldn't fetch data due to error: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationList{}, fmt.Errorf("couldn't read data: %w", err)
	}

	c.cache.Add(url, body)

	var locations LocationList
	if err := json.Unmarshal(body, &locations); err != nil {
		return LocationList{}, fmt.Errorf("unreadable: %w", err)
	}

	return locations, nil
}