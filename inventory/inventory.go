package inventory

import (
	"errors"
	"fmt"
)

type Inventory struct {
	Items []Item
}

func New() *Inventory {
	return &Inventory{
		[]Item{},
	}
}

func ReproduceFromHistory(history History) (Inventory, error) {
	inventory := New()

	for _, t := range history.Transactions {
		if t.Amount > 0 {
			inventory.Add(t.Item)
		} else {
			err := inventory.Remove(t.Item.Name)
			if err != nil {
				return *inventory, err
			}
		}
	}
	return *inventory, nil
}

func (i *Inventory) Add(item Item) {
	i.Items = append(i.Items, item)
}

func (i *Inventory) Remove(name string) error {
	idx := i.Contains(name)
	if idx == -1 {
		return errors.New(fmt.Sprintf("Non existing item %s was removed from the box", name))
	}
	i.Items = append(i.Items[:idx], i.Items[idx+1:]...)

	return nil
}

func (i *Inventory) Contains(name string) int {
	for idx, item := range i.Items {
		if item.Name == name {
			return idx
		}
	}

	return -1
}
