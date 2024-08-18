package lz77

type packet struct{
	distance int
	length int
	char string
}

func (p packet)ToRes() Result {
	if p.char != "" {
		return Literal(p.char)
	} else {
		pair := Pair{distance: p.distance, length: p.length} 
		return pair
	}
}
