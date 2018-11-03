package inventory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateListFromHistory(t *testing.T) {
	t1 := Transaction{
		"2018/10/12",
		"Initial",
		Item{
			"HDMI cable",
			"",
		},
		1,
	}

	t2 := Transaction{
		"2018/10/13",
		"Initial",
		Item{
			"SATA cable",
			"",
		},
		1,
	}

	t3 := Transaction{
		"2018/10/14",
		"Initial",
		Item{
			"HDMI cable",
			"",
		},
		-1,
	}

	h := TransactionHistory{[]Transaction{t1, t2, t3}}

	l, err := ReproduceFromHistory(h)
	if err != nil {
		panic(err)
	}

	assert.Contains(t, l.Items, t2.Item, "The list should only contain a SATA cable")
}

func TestCreateListFromHistoryAndRemoveNonExistingItem(t *testing.T) {
	t1 := Transaction{
		"2018/10/12",
		"Initial",
		Item{
			"HDMI cable",
			"",
		},
		1,
	}

	t2 := Transaction{
		"2018/10/13",
		"Initial",
		Item{
			"SATA cable",
			"",
		},
		1,
	}

	t3 := Transaction{
		"2018/10/14",
		"Initial",
		Item{
			"CD Player",
			"",
		},
		-1,
	}

	h := TransactionHistory{[]Transaction{t1, t2, t3}}

	_, err := ReproduceFromHistory(h)

	assert.Error(t, err, "Should throw error. It is not possible to remove non existing items from a box")
}

func TestInventory_Add(t *testing.T) {
	inventory := New()
	item := Item{
		"SATA cable",
		"",
	}

	assert.NotContains(t, inventory.Items, item)

	inventory.Add(item)

	assert.Contains(t, inventory.Items, item)
}

func TestInventory_Remove(t *testing.T) {
	inventory := New()
	item := Item{
		"SATA cable",
		"",
	}
	inventory.Add(item)

	assert.Contains(t, inventory.Items, item)
	err := inventory.Remove(item.Name)
	assert.NoError(t, err, "Should not throw an error")
	assert.NotContains(t, inventory.Items, item)
}
