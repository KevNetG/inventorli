package main

type Transaction struct {
	Date            string `json:"date"`
	Reason          string `json:"reason"`
	ItemName        string `json:"itemName"`
	ItemDescription string `json:"itemDescription"`
	Amount          int    `json:"amount"`
}
