package main

func CountNucleotidesByByteIndex(dna string) map[byte]int {
	chars := make(map[byte]int)
	for i := range len(dna) {
		chars[dna[i]]++
	}
	return chars
}

func CountNucleotidesByByteIndexAsRune(dna string) map[rune]int {
	chars := make(map[rune]int)
	for i := range len(dna) {
		chars[rune(dna[i])]++
	}
	return chars
}

func CountNucleotidesByByteIndexAsInt16(dna string) map[int16]int {
	chars := make(map[int16]int)
	for i := range len(dna) {
		chars[int16(dna[i])]++
	}
	return chars
}

func CountNucleotidesByRuneRange(dna string) map[rune]int {
	chars := make(map[rune]int)
	for _, r := range dna {
		chars[r]++
	}
	return chars
}

func CountNucleotidesArrayBytes(dna string) map[byte]int {
	chars := [4]int{0, 0, 0, 0}
	for i := range len(dna) {
		switch dna[i] {
		case 'A':
			chars[0]++
		case 'C':
			chars[1]++
		case 'G':
			chars[2]++
		case 'T':
			chars[3]++
		}
	}
	return map[byte]int{
		'A': chars[0],
		'C': chars[1],
		'G': chars[2],
		'T': chars[3],
	}
}

func CountNucleotidesArrayRunes(dna string) map[rune]int {
	chars := [4]int{0, 0, 0, 0}
	for i := range dna {
		switch dna[i] {
		case 'A':
			chars[0]++
		case 'C':
			chars[1]++
		case 'G':
			chars[2]++
		case 'T':
			chars[3]++
		}
	}
	return map[rune]int{
		'A': chars[0],
		'C': chars[1],
		'G': chars[2],
		'T': chars[3],
	}
}

func TranscribeDnaToRnaBytes(dna string, rna []byte) {
	for i := range len(dna) {
		switch dna[i] {
		case 'A':
			rna[i] = 'U'
		case 'C':
			rna[i] = 'G'
		case 'G':
			rna[i] = 'C'
		case 'T':
			rna[i] = 'A'
		}
	}
}

func TranscribeDnaToRnaRunes(dna string, rna []rune) {
	for i, n := range dna {
		switch n {
		case 'A':
			rna[i] = 'U'
		case 'C':
			rna[i] = 'G'
		case 'G':
			rna[i] = 'C'
		case 'T':
			rna[i] = 'A'
		}
	}
}

func main() {
	panic("Please run benchmarks instead.")
}
