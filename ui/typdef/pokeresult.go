package typdef

import (
	"image"

	apistructs "github.com/mtslzr/pokeapi-go/structs"
)

type PokeResult struct {
	Pokemon   apistructs.Pokemon
	Species   apistructs.PokemonSpecies
	Types     []apistructs.Type
	Abilities []Ability
	Image     image.Image
	Error     error
}

type Ability struct {
	Info     apistructs.Ability
	IsHidden bool
}
