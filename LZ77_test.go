package lz77

import (
	"log"
	"testing"
)

func TestLz(t *testing.T) {
	test := "tres tristes tigres tragaban trigo en un trigal"
	c := Compress(test, 6)
	log.Println("Compression finished")
	log.Println(c)
	res := Decompress(c)

	if res != test {
		t.Fatal("Compression failed")
	}
}