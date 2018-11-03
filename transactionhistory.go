package main

import (
	"io"
	"encoding/json"
)

type TransactionHistory struct {
	Transactions []Transaction `json:"transactions"`
}

func (t *TransactionHistory) Write(w io.Writer) {
	b, _ := json.Marshal(*t)
	w.Write(b)
}

func (t *TransactionHistory) Read(r io.Reader, n int64) {
	b := make([]byte, n)
	r.Read(b)
	json.Unmarshal(b, t)
}
