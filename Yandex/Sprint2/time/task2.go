package main

import (
	"fmt"
	"time"
)

func FormatTimeToString(timestamp time.Time, format string) string {
	return timestamp.Format(format)
}

func main() {
	fmt.Println(FormatTimeToString(time.Now(), time.RFC822))
}
