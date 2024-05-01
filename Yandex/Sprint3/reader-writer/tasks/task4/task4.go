package task4

import (
	"io"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
	var buf = make([]byte, n)

	fact_n, err := r.Read(buf)
	if err != nil {
		return err
	}

	if fact_n != int(n) {
		buf = make([]byte, fact_n)
	}
	_, err = w.Write(buf)
	if err != nil {
		return err
	}

	return nil
}
