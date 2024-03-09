package main

import (
	"fmt"
	"time"
)

func TimeDifference(start, end time.Time) time.Duration {
	return end.Sub(start)
}

func main() {
	fmt.Println(TimeDifference(time.Now(), time.Now().Add(time.Minute*5)))
}
