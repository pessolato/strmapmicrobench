package main

import (
	"testing"
)

const DNA string = `TAAGGTATCCCAAGTTGGCTCCTGACTTAGCTGCCTCGGGAGGAACTTAGGGCAACTCCAGTAAAGTTACTTTTGGCGCGGGGTATGTTATGTTCGATAGTAAGGTATCCCAAGTTGGCTCCTGACTTAGCTGCCTCGGGAGGAACTTAGGGCAACTCCAGTAAAGTTACTTTTGGCGCGGGGTATGTTATGTTCGATAGTAAGGTATCCCAAGTTGGCTCCTGACTTAGCTGCCTCGGGAGGAACTTAGGGCAACTCCAGTAAAGTTACTTTTGGCGCGGGGTATGTTATGTTCGATAGTAAGGTATCCCAAGTTGGCTCCTGACTTAGCTGCCTCGGGAGGAACTTAGGGCAACTCCAGTAAAGTTACTTTTGGCGCGGGGTATGTTATGTTCGATAG`

func BenchmarkCountNucleotidesByByteIndex(b *testing.B) {
	for b.Loop() {
		CountNucleotidesByByteIndex(DNA)
	}
}

func BenchmarkCountNucleotidesByByteIndexAsRune(b *testing.B) {
	for b.Loop() {
		CountNucleotidesByByteIndexAsRune(DNA)
	}
}

func BenchmarkCountNucleotidesByByteIndexAsInt16(b *testing.B) {
	for b.Loop() {
		CountNucleotidesByByteIndexAsInt16(DNA)
	}
}

func BenchmarkCountNucleotidesByRuneRange(b *testing.B) {
	for b.Loop() {
		CountNucleotidesByRuneRange(DNA)
	}
}

func BenchmarkCountNucleotidesArrayBytes(b *testing.B) {
	for b.Loop() {
		CountNucleotidesArrayBytes(DNA)
	}
}

func BenchmarkCountNucleotidesArrayRunes(b *testing.B) {
	for b.Loop() {
		CountNucleotidesArrayRunes(DNA)
	}
}

func BenchmarkTranscribeDnaToRnaBytes(b *testing.B) {
	rna := make([]byte, len(DNA))
	for b.Loop() {
		TranscribeDnaToRnaBytes(DNA, rna)
	}
}

func BenchmarkTranscribeDnaToRnaRunes(b *testing.B) {
	rna := make([]rune, len(DNA))
	for b.Loop() {
		TranscribeDnaToRnaRunes(DNA, rna)
	}
}
