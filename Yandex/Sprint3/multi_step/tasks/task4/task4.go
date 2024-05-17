package main

func Tee[T any](done <-chan struct{}, in <-chan T) (<-chan T, <-chan T) {
	out1 := make(chan T)
	out2 := make(chan T)

	go func() {
		defer close(out1)
		defer close(out2)

		for val := range in {
			select {
			case <-done:
				return
			default:
				out1 <- val
				out2 <- val
			}
		}
	}()

	return out1, out2
}
