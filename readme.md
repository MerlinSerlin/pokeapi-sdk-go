# PokeAPI SDK for Go

Welcome to the PokeAPI SDK for Go! This library allows you to easily interact with the PokeAPI to fetch Pokémon data and other related information.

## Features

- **Retrieve Pokémon Data**: Fetch details about Pokémon, including their ID, name, abilities, moves, and more.
- **Access Generation Information**: Get details about Pokémon generations, including main regions, abilities, moves, and species.


Full API documentation can be found at [Poke API](https://pokeapi.co/docs/v2.html).

## Installation

To use the PokeAPI SDK for Go, you need to import the package into your Go project. You can do this by running:

```bash
go get github.com/merlinserlin/pokeapi-sdk-go
```

```go
import "github.com/merlinserlin/pokeapi-sdk-go"
```

## Endpoints
  
#### Get Pokemon

```go
  client := api.NewClient()

// Get a Pokémon
pokemon, err := client.GetPokemon("pokemon")
```

*Must pass an ID (e.g. "1") or name (e.g. "pikachu").*

#### Get Generation


```go
  client := api.NewClient()

// Get a Generation
pokemon, err := client.GetGeneration("generation")
```

*Must pass an ID (e.g. "1") or name (e.g. "Generation I").*

## Tests 

Tests are included in the tests folder. 

It should be noted that the data from PokeAPI seems subject to change - I found that the example response property order for clefairy (id 35) is 64 when I ping the API, not 56 as  outlined in the sample response in the API doc's. 

Considering this information, the focus of tests is primarily fetching data against fields outlined in the models.

## Caching

API calls are cached by default to mitgate number of requests to PokeAPI.

## Dependencies of Note

#### resty
[https://github.com/go-resty/resty/v2][resty]

This project uses resty as an HTTP client to simplify the experience of making HTTP requests. 

#### cache-go
[https://github.com/patrickmn/go-cache][go-cache]

This package is used to simplify caching operations. It's used to cache the requests we're sending to PokeAPI to keep their server costs down.

#### testify
[https://github.com/stretchr/testify][github]

Our tests leverage the assert methods exposed by the testify package. Testify is a pretty standard tool for testing and has good readability. 



