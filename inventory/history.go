package inventory

import (
	"encoding/json"
	"io"
	"os"
)

type History struct {
	Transactions []Transaction `json:"transactions"`
}

func (h *History) Write(w io.Writer) error {
	b, err := json.Marshal(*h)
	if err != nil {
		return err
	}
	_, err = w.Write(b)

	return err
}

// Truncates the file before writing
func (h *History) WriteFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Set cursor to write at the beginning of the file
	f.Truncate(0)
	f.Seek(0, 0)

	err = h.Write(f)

	return err
}

func (h *History) Read(r io.Reader, n int64) {
	b := make([]byte, n)
	r.Read(b)
	json.Unmarshal(b, h)
}

func (h *History) ReadFile(path string) error {
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	stat, err := os.Stat(path)
	if err != nil {
		return err
	}

	h.Read(f, stat.Size())

	return nil
}
