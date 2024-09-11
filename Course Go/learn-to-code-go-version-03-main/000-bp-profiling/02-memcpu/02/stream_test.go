// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// go test -run none -bench . -benchtime 3s -benchmem -memprofile p.out
// go test -run none -bench . -benchtime 3s -benchmem -memprofile p.out -gcflags -m=2
// go test -run none -bench . -benchtime 3s -benchmem -cpuprofile p.out
// go test -bench . -benchmem -memprofile p.out -gcflags -m=2

// OLD WAY: go tool pprof memcpu.test p.out
// go tool pprof -noinlines p.out
// ---then---
// list algOne
// ---and/or---
// weblist algOne


// Tests to see how each algorithm compare.
package main

import (
	"bytes"
	"testing"
)

var output bytes.Buffer
var in = assembleInputStream()
var find = []byte("elvis")
var repl = []byte("Elvis")

// Capture the time it takes to execute algorithm one.
func BenchmarkAlgorithmOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output.Reset()
		algOne(in, find, repl, &output)
	}
}

// Capture the time it takes to execute algorithm two.
func BenchmarkAlgorithmTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output.Reset()
		algTwo(in, find, repl, &output)
	}
}
