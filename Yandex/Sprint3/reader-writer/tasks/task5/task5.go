package task5

import (
	"bytes"
	"io"
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	var bufSize int16 = 1024
	buf1 := make([]byte, bufSize)
	buf2 := make([]byte, bufSize)
	var err error
	for err != io.EOF {

		_, err = r.Read(buf1)
		if err != nil && err != io.EOF {
			return false, err
		}
		_, err = r.Read(buf2)
		if err != nil && err != io.EOF {
			return false, err
		}
		if bytes.Contains(append(buf1, buf2...), seq) {
			return true, nil
		}

		return false, nil
	}
	return false, io.EOF
}
