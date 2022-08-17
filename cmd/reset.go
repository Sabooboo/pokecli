/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Sabooboo/pokecli/dex"
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset all data associated with pokecli",
	Long: `reset is used to delete pokecli data. If you are having issues with
pokecli, resetting the persistent storage may help.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.PersistentFlags().GetInt("cache")
		if err != nil {
			id = 0
		}
		err = dex.InvalidateCache(dex.ID(id))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Reset succesfully")
		}
	},
}

func init() {
	resetCmd.PersistentFlags().Int("cache", 0, "Reset the Pokedex stored on the local device. 0 resets the whole cache.")
	rootCmd.AddCommand(resetCmd)
}
