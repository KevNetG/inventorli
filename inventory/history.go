package inventory

import (
	"encoding/json"
	"io"
)

type History struct {
	Transactions []Transaction `json:"transactions"`
}

func (t *History) Write(w io.Writer) error {
	b, err := json.Marshal(*t)
	if err != nil {
		return err
	}
	_, err = w.Write(b)

	return err
}

func (t *History) Read(r io.Reader, n int64) {
	b := make([]byte, n)
	r.Read(b)
	json.Unmarshal(b, t)
}
