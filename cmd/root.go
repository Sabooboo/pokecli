package cmd

import (
	"fmt"
	"os"

	"github.com/Sabooboo/pokecli/ui"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pokecli",
	Short: "A terminal-based Pokedex app",
	Long: `pokecli is a pokedex located entirely within the terminal.
In pokecli, you can look up information about any Pokemon via
PokeAPI. Just run the app from your terminal and use the built-in
interface to navigate the pokedex! Run with the flag -h for help.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Start()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func init() {}
