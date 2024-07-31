package lz77

type Packet struct{
	distance int
	length int
	char string
}

type MatchResult struct {
	matchList []Packet
	actualMatch *Packet
}