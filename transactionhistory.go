package inventory

import (
    "io"
    "encoding/json"
    "fmt"
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
    fmt.Print(string(b))
    json.Unmarshal(b, t)
}
