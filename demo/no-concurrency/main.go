package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type NewsAggPage struct {
	Title    string
	Pokemons map[string]string
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func newsRoutine(result chan Response, Location string) {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	result <- responseObject
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	pokemons := make(map[string]string)
	Locations := [...]string{
		"https://pokeapi.co/api/v2/pokedex/kanto/",
		"https://pokeapi.co/api/v2/pokedex/national/",
		"https://pokeapi.co/api/v2/pokedex/original-johto/",
		"https://pokeapi.co/api/v2/pokedex/hoenn/",
		"https://pokeapi.co/api/v2/pokedex/original-sinnoh/",
		"https://pokeapi.co/api/v2/pokedex/original-unova/",
		"https://pokeapi.co/api/v2/pokedex/kanto/",
		"https://pokeapi.co/api/v2/pokedex/national/",
		"https://pokeapi.co/api/v2/pokedex/original-johto/",
		"https://pokeapi.co/api/v2/pokedex/hoenn/",
		"https://pokeapi.co/api/v2/pokedex/original-sinnoh/",
		"https://pokeapi.co/api/v2/pokedex/original-unova/",
		"https://pokeapi.co/api/v2/pokedex/kanto/",
		"https://pokeapi.co/api/v2/pokedex/national/",
		"https://pokeapi.co/api/v2/pokedex/original-johto/",
		"https://pokeapi.co/api/v2/pokedex/hoenn/",
		"https://pokeapi.co/api/v2/pokedex/original-sinnoh/",
		"https://pokeapi.co/api/v2/pokedex/original-unova/",
		"https://pokeapi.co/api/v2/pokedex/kanto/",
		"https://pokeapi.co/api/v2/pokedex/national/",
		"https://pokeapi.co/api/v2/pokedex/original-johto/",
		"https://pokeapi.co/api/v2/pokedex/hoenn/",
		"https://pokeapi.co/api/v2/pokedex/original-sinnoh/",
		"https://pokeapi.co/api/v2/pokedex/original-unova/",
		"https://pokeapi.co/api/v2/pokedex/kanto/",
		"https://pokeapi.co/api/v2/pokedex/national/",
		"https://pokeapi.co/api/v2/pokedex/original-johto/",
		"https://pokeapi.co/api/v2/pokedex/hoenn/",
		"https://pokeapi.co/api/v2/pokedex/original-sinnoh/",
		"https://pokeapi.co/api/v2/pokedex/original-unova/",
	}
	for _, Location := range Locations {
		response, err := http.Get(Location)
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		var responseObject Response
		json.Unmarshal(responseData, &responseObject)
		for i := 0; i < len(responseObject.Pokemon); i++ {
			pokemons[responseObject.Pokemon[i].Species.Name] = responseObject.Pokemon[i].Species.Name
		}
	}
	p := NewsAggPage{Title: "Pokemon Non-Use Concurrency", Pokemons: pokemons}

	t, err := template.ParseFiles("main.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/pokemon/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}
