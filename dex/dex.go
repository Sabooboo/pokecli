package dex

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/mtslzr/pokeapi-go"
)

type ID string

type Pokemon string

func (p Pokemon) FilterValue() string { return string(p) }

type cache struct {
	Dexes []Pokedex `json:"dexes"`
}

type Pokedex struct {
	Id    ID        `json:"id"`
	Names []Pokemon `json:"pokemon"`
}

const (
	fileName = "dex.json"
)

const (
	National ID = "1"
)

func GetPokedex(id ID) (Pokedex, error) {
	var err error
	var pokedex Pokedex

	// Look in fs cache
	pokedex, err = getPokedex(id)

	// If not in fs cache
	if err != nil {
		pokedex, err = FetchPokedex(id) // Todo: handle failure case
	}

	// Update cache
	err = updateCache(pokedex)

	// Return
	return pokedex, err
}

func getCache() (cache, error) {
	file, err := os.Open(fileName)
	if err != nil {
		file, _ = os.Create(fileName)
	}
	bytes, _ := io.ReadAll(file)
	var dexes cache
	json.Unmarshal(bytes, &dexes)
	return dexes, nil
}

func getPokedex(id ID) (Pokedex, error) {

	dexes, err := getCache()
	if err != nil {
		return Pokedex{}, err
	}
	for _, v := range dexes.Dexes {
		if v.Id == id {
			return v, nil
		}
	}
	return Pokedex{}, errors.New("No pokedex found at id.")
}

func FetchPokedex(id ID) (Pokedex, error) {
	pokedex := Pokedex{Id: id, Names: make([]Pokemon, 0)}
	all, err := pokeapi.Pokedex(string(id))
	if err != nil {
		return Pokedex{}, err
	}
	for _, v := range all.PokemonEntries {
		pokedex.Names = append(pokedex.Names, Pokemon(v.PokemonSpecies.Name))
	}
	return pokedex, nil
}

func updateCache(pkmn Pokedex) error {
	// Get cache
	dexCache, err := getCache()
	if err != nil {
		return err
	}

	// Change entry where id = pkmn.id, or create if not exists
	{
		// Look for matching entry
		var i int
		var v Pokedex
		for i, v = range dexCache.Dexes {
			if v.Id == pkmn.Id {
				break
			}
		}

		// If matching entry
		if v.Id == pkmn.Id {
			dexCache.Dexes[i] = pkmn
		} else {
			dexCache.Dexes = append(dexCache.Dexes, pkmn)
		}
	}

	// Serialize cache
	entry, err := json.MarshalIndent(dexCache, "", " ")
	if err != nil {
		return err
	}

	// Write back to file
	err = os.WriteFile(filepath.Join(fileName), entry, 0644)
	return err
}
