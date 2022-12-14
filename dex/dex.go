package dex

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	e "github.com/Sabooboo/pokecli/dex/errors"

	"github.com/mtslzr/pokeapi-go"
)

type ID int

type Pokemon string

func (p Pokemon) FilterValue() string { return string(p) }

// Pokedex with an ID and list of Pokémon. Many of these may be stored in cache.
type Pokedex struct {
	Id    ID        `json:"id"`
	Names []Pokemon `json:"pokemon"`
}

type cache struct {
	Dexes []Pokedex `json:"dexes"`
}

const (
	FileName = "dex.json"
)

const (
	National ID = 1
)

// GetPokedex Retrieves the Pokedex matching id from the cache or from a webserver if not found.
func GetPokedex(id ID) (Pokedex, error) {
	var err error
	var pokedex Pokedex

	// Look in fs cache.
	// TODO: Store Pokedex in memory on startup and reference those instead of fs cache.
	pokedex, err = GetPokedexFromCache(id, true)

	// If not in fs cache or cache is invalid
	if err != nil || len(pokedex.Names) == 0 {
		// pokeapi-go caches all requests in memory
		// so local cache invalidation should not
		// affect web requests.
		pokedex, _ = FetchPokedex(id) // Todo: handle failure case
	}

	// Update cache
	err = UpdateCache(pokedex)

	// Return
	return pokedex, err
}

// FetchPokedex retrieves the Pokedex matching id from PokeAPI.
// Note that this does not cache requests.
// Use GetPokedex instead if you want to cache
// requests in filesystem for later starts
func FetchPokedex(id ID) (Pokedex, error) {
	pokedex := Pokedex{Id: id, Names: make([]Pokemon, 0)}
	all, err := pokeapi.Pokedex(fmt.Sprint(id))
	if err != nil {
		return Pokedex{}, e.FetchFailed
	}
	for _, v := range all.PokemonEntries {
		pokedex.Names = append(pokedex.Names, Pokemon(v.PokemonSpecies.Name))
	}
	return pokedex, nil
}

// InvalidateCache invalidates the data located in the persistent cache under id.
// If 0 is passed as the id, the whole cache is deleted.
func InvalidateCache(id ID) error {
	var err error
	if id == 0 {
		err = delCache()
		if err != nil {
			return e.FileNotFound
		}
		return nil
	}

	return UpdateCache(Pokedex{
		Id:    id,
		Names: make([]Pokemon, 0),
	})
}

// GetPokedexFromCache Retrieves a Pokedex from the cache, optionally creating if not exists.
func GetPokedexFromCache(id ID, create bool) (Pokedex, error) {
	dexes, err := getCache(create)
	if err != nil {
		return Pokedex{}, err
	}
	for _, v := range dexes.Dexes {
		if v.Id == id {
			return v, nil
		}
	}
	return Pokedex{}, e.NotFound
}

// Retrieves the cache file from disk, optionally creating if not exists.
func getCache(create bool) (cache, error) {
	file, err := os.Open(FileName)
	if err != nil && create {
		file, _ = os.Create(FileName)
	}
	bytes, _ := io.ReadAll(file)
	var dexes cache
	err = json.Unmarshal(bytes, &dexes)
	return dexes, err
}

func delCache() error {
	return os.Remove(FileName)
}

func UpdateCache(pkmn Pokedex) error {
	// Get cache
	dexCache, err := getCache(true)
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
	err = os.WriteFile(filepath.Join(FileName), entry, 0644)
	return err
}
