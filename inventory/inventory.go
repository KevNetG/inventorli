package inventory

import (
	"errors"
	"fmt"
)

type Box struct {
	Items []Item
}

func New() *Box {
	return &Box{
		[]Item{},
	}
}

func ReproduceFromHistory(history History) (Box, error) {
	inventory := New()

	for _, t := range history.Transactions {
		if t.Amount > 0 {
			for i := 0; i < t.Amount; i++ {
				inventory.Add(t.Item)
			}
		} else {
			err := inventory.Remove(t.Item.Name)
			if err != nil {
				return *inventory, err
			}
		}
	}
	return *inventory, nil
}

func (i *Box) Add(item Item) {
	i.Items = append(i.Items, item)
}

func (i *Box) Remove(name string) error {
	idx := i.Contains(name)
	if idx == -1 {
		return errors.New(fmt.Sprintf("Non existing item %s was removed from the box", name))
	}
	i.Items = append(i.Items[:idx], i.Items[idx+1:]...)

	return nil
}

func (i *Box) Contains(name string) int {
	for idx, item := range i.Items {
		if item.Name == name {
			return idx
		}
	}

	return -1
}

func (i *Box) ContainsItem(item Item) int {
	for idx, it := range i.Items {
		if it.Name == item.Name && it.Description == item.Description {
			return idx
		}
	}

	return -1
}
