package cmd

import (
	"fmt"
	"github.com/Sabooboo/pokecli/util"

	"github.com/spf13/cobra"
)

const (
	shinyFlag  = "shiny"
	shinyFlagP = "s"
)

// imgCmd represents the img command
var imgCmd = &cobra.Command{
	Use:   "img",
	Short: "Prints the image located at a given URL",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		isShiny, err := cmd.Flags().GetBool(shinyFlag)
		if err != nil {
			fmt.Println(err)
			return
		}

		name := args[0]
		img := util.GetImage(name, isShiny)
		if img.Err != nil {
			fmt.Println(img.Err)
			return
		}
		ascii := util.ImageToASCII(img.Img, -1, -1, true)
		fmt.Println(ascii)
	},
}

func init() {
	rootCmd.AddCommand(imgCmd)

	imgCmd.Flags().BoolP(
		shinyFlag,
		shinyFlagP,
		false,
		"Use this to get a shiny image.",
	)
}
