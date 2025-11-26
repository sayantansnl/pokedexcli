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

func (c *Client) FetchLocationAreaDetails(name string) (LocationDetails, error) {
    url := c.baseUrl + "/location-area/" + name

    if data, ok := c.cache.Get(url); ok {
        var location LocationDetails
        if err := json.Unmarshal(data, &location); err != nil {
            return LocationDetails{}, fmt.Errorf("unable to get details: %w", err)
        }
        return location, nil
    }

    res, err := http.Get(url)
    if err != nil {
        return LocationDetails{}, fmt.Errorf("unable to fetch details: %w", err)
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return LocationDetails{}, fmt.Errorf("unreadable location: %w", err)
    }

    c.cache.Add(url, body)

    var location LocationDetails
    if err := json.Unmarshal(body, &location); err != nil {
        return LocationDetails{}, fmt.Errorf("unreadable location: %w", err)
    }

    return location, nil
}

func (c *Client) FetchPokemon(pokemonName string) (PokemonDetails, error) {
	url := c.baseUrl + "/pokemon/" + pokemonName

	if data, ok := c.cache.Get(url); ok {
        var pokemon PokemonDetails
        if err := json.Unmarshal(data, &pokemon); err != nil {
            return PokemonDetails{}, fmt.Errorf("unable to get details: %w", err)
        }
        return pokemon, nil
    }

	res, err := http.Get(url)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("cannot fetch pokemon details: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("unreadable pokemon details: %w", err)
	}

	c.cache.Add(url, body)

	var pokemon PokemonDetails
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return PokemonDetails{}, fmt.Errorf("unable to unmarshal: %w", err)
	}

	return pokemon, nil
}