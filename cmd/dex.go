package cmd

import (
	"fmt"
	"strings"

	"github.com/Sabooboo/pokecli/dex"
	"github.com/spf13/cobra"
)

const (
	IdFlag        = "id"
	cacheOnlyFlag = "cache-only"
	overrideFlag  = "override"
	overrideFlagP = "o"
)

// dexCmd represents the dex command
var dexCmd = &cobra.Command{
	Use:   "dex",
	Short: "Retrieve all Pokemon from the national or any regional dex.",
	Long: `This command will retrieve a list of newline-separated Pokemon names from
the Pokedex matching the specified id.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, cacheOnly, override := getFlags(cmd)
		// set fetch location
		location := dex.Web
		if cacheOnly {
			location = dex.Disk
		}
		// if override is true, cache result to dex internal
		pkmn, err := dex.GetPokedexFrom(location, id, override)
		if err != nil {
			fmt.Printf("Trouble reading Pokédex:\n%s\n", err)
			return
		}
		// if override is true, result was cached, so write to disk
		if override {
			if err = dex.WriteCache(); err != nil {
				fmt.Printf("Trouble writing Pokédex at id %d:\n%s\n", id, err)
				return
			}
		}
		// output
		printPokedex(pkmn)
	},
}

func printPokedex(pkmn dex.Pokedex) {
	s := strings.Builder{}
	for _, v := range pkmn.Names {
		s.WriteString(fmt.Sprintf("%s\n", v))
	}
	fmt.Println(s.String())
}

func getFlags(cmd *cobra.Command) (int, bool, bool) {
	// id, default 1
	id, _ := cmd.PersistentFlags().GetInt(IdFlag)

	// Only use cache, default no
	cacheOnly := false
	cacheOnly, _ = cmd.Flags().GetBool(cacheOnlyFlag)

	// override cache, default no
	override := false
	override, _ = cmd.Flags().GetBool(overrideFlag)

	return id, cacheOnly, override
}

func init() {
	rootCmd.AddCommand(dexCmd)

	dexCmd.Flags().IntP(
		IdFlag,
		"i",
		1,
		"The id of the pokedex you want to retrieve. 1 retrieves the national dex. Other numbers correspond to generations.",
	)

	dexCmd.Flags().Bool(
		cacheOnlyFlag,
		false,
		"Define whether this should only retrieve Pokemon from locally stored pokedex.",
	)

	dexCmd.Flags().BoolP(
		overrideFlag,
		overrideFlagP,
		false,
		"Define whether this search the web for the specified dex and override any cached values.",
	)
}
