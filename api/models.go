package api

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

func PokemonToSummary(p *Pokemon) PokemonSummary {
	stats := make(map[string]int)
	for _, s := range p.Stats {
		stats[s.Stat.Name] = s.BaseStat
	}

	var types []string
	for _, t := range p.Types {
		types = append(types, t.Type.Name)
	}

	var abilities []string
	for _, a := range p.Abilities {
		abilities = append(abilities, a.Ability.Name)
	}

	var moves []string
	for _, m := range p.Moves {
		moves = append(moves, m.Move.Name)
	}

	var heldItems []string
	for _, item := range p.HeldItems {
		heldItems = append(heldItems, item.Item.Name)
	}

	return PokemonSummary{
		Name:      p.Name,
		ID:        p.ID,
		Height:    p.Height,
		Weight:    p.Weight,
		Types:     types,
		Abilities: abilities,
		Stats:     stats,
		Species:   p.Species.Name,
		Moves:     moves,
		HeldItems: heldItems,
	}
}
