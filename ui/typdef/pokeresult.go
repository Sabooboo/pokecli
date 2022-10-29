package typdef

import (
	apistructs "github.com/mtslzr/pokeapi-go/structs"
	"image"
)

const (
	Health         = "hp"
	Attack         = "attack"
	SpecialAttack  = "special-attack"
	Defense        = "defense"
	SpecialDefense = "special-defense"
	Speed          = "speed"
)

type PokeResult struct {
	Pokemon   apistructs.Pokemon
	Species   apistructs.PokemonSpecies
	Types     []apistructs.Type
	Abilities []Ability
	Images    ShinyToggleable[PokeImg]
	Stats     Stats[int]
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

type Stats[T any] struct {
	Health         T
	Attack         T
	SpecialAttack  T
	Defense        T
	SpecialDefense T
	Speed          T
}

type PokeImg struct {
	Img image.Image
	Err error
}
