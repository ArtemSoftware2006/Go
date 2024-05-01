/*

Создайте структуру UpperWriter, содержащую поле UpperString string.
Реализуйте интерфейс io.Writer.
Метод Write должен переводить строку в верхний регистр и записывать данные в поле UpperString.
В случае ошибки - верните её.

*/

package task3

import "strings"

type UpperWriter struct {
	UpperString string
}

func (w *UpperWriter) Write(buf []byte) (int, error) {
	str := string(buf)
	w.UpperString = strings.ToUpper(str)
	return len(buf), nil
}
