package igo

import "testing"

func TestLineIterator(t *testing.T) {
	iter, err := NewLineIterator("testdata/lines")
	if err != nil {
		t.Fatal(err)
	}
	lines := [...]string{"12", "345", "67"}
	i := 0
	for iter.HasNext() {
		line := iter.Next()
		if line != lines[i] {
			t.Fatalf("%s %s", line, lines[i])
		}
		i++
	}
}
