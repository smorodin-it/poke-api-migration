package models

// Pokemon list

type PokemonListResponse struct {
	Count    uint
	Next     *string
	Previous *string
	Results  []PokemonListResponseResultModel
}

type PokemonListResponseResultModel struct {
	Name string
	Url  string
}

// Pokemon response

type PokemonResponse struct {
	Id    uint
	Name  string
	Types []PokemonResponseTypeModel
}

type PokemonResponseTypeModel struct {
	Slot uint
	Type struct {
		Name string
		Url  string
	}
}
