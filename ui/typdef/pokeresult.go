package typdef

import (
	apistructs "github.com/mtslzr/pokeapi-go/structs"
	"image"
)

type PokeResult struct {
	Pokemon   apistructs.Pokemon
	Species   apistructs.PokemonSpecies
	Types     []apistructs.Type
	Abilities []Ability
	Images    ShinyToggleable[PokeImg]
	Error     error
}

type ShinyToggleable[T any] struct {
	Normal T
	Shiny  T
}

type Ability struct {
	Info     apistructs.Ability
	IsHidden bool
}

type PokeImg struct {
	Img image.Image
	Err error
}
