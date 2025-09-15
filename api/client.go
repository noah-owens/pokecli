package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchPokemon(name string) (*Pokemon, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	resp, err := http.Get(url)

	// Failure to fetch (likely incorrect argument)
	if err != nil {
		// DID YOU MEAN: get [similar pokemon names]
		return nil, fmt.Errorf("failed to fetch %w", err)
	}

	// Ensure resources are closed when FetchPokemon() returns
	defer resp.Body.Close()

	// HTTP status error
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	// Decoding Error
	var p Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, fmt.Errorf("failed to decode: %w", err)
	}

	return &p, nil
}
