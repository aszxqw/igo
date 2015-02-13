package igo

import "testing"

func TestWait(t *testing.T) {
	wg := WaitGroup{}
	wg.Go(func() {
		i := 1
		i++
	})
	wg.Wait()
}
