package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"inventorli/inventory"
	"math"
)

var cmdList = &cobra.Command{
	Use:   "list [string to echo]",
	Short: "list the items inside a box",
	Long: `echo is for echoing anything back.
    Echo echo’s.
    `,
	Run: listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	h := inventory.History{}
	h.ReadFile(file)

	inv, err := inventory.ReproduceFromHistory(h)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	for i, item := range inv.Items {
		fmt.Println(fmt.Sprintf("[%d] %s", i, item.Name))
		fmt.Printf("   ")

		for j := 0; j < int(math.Log10(float64(i+1))+1); j++ {
			fmt.Print(" ")
		}

		fmt.Printf(fmt.Sprintf("%s", item.Description))
		fmt.Println()
	}
}

func init() {
	cmdList.Flags().StringVarP(&file, "file", "f", "", "Box files")

	rootCmd.AddCommand(cmdList)
}
