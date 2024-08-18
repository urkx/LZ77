package lz77

import (
	"fmt"
	"testing"

	"github.com/urkx/lgo"
)

func TestLz(t *testing.T) {
	test := "tres tristes tigres tragaban trigo en un trigal"
	c := Compress(test, 32000)
	fmt.Print(lgo.Info("Compression finished"))
	fmt.Print(lgo.Debug(c))
	res := Decompress(c)

	if res != test {
		t.Fatal("Compression failed")
	}
}