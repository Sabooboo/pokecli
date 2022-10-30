package util

import (
	"image"
	"image/png"
	"io"
	"math"
	"net/http"
	"strings"
	"sync"

	"github.com/Sabooboo/pokecli/ui/typdef"
	"github.com/mtslzr/pokeapi-go"
	apistructs "github.com/mtslzr/pokeapi-go/structs"
	imgascii "github.com/qeesung/image2ascii/convert"
)

func GetPokemon(id string, out chan<- typdef.PokeResult) {
	id = strings.ToLower(id)
	pkmn, errA := pokeapi.Pokemon(id)
	species, errB := pokeapi.PokemonSpecies(id)

	types := make([]apistructs.Type, 0, 2)
	for _, v := range pkmn.Types {
		res, err := pokeapi.Type(v.Type.Name)
		if err == nil {
			types = append(types, res)
		}
	}

	abilities := make([]typdef.Ability, 0)
	for _, v := range pkmn.Abilities {
		res, err := pokeapi.Ability(v.Ability.Name)
		if err != nil {
			continue
		}
		ability := typdef.Ability{
			Info:     res,
			IsHidden: v.IsHidden,
		}
		abilities = append(abilities, ability)
	}

	stats := make(map[string]int)
	for _, v := range pkmn.Stats {
		stats[v.Stat.Name] = v.BaseStat
	}

	resultStats := typdef.Stats[int]{
		Health:         stats[typdef.Health],
		Attack:         stats[typdef.Attack],
		SpecialAttack:  stats[typdef.SpecialAttack],
		Defense:        stats[typdef.Defense],
		SpecialDefense: stats[typdef.SpecialDefense],
		Speed:          stats[typdef.Speed],
	}

	result := typdef.PokeResult{
		Pokemon:   pkmn,
		Species:   species,
		Types:     types,
		Abilities: abilities,
		Stats:     resultStats,
		Error:     leastNil(errA, errB), // Ensure that if there was any error, nil will not be returned.
	}

	imgUrl := pkmn.Sprites.FrontDefault
	shinyUrl := pkmn.Sprites.FrontShiny

	var wg sync.WaitGroup

	// Opted not to use channels since the request is multithreaded but still encapsulated here.
	getImage := func(url string, isShiny bool) {
		defer wg.Done()
		img, err := URLToImage(url)

		out := &result.Images.Normal
		if isShiny {
			out = &result.Images.Shiny
		}
		*out = typdef.PokeImg{
			Img: img,
			Err: err,
		}
	}

	wg.Add(2)

	go getImage(imgUrl, false)
	go getImage(shinyUrl, true)

	wg.Wait()

	out <- result
}

func GetImage(id string, isShiny bool) typdef.PokeImg {
	pkmn, errA := pokeapi.Pokemon(id)
	url := pkmn.Sprites.FrontDefault
	if isShiny {
		url = pkmn.Sprites.FrontShiny
	}
	img, errB := URLToImage(url)
	return typdef.PokeImg{
		Img: img,
		Err: leastNil(errA, errB),
	}
}

func URLToImage(url string) (image.Image, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	img, err := png.Decode(res.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func ImageToASCII(img image.Image, width, height int, fit bool) string {
	convert := imgascii.NewImageConverter()
	opt := imgascii.Options{
		Ratio:           1,
		FixedWidth:      width,
		FixedHeight:     height,
		FitScreen:       fit,
		StretchedScreen: false,
		Colored:         true,
		Reversed:        false,
	}
	return convert.Image2ASCIIString(img, &opt)
}

func leastNil(errs ...error) error {
	for _, v := range errs {
		if v != nil {
			return v
		}
	}
	return nil
}

func Title(str string) string {
	s := strings.Fields(str)
	for i, v := range s {
		s[i] = strings.ToUpper(Substring(v, 0, 1)) + strings.ToLower(Substring(v, 1, len(v)))
	}
	return strings.Join(s, " ")
}

// Substring returns a substring of str starting at a 0-based index and ending
// after a certain amount of characters. Negative numbers do not work.
func Substring(str string, start, length int) string {
	whole := []rune(str)

	if start >= len(whole) {
		return ""
	}

	if start+length > len(whole) {
		length = len(whole) - start
	}

	return string(whole[start : start+length])
}

func Min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func Max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func Reverse[T any](a []T) []T {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
