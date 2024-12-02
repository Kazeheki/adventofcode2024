package days

import (
	"log/slog"
	"testing"
)

func Test(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	content := []byte(input)

	result1, result2, err := Day01(&content)

	if err != nil {
		t.Fatal(err)
	}

	if result1 != "11" {
		t.Error("result1 be 11 but was " + result1)
	}

	if result2 != "31" {
		t.Error("result2 should be 31 but was " + result2)
	}
}
