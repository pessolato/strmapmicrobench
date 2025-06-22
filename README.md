# Go Strings & Maps Microbenchmarks

This project provides a set of microbenchmarks to evaluate the performance of Go's built-in maps and various methods for iterating over strings, specifically for counting nucleotides and transcribing DNA to RNA. The benchmarks compare different approaches using byte and rune indexing, as well as array-based counting.

## Purpose

The goal is to measure and compare the efficiency of:

- Using Go maps with different key types (`byte`, `rune`, `int16`)
- Iterating over strings by byte index vs. rune range
- Counting nucleotides using arrays vs. maps
- Transcribing DNA to RNA using byte and rune slices

This is a microbenchmark-only project and does not provide any production-ready functionality.

## Benchmarked Functions

The following functions are benchmarked (see [main.go](main.go)):

- [`CountNucleotidesByByteIndex`](main.go)
- [`CountNucleotidesByByteIndexAsRune`](main.go)
- [`CountNucleotidesByByteIndexAsInt16`](main.go)
- [`CountNucleotidesByRuneRange`](main.go)
- [`CountNucleotidesArrayBytes`](main.go)
- [`CountNucleotidesArrayRunes`](main.go)
- [`TranscribeDnaToRnaBytes`](main.go)
- [`TranscribeDnaToRnaRunes`](main.go)

See [main_test.go](main_test.go) for the benchmark definitions.

## Running the Benchmarks

To run all benchmarks with CPU profiling:

```sh
go test -bench=. -benchtime=1000000x -benchmem -cpuprofile=cpu.out
```

To generate a graph image of the CPU profile:

```sh
go tool pprof -png cpu.out
```

## Conclusions

Based on the microbenchmarks, the following conclusions can be drawn:

1. **Byte vs. Rune Iteration:** Iterating over strings by bytes is slightly faster than iterating by runes, provided the strings contain only ASCII characters. This is because byte iteration avoids the overhead of decoding runes.

2. **Map Key Types and Performance:** Maps with `byte` keys use the generic `internal/runtime/maps.(*Map).putSlotSmall` method, which is significantly slower than the specialized `internal/runtime/maps.(*Map).putSlotSmallFast32` method used for maps with `int32` keys (such as `rune`). As a result, using `rune` as the map key can yield better performance than using `byte`.

3. **Array-Based Counting:** When the set of possible keys is small and known in advance, using an array with direct indexing (for example, via a `switch` statement) is much faster than using a map. This approach eliminates hashing and lookup overhead.

4. **Type Conversion for Maps:** If a map must be used, converting a small value (like a `byte`) to a larger type (such as a `rune`) can significantly improve performance due to the more efficient map implementation for larger key types.

These findings highlight the importance of choosing the right data structures and iteration methods for optimal performance in Go, especially when processing large strings or performing