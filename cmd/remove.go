package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"inventorli/inventory"
	"strconv"
	"time"
)

var cmdRemove = &cobra.Command{
	Use: "remove [string to echo]",
	Short: `Removes an item with a specific id from a box.
		The id can be found when listing the items using the inventorli list command`,
	Long: `echo is for echoing anything back.
    EAddedcho echo’s.
    `,
	Run: runRemove,
}

func runRemove(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Printf("you must provide an index, which item to remove from an inventory list")
		return
	}

	idx, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	h := inventory.History{}
	h.ReadFile(file)
	inv, err := inventory.ReproduceFromHistory(h)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	if len(inv.Items) <= idx {
		fmt.Printf("The item %d does not exist int the inventory list", idx)
		return
	}

	h.Transactions = append(h.Transactions, inventory.Transaction{
		time.Now().Format("2006/01/02"),
		reason,
		inv.Items[idx],
		-amount,
	})

	err = h.WriteFile(file)
	if err != nil {
		panic(err)
	}
}

func init() {
	cmdRemove.Flags().StringVarP(&file,
		"file",
		"f",
		"",
		"path to a history file",
	)
	cmdRemove.Flags().StringVarP(&reason,
		"reason",
		"r",
		"",
		"Reason what the item is or was used for",
	)
	cmdRemove.Flags().IntVarP(&amount,
		"amount",
		"n",
		1,
		"How many Items shall be removed from the box",
	)

	rootCmd.AddCommand(cmdRemove)
}
