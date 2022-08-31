package util

import (
	"image"
	"image/png"
	"net/http"
	"strings"

	"github.com/Sabooboo/pokecli/ui/typdef"
	"github.com/mtslzr/pokeapi-go"
	apistructs "github.com/mtslzr/pokeapi-go/structs"
	imgascii "github.com/qeesung/image2ascii/convert"
)

func GetPokemon(id string, out chan<- typdef.PokeResult) {
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

	// TODO Fetch official artwork front-default dynamically. Need to find where that is if it is in the wrapper at all.
	imgUrl := pkmn.Sprites.FrontDefault // "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/112.png"
	img, _ := URLToImage(imgUrl)
	// TODO Error handling, or maybe not since img is nil in this case and it therefore will not display.

	out <- typdef.PokeResult{
		Pokemon:   pkmn,
		Species:   species,
		Types:     types,
		Abilities: abilities,
		Image:     img,
		Error:     leastNil(errA, errB), // Ensure that if there was any error, nil will not be returned.
	}
}

// Fetches the image located at a URL and returns an ASCII representation.
func URLToASCII(url string) string {
	img, err := URLToImage(url)
	if err != nil {
		return ""
	}
	return ImageToASCII(img, &imgascii.DefaultOptions)
}

func URLToImage(url string) (image.Image, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	img, err := png.Decode(res.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func ImageToASCII(img image.Image, options *imgascii.Options) string {
	convert := imgascii.NewImageConverter()
	return convert.Image2ASCIIString(img, options)
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

// Returns a substring of str starting at a 0-based index and ending
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
