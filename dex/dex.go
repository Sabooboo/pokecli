package dex

import (
	"encoding/json"
	"fmt"
	e "github.com/Sabooboo/pokecli/dex/errors"
	"github.com/mtslzr/pokeapi-go"
	"io"
	"os"
)

type Pokemon string

func (p Pokemon) FilterValue() string { return string(p) }

// Pokedex with an ID and list of Pokémon. Many of these may be stored in cache.
type Pokedex struct {
	Id    int       `json:"id"`
	Names []Pokemon `json:"pokemon"`
}

type Location uint8

type Cache struct {
	Dexes []Pokedex `json:"dexes"`
}

const (
	FileName = "dex.json"
)

const (
	National int = 1
)

const (
	Memory Location = 0
	Disk   Location = 1
	Web    Location = 2
)

var localCache = Cache{}

// GetPokedex Retrieves the Pokedex matching id. This result is cached in
// memory.
func GetPokedex(id int) (Pokedex, error) {
	var (
		pkmn Pokedex
		err  error
	)
	// Check if cache exists in ram
	pkmn, err = GetPokedexFrom(Memory, id, true)
	if err == nil {
		return pkmn, nil
	}

	// Check if cache exists on disk
	pkmn, err = GetPokedexFrom(Disk, id, true)
	if err == nil {
		return pkmn, nil
	}

	// Fetch from internet
	pkmn, err = GetPokedexFrom(Web, id, true)
	if err == nil {
		return pkmn, nil
	}

	// Failure case: Pokedex could not be found anywhere.
	return Pokedex{}, err
}

// GetPokedexFrom retrieves the Pokedex matching id from the Location provided,
// optionally caching the result. If there is an error retrieving the Pokedex,
// the result will not be cached, and the error will be returned.
func GetPokedexFrom(location Location, id int, cacheRes bool) (Pokedex, error) {
	var (
		pkmn Pokedex
		err  error
	)
	switch location {
	case Disk:
		pkmn, err = getPokedexFromDisk(id)
	case Memory:
		pkmn, err = getPokedexFromMemory(id)
	case Web:
		pkmn, err = getPokedexFromPokeAPI(id)
	}
	if err != nil {
		return pkmn, err
	}

	if cacheRes {
		updateCache(pkmn)
	}
	return pkmn, nil
}

// InvalidateCache invalidates the data located in the persistent cache under id.
// If 0 or lower is passed as the id, the whole cache is reset.
func InvalidateCache(id int) {
	if id <= 0 {
		localCache = Cache{}
		return
	}

	updateCache(Pokedex{
		Id:    id,
		Names: make([]Pokemon, 0),
	})
}

// Following 4 are a set of non-caching, request-only functions

// getPokedexFromMemory returns the Pokedex matching id in the local Cache.
// If the matching Pokedex could not be found, an error will be returned.
func getPokedexFromMemory(id int) (Pokedex, error) {
	for _, v := range localCache.Dexes {
		if v.Id == id {
			return v, nil
		}
	}
	return Pokedex{}, e.IdNotFound
}

// getPokedexFromDisk returns the Pokedex matching id in the file matching
// FileName. If the file is missing or malformed, or the Pokedex can not be
// found, an error will be returned.
func getPokedexFromDisk(id int) (Pokedex, error) {
	dexes, err := GetCacheFromDisk(false)
	if err != nil {
		return Pokedex{}, err
	}
	for _, v := range dexes.Dexes {
		if v.Id == id {
			return v, nil
		}
	}
	return Pokedex{}, e.IdNotFound
}

// getPokedexFromPokeAPI returns the Pokedex matching id from the PokéAPI.
// If the request does not succeed, an error will be returned.
func getPokedexFromPokeAPI(id int) (Pokedex, error) {
	pokedex := Pokedex{Id: id, Names: make([]Pokemon, 0)}
	all, err := pokeapi.Pokedex(fmt.Sprint(id))
	if err != nil {
		return Pokedex{}, err
	}
	for _, v := range all.PokemonEntries {
		pokedex.Names = append(pokedex.Names, Pokemon(v.PokemonSpecies.Name))
	}
	return pokedex, nil
}

// GetCacheFromDisk Retrieves the cache located in the file matching FileName.
// If the file is missing or malformed, an error will be returned, and the
// result will not be cached.
func GetCacheFromDisk(cacheRes bool) (Cache, error) {
	file, pathErr := os.Open(FileName)
	if pathErr != nil {
		return Cache{}, pathErr
	}
	bytes, _ := io.ReadAll(file)
	var dexes Cache
	err := json.Unmarshal(bytes, &dexes)
	if err == nil && cacheRes {
		localCache = dexes
	}
	return dexes, err
}

// WriteCache serializes localCache and writes it to the file matching FileName,
// creating it if not exists.
func WriteCache() error {
	entry, err := json.MarshalIndent(localCache, "", " ")
	if err != nil {
		return err
	}

	// Write back to file
	err = os.WriteFile(FileName, entry, 0666)
	return err
}

// DelCache wraps os.Remove, targeting the file matching FileName.
func DelCache() error {
	return os.Remove(FileName)
}

// updateCache edits localCache, changing the Pokedex matching pkmn.Id to pkmn,
// or appending pkmn if not exists.
func updateCache(pkmn Pokedex) {
	for i, v := range localCache.Dexes {
		if v.Id == pkmn.Id {
			localCache.Dexes[i] = v
			return
		}
	}
	localCache.Dexes = append(localCache.Dexes, pkmn)
}
