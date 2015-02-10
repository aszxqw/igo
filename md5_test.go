package igo

import "testing"

func TestMd5(t *testing.T) {
	if GetMd5String("1234") != "81dc9bdb52d04dc20036dbd8313ed055" {
		t.Fatal("failed.")
	}
}
