package days

import (
	"log/slog"
	"testing"
)

func Test(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	content := []byte(input)

	result1, result2, err := Process(&content)

	if err != nil {
		t.Fatal(err)
	}

	expected := "2"
	if result1 != expected {
		t.Error("expected result1 to be " + expected + " but was " + result1)
	}

	expected = "4"
	if result2 != expected {
		t.Error("expected result2 to be " + expected + " but was " + result2)
	}
}
