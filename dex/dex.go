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

// Pokemon implements list.Item.
func (p Pokemon) FilterValue() string { return string(p) }

// Pokedex with an ID and list of pokemon. Many of these may be stored in cache.
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

// Retrieves the pokedex matching id from the cache or from a webserver if not found.
func GetPokedex(id ID, updateCache bool) (Pokedex, error) {
	var err error
	var pokedex Pokedex

	// Look in fs cache.
	// TODO: Store pokedexes in memory on startup and reference those instead of fs cache.
	pokedex, err = GetPokedexFromCache(id, updateCache)

	// If not in fs cache or cache is invalid
	if err != nil || len(pokedex.Names) == 0 {
		// pokeapi-go caches all requests in memory
		// so local cache invalidation should not
		// affect web requests.
		pokedex, _ = FetchPokedex(id) // Todo: handle failure case
	}

	// Update cache
	if updateCache {
		UpdateCache(pokedex)
	}

	// Return
	return pokedex, nil
}

// Retrieves the pokedex matching id from PokeAPI.
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

// Invalidates the data located in the persistant cache under id.
// If 0 is passed as the id, the whole cache is deleted.
func InvalidateCache(id ID) error {
	var err error
	if id == 0 {
		err = delCache()
		if err != nil {
			return e.FileNotFound
		}
	}

	return UpdateCache(Pokedex{
		Id:    id,
		Names: make([]Pokemon, 0),
	})
}

// Retrieves the pokedex from cache if exists. Optionally creates if not exists.
func GetPokedexFromCache(id ID, createCache bool) (Pokedex, error) {
	dexes, err := getCache(createCache)
	if err != nil {
		return Pokedex{}, err
	}
	for _, v := range dexes.Dexes {
		if v.Id == id {
			return v, nil
		}
	}
	return Pokedex{}, e.DexNotFound
}

// Retrieves the cache file from disk. Optionally creates if not exists.
func getCache(createCache bool) (cache, error) {
	file, err := os.Open(FileName)
	var dexes cache
	if err != nil {
		if createCache {
			file, _ = os.Create(FileName)
		} else {
			return dexes, err
		}
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		return dexes, err
	}
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
