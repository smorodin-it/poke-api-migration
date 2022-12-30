package api_responses

// List Response Model
type ListResponseModel struct {
	Count    uint
	Next     *string
	Previous *string
}

// Pokemon list

type PokemonListResponse struct {
	ListResponseModel
	Results []PokemonListResponseResultModel
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

// Pokemon Types response

type PokemonTypesResponse struct {
	ListResponseModel
	Result []PokemonTypeResponseResultModel
}

type PokemonTypeResponseResultModel struct {
	Name string
	Url  string
}
