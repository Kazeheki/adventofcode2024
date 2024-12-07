package days

import (
	"log/slog"
	"testing"
)

func Test(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	input := `xmul(2,4)&mul[3,7]
!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

	content := []byte(input)

	result1, result2, err := Process(&content)

	if err != nil {
		t.Fatal(err)
	}

	expected := "161"
	if result1 != expected {
		t.Error("expected result1 to be " + expected + " but was " + result1)
	}

	expected = "48"
	if result2 != expected {
		t.Error("expected result2 to be " + expected + " but was " + result2)
	}
}
