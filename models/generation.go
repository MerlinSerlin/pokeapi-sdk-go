package models

// type Generation struct {
// 	ID        int    `json:"id"`
// 	Name      string `json:"name"`
// 	Abilities []struct {
// 		Name string `json:"name"`
// 		Url  string `json:"url"`
// 	} `json:"abilities"`
// 	MainRegion struct {
// 		Name string `json:"name"`
// 		Url  string `json:"url"`
// 	} `json:"main_region"`
// 	Moves []struct {
// 		Name string `json:"name"`
// 		Url  string `json:"url"`
// 	} `json:"moves"`
// 	PokemonSpecies []struct {
// 		Name string `json:"name"`
// 		Url  string `json:"url"`
// 	} `json:"pokemon_species"`
// 	Types []struct {
// 		Name string `json:"name"`
// 		Url  string `json:"url"`
// 	} `json:"types"`
// 	VersionGroups []struct {
// 		Name string `json:"name"`
// 		Url  string `json:"url"`
// 	} `json:"version_groups"`
// }

type Generation struct {
	ID             int                `json:"id"`
	Name           string             `json:"name"`
	Abilities      []Ability          `json:"abilities"`
	MainRegion     NamedAPIResource   `json:"main_region"`
	Moves          []NamedAPIResource `json:"moves"`
	Names          []Name             `json:"names"`
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
	Types          []NamedAPIResource `json:"types"`
	VersionGroups  []NamedAPIResource `json:"version_groups"`
}

type Ability struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Name struct {
	Language NamedAPIResource `json:"language"`
}
