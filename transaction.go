package main

type Transaction struct {
	Date   string `json:"date"`
	Reason string `json:"reason"`
	Item   Item   `json:"item"`
	Amount int    `json:"amount"`
}
