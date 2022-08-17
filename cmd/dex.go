/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
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
		var pkdx dex.Pokedex
		var err error
		id, cacheOnly, override := getFlags(cmd)
		if id < 1 {
			cmd.Help()
			return
		}

		if cacheOnly {
			pkdx, err = dex.GetPokedexFromCache(dex.ID(id), false)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			pkdx, err = dex.GetPokedex(dex.ID(id), override)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		printPokedex(pkdx)
	},
}

func printPokedex(pkdx dex.Pokedex) {
	s := strings.Builder{}
	for _, v := range pkdx.Names {
		s.WriteString(fmt.Sprintf("%s\n", v))
	}

	fmt.Print(s.String())
}

func getFlags(cmd *cobra.Command) (dex.ID, bool, bool) {
	// id, default 1
	id, _ := cmd.PersistentFlags().GetInt(IdFlag)

	// Only use cache, default no
	cacheOnly := false
	cacheOnly, _ = cmd.Flags().GetBool(cacheOnlyFlag)

	// override cache, default no
	override := false
	override, _ = cmd.Flags().GetBool(overrideFlag)

	return dex.ID(id), cacheOnly, override
}

func init() {
	rootCmd.AddCommand(dexCmd)

	dexCmd.PersistentFlags().IntP(
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
