package main

func DoubleNumbers(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case <-done:
				return
			default:
				out <- num * 2
			}
		}
	}()
	return out
}
