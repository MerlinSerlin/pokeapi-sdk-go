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

### Caching

API calls are cached by default to mitgate number of requests to PokeAPI.




