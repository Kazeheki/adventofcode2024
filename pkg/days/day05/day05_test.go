package days

import (
	"log/slog"
	"testing"
)

func TestWithExampleInput(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
	content := []byte(input)
	result1, result2, err := Process(&content)

	if err != nil {
		t.Fatal(err)
	}

	expected := "143"
	if result1 != expected {
		t.Error("expected result1 to be " + expected + " but was " + result1)
	}

	expected = "123"
	if result2 != expected {
		t.Error("expected result2 to be " + expected + " but was " + result2)
	}
}
