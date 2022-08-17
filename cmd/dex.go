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

// dexCmd represents the dex command
var dexCmd = &cobra.Command{
	Use:   "dex",
	Short: "Retrieve all Pokemon from the national or any regional dex.",
	Long: `This command will retrieve a list of newline-separated Pokemon names from
the Pokedex matching the specified id.`,
	Run: func(cmd *cobra.Command, args []string) {
		// id, default 1
		id, err := cmd.PersistentFlags().GetInt("id")
		if id < 1 || err != nil {
			cmd.Help()
			return
		}
		fmt.Println(id)

		// Only use cache, default no
		cacheOnly := false
		cacheOnly, _ = cmd.Flags().GetBool("cache-only")
		fmt.Println(cacheOnly)

		// override cache, default no
		override := false
		override, _ = cmd.Flags().GetBool("override")
		fmt.Println(override)

		var pkdx dex.Pokedex

		if !override {
			pkdx, err = dex.GetPokedex(dex.ID(id))
			if err != nil {
				fmt.Println(err)
				return
			}
			printPokedex(pkdx)
			return
		}

		pkdx, err = dex.GetPokedexFromCache(dex.ID(id))
		if err != nil || !cacheOnly {
			pkdx, err = dex.FetchPokedex(dex.ID(id))
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		printPokedex(pkdx)
		fmt.Println("Override", override)
		fmt.Println("Cache Only", cacheOnly)
	},
}

func printPokedex(pkdx dex.Pokedex) {
	s := strings.Builder{}
	for _, v := range pkdx.Names {
		s.WriteString(fmt.Sprintf("%s\n", v))
	}

	fmt.Println(s.String())
}

func init() {
	rootCmd.AddCommand(dexCmd)

	dexCmd.PersistentFlags().IntP(
		"id",
		"i",
		1,
		"The id of the pokedex you want to retrieve. 1 retrieves the national dex. Other numbers correspond to generations.",
	)

	dexCmd.Flags().Bool(
		"cache-only",
		false,
		"Define whether this should only retrieve Pokemon from locally stored pokedex.",
	)

	dexCmd.Flags().BoolP(
		"override",
		"o",
		false,
		"Define whether this search the web for the specified dex and override any cached values.",
	)
}
