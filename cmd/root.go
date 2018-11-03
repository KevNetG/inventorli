package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "invemtory",
	Short: "Inventory is a tool for managing and tracking your stuff",
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
