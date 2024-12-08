package days

import (
	"fmt"
	"log/slog"
	"os"
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

	expected = ""
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

func TestPart1WithRealInput(t *testing.T) {
	input, err := os.ReadFile(fmt.Sprintf(`../../../input/day%02d`, 4))
	if err != nil {
		panic("could not find file")
	}
	content := []byte(input)
	_, _, err = Process(&content)

	if err != nil {
		t.Fatal(err)
	}
}

func TestRemoveMe(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	input := `1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
18
19
20
21
22
23
24
25
26
27
28
29
30
31
32
`

	content := []byte(input)
	_, _, err := Process(&content)

	if err != nil {
		t.Fatal(err)
	}
}
