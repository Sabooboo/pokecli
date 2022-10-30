package cmd

import (
	"github.com/Sabooboo/pokecli/ui"
	"github.com/spf13/cobra"
)

// uiCmd represents the ui command
var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Start the pokecli user interface",
	Long:  `This command turns your terminal into a working pokedex where you can look up data on Pokemon.`,
	Run: func(cmd *cobra.Command, args []string) {
		ui.Start()
	},
}

func init() { rootCmd.AddCommand(uiCmd) }
