package cmd

import (
	"fmt"
	"github.com/Sabooboo/pokecli/dex"
	"github.com/spf13/cobra"
	"strconv"
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
		// Explicitly delete the file in this case if no id or <=0 passed.
		if id <= 0 {
			_ = dex.DelCache()
			fmt.Println("Successfully reset cache")
			return
		}
		// fetch fs cache, save to dex internal cache
		_, err := dex.GetCacheFromDisk(true)
		if err != nil { // Missing/malformed
			fmt.Printf("Trouble reading Pokédex from disk:\n%s\n", err)
			return
		}
		// invalidate Pokedex matching id
		dex.InvalidateCache(id)

		// write back to fs
		err = dex.WriteCache()
		if err != nil { // Error writing
			fmt.Printf("Trouble writing Pokédex at id %d:\n%s\n", id, err)
			return
		}

		fmt.Printf("Successfully reset cache %s\n", strconv.Itoa(id))
	},
}

func init() {
	resetCmd.Flags().Int(cacheFlag, 0, "Reset the Pokedex stored on the local device. 0 resets the whole cache.")
	rootCmd.AddCommand(resetCmd)
}
