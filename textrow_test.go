package textbackend

import (
	"testing"
)

// tests

func Test_TextRow_NewTextRow(t *testing.T) {
	row := NewTextRow(1, "hello world")
	if row.GetLevel() != 1 {
		t.Error("TextRow.GetLevel() not equal")
	}
	if string(row.GetText()) != "hello world" {
		t.Error("TextRow.GetText() not equal")
	}
}

func Test_TextRow_GetLevelWhitespace(t *testing.T) {
	textRow := TextRow{}
	textRow.Level = 1
	if textRow.GetLevelWhitespace("\t") != "\t" {
		t.Error("GetLevelWhitespace failed")
	}
}

func Test_TextRow_GetString(t *testing.T) {
	textRow := NewTextRow(0, "hello world")
	if textRow.GetString("\t") != "hello world" {
		t.Error("GetString failed")
	}

	textRow2 := NewTextRow(1, "hello world")
	if textRow2.GetString("\t") != "\thello world" {
		t.Error("GetString failed")
	}
}

// benchmarks

var resultTextRow string

func benchmark_TextRow(c string, b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		row := NewTextRow(2, c)
		r = row.GetString("<br>")
	}
	resultTextRow = r
}

func Benchmark_TextRow_small(b *testing.B) { benchmark_TextRow("hello world", b) }

func Benchmark_TextRow_middle(b *testing.B) { benchmark_TextRow("hello world, this is just a test", b) }

func Benchmark_TextRow_large(b *testing.B) {
	benchmark_TextRow("Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", b)
}
