package models

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
