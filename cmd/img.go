/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Sabooboo/pokecli/util"
	"github.com/spf13/cobra"
)

// imgCmd represents the img command
var imgCmd = &cobra.Command{
	Use:   "img",
	Short: "Prints the image located at a given URL",
	// Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		// Perform some sort of validation here -- I can't be asked right now in testing.

		fmt.Println(util.URLToASCII(url))
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
