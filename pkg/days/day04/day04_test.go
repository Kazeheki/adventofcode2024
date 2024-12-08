package days

import (
	"log/slog"
	"strconv"
	"testing"
)

func TestWithExampleInput(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	content := []byte(input)
	result1, result2, err := Process(&content)

	if err != nil {
		t.Fatal(err)
	}

	expected := "18"
	if result1 != expected {
		t.Error("expected result1 to be " + expected + " but was " + result1)
	}

	expected = "9"
	if result2 != expected {
		t.Error("expected result2 to be " + expected + " but was " + result2)
	}
}

func TestPart1(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	testCases := []struct {
		input       string
		expected    int
		description string
	}{
		{input: `XMAS`, expected: 1, description: "forward"},
		{input: `SAMX`, expected: 1, description: "backwards"},
		{
			input: `X
M
A
S`,
			expected:    1,
			description: "vertical TB",
		},
		{
			input: `S
A
M
X`,
			expected:    1,
			description: "vertical BT",
		},
		{
			input: `X...
.M..
..A.
...S`,
			expected:    1,
			description: "diagonal TL-BR",
		},
		{
			input: `...X
..M.
.A..
S...`,
			expected:    1,
			description: "diagonal TR-BL",
		},
		{
			input: `S...
.A..
..M.
...X`,
			expected:    1,
			description: "diagonal BR-TL",
		},
		{
			input: `...S
..A.
.M..
X...`,
			expected:    1,
			description: "diagonal BL-TR",
		},
		{
			input: `
...S
..A.
.M..
X...
`,
			expected:    1,
			description: "diagonal BL-TR with leading/tailing newlines",
		},
		{
			input: `

...S
..A.
.M..
X...

`,
			expected:    1,
			description: "diagonal BL-TR with leading/tailing newlines2",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			content := []byte(testCase.input)
			result1, _, err := Process(&content)

			if err != nil {
				t.Fatal(err)
			}

			expected := strconv.Itoa(testCase.expected)
			if result1 != expected {
				t.Error("expected result1 to be " + expected + " but was " + result1)
			}
		})
	}
}
