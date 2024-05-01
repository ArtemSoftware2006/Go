package task1

import (
	"io"
)

func WriteString(s string, w io.Writer) error {
	_, err := w.Write([]byte(s))
	if err != nil {
		return err
	}
	return nil
}
