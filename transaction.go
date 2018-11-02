package main

import "io"
import "encoding/json"

type Transaction struct {
	Date            string `json:"date"`
	Reason          string `json:"reason"`
	ItemName        string `json:"itemName"`
	ItemDescription string `json:"itemDescription"`
	Amount          int    `json:"amount"`
}

func (t *Transaction) Write(w io.Writer) {
	b, _ := json.Marshal(*t)
	w.Write(b)
}
