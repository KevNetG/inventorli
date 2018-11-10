package cmd

import (
	"fmt"
	"github.com/agnivade/levenshtein"
	"github.com/spf13/cobra"
	"inventorli/inventory"
	"io/ioutil"
	"math"
	"path"
	"sort"
)

var cmdSearch = &cobra.Command{
	Use:   "search [string to echo]",
	Short: "list the items inside Search box",
	Long: `echo is for echoing anything back.
    Echo echo’s.
    `,
	Run: searchRun,
}

func searchRun(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		return
	}
	boxes := []inventory.Box{}

	if directory != "" {
		d, err := ioutil.ReadDir(directory)
		if err != nil {
			fmt.Printf("%s", err)
		}

		for _, f := range d {
			inv := createInventory(path.Join(directory, f.Name()))
			boxes = append(boxes, inv)
		}
	}

	if file != "" {
		inv := createInventory(file)
		boxes = append(boxes, inv)
	}

	type Search struct {
		Item     inventory.Item
		Distance int
	}

	ss := []Search{}

	for _, box := range boxes {
		for _, item := range box.Items {
			ss = append(ss, Search{
				item,
				levenshtein.ComputeDistance(args[0], item.Name),
			})
		}
	}
	sort.SliceStable(ss, func(i, j int) bool {
		return ss[i].Distance < ss[j].Distance
	})

	for i, a := range ss {
		fmt.Printf("[%d] %s", i, a.Item.Name)
		fmt.Println()

		for j := 0; j < int(math.Log10(float64(i+1))+1); j++ {
			fmt.Print(" ")
		}

		fmt.Printf(fmt.Sprintf("%s", a.Item.Description))
		fmt.Println()
	}
}

func createInventory(path string) inventory.Box {
	h := inventory.History{}
	h.ReadFile(path)
	inv, err := inventory.ReproduceFromHistory(h)
	if err != nil {
		panic(err)
	}

	return inv
}

func init() {
	cmdSearch.Flags().StringVarP(
		&file,
		"file",
		"f",
		"",
		"Box file",
	)
	cmdSearch.Flags().StringVarP(
		&directory,
		"directory",
		"D",
		"",
		"Path to directory with history files",
	)
	cmdSearch.Flags().StringVarP(&description,
		"description",
		"d",
		"",
		"Additional information to the Item",
	)

	rootCmd.AddCommand(cmdSearch)
}
