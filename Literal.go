package lz77

type Literal string

func (l Literal) res() string {
	return string(l)
}