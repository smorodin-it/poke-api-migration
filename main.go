package main

import (
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"poke-api-go/constants"
	"poke-api-go/models"
)

func getPokemon(url string) {
	pokeResp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}

	pokeRespBody, err := io.ReadAll(pokeResp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var entityResponse models.PokemonResponse
	if err = json.Unmarshal(pokeRespBody, &entityResponse); err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("ID: %d Name: %s", entityResponse.Id, entityResponse.Name)

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
		log.Println(err.Error())
	}

	pokeListRespBody, err := io.ReadAll(pokeListResp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var listResponse models.PokemonListResponse
	if err = json.Unmarshal(pokeListRespBody, &listResponse); err != nil {
		log.Fatal(err.Error())
	}

	//var pokeList []models.Pokemon
	for _, pkm := range listResponse.Results {
		getPokemon(pkm.Url)
	}

	//db.Create(&pokeList)

}
