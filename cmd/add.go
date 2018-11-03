package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"inventorli/inventory"
	"time"
)

var description string

var cmdAdd = &cobra.Command{
	Use:   "add [string to echo]",
	Short: "add the items inside a box",
	Long: `echo is for echoing anything back.
    Echo echo’s.
    `,
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Printf("you must provide exactly one item name")
		return
	}

	if file == "" {
		fmt.Printf("you must provide at least one file")
		return
	}

	h := inventory.History{}
	h.ReadFile(file)
	h.Transactions = append(h.Transactions, inventory.Transaction{
		time.Now().Format("2006/01/02"),
		reason,
		inventory.Item{
			args[0],
			description,
		},
		amount,
	})

	err := h.WriteFile(file)
	if err != nil {
		panic(err)
	}
}

func init() {
	cmdAdd.Flags().StringVarP(&file,
		"file",
		"f",
		"",
		"Box file",
	)
	cmdAdd.Flags().StringVarP(&reason,
		"reason",
		"r",
		"",
		"Reason what the item is or was used for",
	)
	cmdAdd.Flags().StringVarP(&description,
		"description",
		"d",
		"",
		"Additional information to the item",
	)
	cmdAdd.Flags().IntVarP(&amount,
		"amount",
		"n",
		1,
		"How many Items shall be added to the box",
	)

	rootCmd.AddCommand(cmdAdd)
}
