package days

import (
	"log/slog"
	"testing"
)

func TestWithExampleInput(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`
	content := []byte(input)
	result1, result2, err := Process(&content)

	if err != nil {
		t.Fatal(err)
	}

	expected := "41"
	if result1 != expected {
		t.Error("expected result1 to be " + expected + " but was " + result1)
	}

	expected = "6"
	if result2 != expected {
		t.Error("expected result2 to be " + expected + " but was " + result2)
	}
}
