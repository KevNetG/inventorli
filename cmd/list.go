package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"inventorli/inventory"
	"os"
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
	if file == "" {
		return
	}
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	defer f.Close()

	stat, err := os.Stat(file)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	h := inventory.TransactionHistory{}
	h.Read(f, stat.Size())

	inventory, err := inventory.ReproduceFromHistory(h)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	for i, item := range inventory.Items {
		fmt.Println(fmt.Sprintf("[%d] %s", i, item.Name))
	}
}

func init() {
	cmdList.Flags().StringVarP(&file, "file", "f", "", "Box files")

	rootCmd.AddCommand(cmdList)
}
