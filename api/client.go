// Package api defines data models for PokeAPI responses and provides functions for fetching and transforming Pokémon data.
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchPokemon retrieves data for a given Pokémon name from the PokeAPI.
// On success, it returns a pointer to a populated Pokemon struct and a nil error.
// On failure, it returns nil and a descriptive error.
func FetchPokemon(name string) (*Pokemon, error) {

	// Create the url path to get a pokemon with name name.
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)

	// HTTP Get Request to url, return the response and any error encountered
	resp, err := http.Get(url)

	// Failure to fetch error (likely incorrect argument)
	if err != nil {
		// TODO -- DID YOU MEAN: get [similar pokemon names]
		return nil, fmt.Errorf("failed to fetch %w", err)
	}

	// Ensure resources are closed when FetchPokemon() returns
	defer resp.Body.Close()

	// HTTP status error
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	// Init zero-value Pokemon struct
	var p Pokemon

	// Decode response body into the zero-value Pokemon struct pointed to by &p.
	// Return wrapped error if decoding fails
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, fmt.Errorf("failed to decode: %w", err)
	}

	return &p, nil
}
