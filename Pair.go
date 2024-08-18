package lz77

import "fmt"

type Pair struct {
	distance int
	length int
}

func (p Pair) res() string {
	return p.toString()
}


func (p Pair)toString() string {
	return fmt.Sprintf("[ distannce=%d, length=%d ]", p.distance, p.length)
}