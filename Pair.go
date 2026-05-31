package lz77

import "fmt"

type Pair struct {
	distance int
	length   int
}

func (p Pair) res() string {
	return p.toString()
}

func (p Pair) Byte() []byte {
	res := make([]byte, 0)
	res = append(res, byte('['))
	res = append(res, byte(p.distance))
	res = append(res, byte(p.length))
	res = append(res, byte(']'))
	return res
}

func (p Pair) toString() string {
	return fmt.Sprintf("[ distance=%d, length=%d ]", p.distance, p.length)
}

func PairFromBytes(b []byte) Pair {
	if len(b) != 4 || (b[0] != byte('[') && b[3] != byte(']')) {
		panic("Not a pair")
	}

	return Pair{int(b[1]), int(b[2])}
}
