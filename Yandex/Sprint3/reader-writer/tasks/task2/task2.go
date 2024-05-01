package task2

import (
	"io"
)

func ReadString(r io.Reader) (string, error) {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
