package task1

func StringsGen(lines ...string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, str := range lines {
			out <- str
		}
	}()
	return out
}
