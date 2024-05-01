package task5

/*
	Напишите функцию Contains(r io.Reader, seq []byte) (bool, error)
	которая должна найти первое вхождение байт seq в данных, доступных через Reader r.
	Если последовательность найдена - верните true, nil, иначе false, nil.
	В случае возникновения ошибки - false и возникшую ошибку.
	**/

import (
	"io"
)

func slicesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Contains(r io.Reader, seq []byte) (bool, error) {
	var buf = make([]byte, len(seq))

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return false, err
		}

		if err == io.EOF {
			break
		}

		if slicesEqual(buf[:n], seq) {
			return true, nil
		}
	}

	return false, nil
}
