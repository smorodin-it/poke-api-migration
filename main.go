package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"poke-api-go/constants"
	"poke-api-go/models"
	"sync"
)

func getPokemon(url string, pkmEntityList *[]models.PokemonResponse, wg *sync.WaitGroup) {
	defer wg.Done()
	retrievePokemonResp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	retrievePokemonRespBody, err := io.ReadAll(retrievePokemonResp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var pokemonResp models.PokemonResponse
	if err = json.Unmarshal(retrievePokemonRespBody, &pokemonResp); err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("ID: %d Name: %s", pokemonResp.Id, pokemonResp.Name)
	*pkmEntityList = append(*pkmEntityList, pokemonResp)
}

func main() {
	dsn := constants.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&models.Pokemon{}); err != nil {
		log.Fatal(err.Error())
	}

	pokeListResp, err := http.Get(constants.BaseApiUrl)
	if err != nil {
		log.Fatal(err.Error())
	}

	pokeListRespBody, err := io.ReadAll(pokeListResp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var listResponse models.PokemonListResponse
	if err = json.Unmarshal(pokeListRespBody, &listResponse); err != nil {
		log.Fatal(err.Error())
	}

	var pkmEntityList []models.PokemonResponse

	//var pokeList []models.Pokemon

	var wg sync.WaitGroup
	for _, pkm := range listResponse.Results {
		wg.Add(1)
		go getPokemon(pkm.Url, &pkmEntityList, &wg)
	}
	wg.Wait()

	fmt.Println(pkmEntityList)

	//db.Create(&pokeList)
}
