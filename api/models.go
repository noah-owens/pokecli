// Package api defines data models for PokeAPI responses and display-friendly output
package api

// PokemonSummary contains display-ready fields for rendering a concise Pokémon summary in CLI output.
type PokemonSummary struct {
	Name      string
	ID        int
	Height    int
	Weight    int
	Types     []string
	Abilities []string
	Stats     map[string]int
	Species   string
	Moves     []string
	HeldItems []string
}

// Pokemon maps the JSON structure returned by https://pokeapi.co/api/v2/pokemon/<name>.
// Pokemon includes all fields necessary for decoding API response.
type Pokemon struct {
	Name   string `json:"name"`
	ID     int    `json:"id"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Order  int    `json:"order"`

	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`

	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`

	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`

	Moves []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			Order        any `json:"order"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`

	PastAbilities []struct {
		Abilities []struct {
			Ability  any  `json:"ability"`
			IsHidden bool `json:"is_hidden"`
			Slot     int  `json:"slot"`
		} `json:"abilities"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"past_abilities"`

	PastTypes []any `json:"past_types"`

	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`

	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`

	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}
