package main

import "time"

type Transaction struct {
	Date            time.Time `json:"date"`
	Reason          string 	  `json:"reason"`
	ItemName        string    `json:"itemName"`
	ItemDescription string    `json:"itemDescription"`
	Amount          int       `json:"amount"`
}
