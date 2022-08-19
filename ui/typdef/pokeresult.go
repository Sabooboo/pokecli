package typdef

import apistructs "github.com/mtslzr/pokeapi-go/structs"

type PokeResult struct {
	Pokemon   apistructs.Pokemon
	Species   apistructs.PokemonSpecies
	Types     []apistructs.Type
	Abilities []Ability
	Error     error
}

type Ability struct {
	Info     apistructs.Ability
	IsHidden bool
}
