package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"inventorli/inventory"
	"os"
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
	f, err := os.OpenFile(file, os.O_RDWR, 0644)
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

	h := inventory.History{}
	h.Read(f, stat.Size())
	h.Transactions = append(h.Transactions, inventory.Transaction{
		time.Now().Format("2006/01/02"),
		reason,
		inventory.Item{
			args[0],
			description,
		},
		amount,
	})

	fmt.Printf("%s", len(h.Transactions))
	f.Truncate(0)
	f.Seek(0, 0)

	err = h.Write(f)
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
	cmdRemove.Flags().StringVarP(&reason,
		"reason",
		"r",
		"",
		"Reason what the item is or was used for",
	)
	cmdRemove.Flags().StringVarP(&description,
		"description",
		"d",
		"",
		"Additional information to the item",
	)
	cmdRemove.Flags().IntVarP(&amount,
		"amount",
		"n",
		1,
		"How many Items shall be added to the box",
	)

	rootCmd.AddCommand(cmdAdd)
}
