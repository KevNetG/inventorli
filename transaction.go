package main

import "time"

type Transaction struct {
	Date            time.Time `json:"date"`
	Reason          string
	ItemName        string
	ItemDescription string
	Amount          int
}
