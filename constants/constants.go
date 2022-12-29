package constants

import "fmt"

type DbConfig struct {
	host     string
	port     uint
	dbname   string
	user     string
	password string
}

var config = DbConfig{
	host:     "localhost",
	port:     5432,
	dbname:   "poke-api",
	user:     "poke-api",
	password: "poke-api",
}

var Dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Moscow", config.host, config.user, config.password, config.dbname, config.port)

const (
	BaseApiUrl = "https://pokeapi.co/api/v2/pokemon?limit=1154"
)
