package cmd

import (
	"fmt"

	"github.com/Sabooboo/pokecli/dex"
	e "github.com/Sabooboo/pokecli/dex/errors"
	"github.com/spf13/cobra"
)

const (
	cacheFlag = "cache"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset all data associated with pokecli",
	Long: `reset is used to delete pokecli data. If you are having issues with
pokecli, resetting the persistent storage may help.`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.PersistentFlags().GetInt(cacheFlag)
		if id > -1 {
			err := dex.InvalidateCache(dex.ID(id))
			if err != nil && err != e.FileNotFound {
				fmt.Println(err)
				return
			}
			s := fmt.Sprint(id)
			if id == 0 {
				s = ""
			}
			fmt.Printf("Succesfully reset cache %s\n", s)
		}

	},
}

func init() {
	resetCmd.PersistentFlags().Int(cacheFlag, 0, "Reset the Pokedex stored on the local device. 0 resets the whole cache.")
	rootCmd.AddCommand(resetCmd)
}
