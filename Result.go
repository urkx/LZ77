package lz77

type Result interface {
	res() string
	Byte() []byte
}
