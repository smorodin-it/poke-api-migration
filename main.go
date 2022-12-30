package main

import (
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	apiResp "poke-api-go/api_responses"
	"poke-api-go/constants"
	"poke-api-go/models"
	"sync"
)

func getPokemon(url string, pokeList *[]models.Pokemon, wg *sync.WaitGroup) {
	defer wg.Done()
	retrievePokemonResp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	retrievePokemonRespBody, err := io.ReadAll(retrievePokemonResp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var pokemonResp apiResp.PokemonResponse
	if err = json.Unmarshal(retrievePokemonRespBody, &pokemonResp); err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("ID: %d Name: %s", pokemonResp.Id, pokemonResp.Name)

	pokemon := models.Pokemon{
		Name:   pokemonResp.Name,
		PkdxId: pokemonResp.Id,
	}

	if err = pokemon.Types.Set(pokemonResp.Types); err != nil {
		log.Fatal(err.Error())
	}

	*pokeList = append(*pokeList, pokemon)
}

func main() {
	dsn := constants.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// Pokemon list
	if err = db.Migrator().DropTable(&models.Pokemon{}); err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&models.Pokemon{}); err != nil {
		log.Fatal(err.Error())
	}

	pokeListResp, err := http.Get(constants.PokemonListUrl)
	if err != nil {
		log.Fatal(err.Error())
	}

	pokeListRespBody, err := io.ReadAll(pokeListResp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var listResponse apiResp.PokemonListResponse
	if err = json.Unmarshal(pokeListRespBody, &listResponse); err != nil {
		log.Fatal(err.Error())
	}

	var pokeList []models.Pokemon

	var wg sync.WaitGroup
	for _, pkm := range listResponse.Results {
		wg.Add(1)
		go getPokemon(pkm.Url, &pokeList, &wg)
	}

	wg.Wait()
	db.Create(&pokeList)

}
