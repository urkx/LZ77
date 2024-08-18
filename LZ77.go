package lz77

import (
	"slices"
)

func searchLongestMatch(data string, searchBuffer string, sp int, lp int, wp int) matchResult{
	matchList := []packet{}
	var actualMatch *packet = nil
	m := ""
	matching := false
	x := sp
	for x < len(searchBuffer) {
		if searchBuffer[x] == data[lp] {
			m = m + string(searchBuffer[x])
			if actualMatch == nil {
				actualMatch = &packet{distance: wp - x, length: len(m), char: ""}
			} else {
				actualMatch.length = len(m)
			}
			lp += 1
			if lp >= len(data) {
				lp = len(data) - 1
			}
			if !matching {
				matching = true
			}
			x += 1
		} else if matching {
			lp = wp
			matchList = append(matchList, *actualMatch)
			m = ""
			matching = false
			actualMatch = nil
		} else {
			x += 1
		}
	}
	return matchResult{matchList: matchList, actualMatch: actualMatch}
}

func orderMatchList(x packet, y packet) int {
	if x.length > y.length { 
		return 1 
	} else if x.length < y.length { 
		return -1 
	} else { return 0 }
}

func Compress(data string, winSize int) []Result {
	buffer_output := []packet{}

	wp := 0
	
	for wp < len(data) {
		sp := 0
		lp := wp
		if wp > winSize { sp = wp - winSize }
		search_buffer := data[sp:wp]
		matchResult := searchLongestMatch(data, search_buffer, sp, lp, wp)
		dwp := 1
		if matchResult.actualMatch == nil && len(matchResult.matchList) == 0 {
			matchResult.matchList = append(matchResult.matchList, packet{distance: 0, length: 0, char: string(data[wp])})
		} else if matchResult.actualMatch != nil {
			matchResult.matchList = append(matchResult.matchList, *matchResult.actualMatch)
			dwp = matchResult.actualMatch.length
		}
		slices.SortFunc(matchResult.matchList, orderMatchList)
		newPacket := matchResult.matchList[len(matchResult.matchList)-1]
		buffer_output = append(buffer_output, newPacket)
		if newPacket.length > 0 {
			dwp = newPacket.length
		}
		wp += dwp
	}
	output := []Result{}
	i := 0
	for i < len(buffer_output) {
		output = append(output, buffer_output[i].ToRes())
		i += 1
	}
	return output
}

func Decompress(input []Result) string {
	output := ""
	x := 0
	for x < len(input) {
		packet := input[x]
		switch t := packet.(type) {
			case Pair:
				d := t.distance
				l := t.length
				actualPointer := len(output)
				newData := output[actualPointer - d : (actualPointer - d) + l]
				output += newData
			default:
				output += packet.res()
		}
		x += 1
	}
	return output
}