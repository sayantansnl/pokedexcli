package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchLocationAreas(baseUrl string) (LocationList, error) {
	//baseUrl := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=0&limit=20")

	res, err := http.Get(baseUrl)
	if err != nil {
		return LocationList{}, fmt.Errorf("couldn't fetch data due to error: %v", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationList{}, fmt.Errorf("couldn't read data: %w", err)
	}

	defer res.Body.Close()

	var locations LocationList
	if err := json.Unmarshal(body, &locations); err != nil {
		return LocationList{}, fmt.Errorf("unreadable: %w", err)
	}

	return locations, nil
}