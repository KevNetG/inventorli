package inventory

import (
	"encoding/json"
	"io"
	"os"
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

func (t *History) ReadFile(path string) error {
	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	stat, err := os.Stat(path)
	if err != nil {
		return err
	}

	h := History{}
	h.Read(f, stat.Size())

	return nil
}
