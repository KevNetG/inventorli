package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var file string
var reason string
var amount int
var directory string

var rootCmd = &cobra.Command{
	Use:   "inventorli",
	Short: "inventorli is Search tool for managing and tracking your stuff",
	Long:  `TODO: Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("lol")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
