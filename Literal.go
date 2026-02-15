package lz77

type Literal string

func (l Literal) res() string {
	return string(l)
}

func (l Literal) Byte() []byte {
	return []byte(l.res())
}
