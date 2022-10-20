/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Sabooboo/pokecli/ui/typdef"
	"github.com/Sabooboo/pokecli/util"
	"github.com/spf13/cobra"
)

// imgCmd represents the img command
var imgCmd = &cobra.Command{
	Use:   "img",
	Short: "Prints the image located at a given URL",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		in := make(chan typdef.PokeResult)
		go util.GetPokemon(name, in) // TODO Make lightweight util.GetImage to ease API calls
		mon := <-in
		// TODO add support for --shiny flag
		img := util.ImageToASCII(mon.Images.Normal.Img, -1, -1, true)
		fmt.Println(img)
	},
}

func init() {
	rootCmd.AddCommand(imgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
