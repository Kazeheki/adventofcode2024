package days

import (
	"regexp"
	"strconv"

	"aoc2024/pkg/common"
)

func Process(content *[]byte) (string, string, error) {
	result1, err := part1(content)
	if err != nil {
		return "", "", err
	}
	result2, err := part2(content)
	if err != nil {
		return result1, "", err
	}
	return result1, result2, nil
}

func part1(content *[]byte) (string, error) {
	input := common.ReadAsTwoDimensionalArray(content)

	findXRegEx := regexp.MustCompile("X")

	height := Range{start: 0, end: len(input)}

	sum := 0
	for y, line := range input {
		for _, indexes := range findXRegEx.FindAllIndex(line, -1) {
			x := indexes[0]
			for _, direction := range Directions {
				if hasXMAS(input, x, y, height, direction) {
					sum += 1
				}
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func part2(content *[]byte) (string, error) {
	input := common.ReadAsTwoDimensionalArray(content)

	findARegEx := regexp.MustCompile("A")

	height := Range{start: 0, end: len(input)}

	sum := 0
	for y, line := range input {
		for _, indexes := range findARegEx.FindAllIndex(line, -1) {
			x := indexes[0]
			if hasXShapeMAS(input, x, y, height) {
				sum += 1
			}
		}
	}

	return strconv.Itoa(sum), nil
}

func hasXShapeMAS(input [][]byte, x int, y int, height Range) bool {
	// x shape MAS
	// may be
	// M.M   S.S   M.S   S.M
	// .A.   .A.   .A.   .A.
	// S.S   M.M   M.S   S.M
	if !height.canHandle(y-1) || !height.canHandle(y+1) {
		return false
	}
	// assuming all have same width
	width := Range{start: 0, end: len(input[y])}
	if !width.canHandle(x-1) || !width.canHandle(x+1) {
		return false
	}

	diagonalStr1 := ""
	for _, direction := range []Direction{TopLeft, Middle, BottomRight} {
		dirX, dirY := direction.Values()
		newX := x + dirX
		newY := y + dirY
		diagonalStr1 += string(input[newY][newX])
	}
	if diagonalStr1 != "MAS" && diagonalStr1 != "SAM" {
		return false
	}

	diagonalStr2 := ""
	for _, direction := range []Direction{TopRight, Middle, BottomLeft} {
		dirX, dirY := direction.Values()
		newX := x + dirX
		newY := y + dirY
		diagonalStr2 += string(input[newY][newX])
	}
	return diagonalStr2 == "MAS" || diagonalStr2 == "SAM"
}

func hasXMAS(input [][]byte, x int, y int, height Range, direction Direction) bool {
	dirX, dirY := direction.Values()
	str := ""
	for step := range len("XMAS") {
		nextX := x + (dirX * step)
		nextY := y + (dirY * step)
		if !height.canHandle(nextY) {
			break
		}
		width := Range{start: 0, end: len(input[nextY])}
		if !width.canHandle(nextX) {
			break
		}
		str += string(input[nextY][nextX])
	}
	return str == "XMAS"
}

type Range struct {
	start int
	end   int
}

func (r *Range) canHandle(val int) bool {
	return r.start <= val && val < r.end
}

type Direction int

const (
	TopLeft Direction = iota
	TopMiddle
	TopRight
	MiddleRight
	BottomRight
	BottomMiddle
	BottomLeft
	MiddleLeft
	Middle
)

var Directions = []Direction{
	TopLeft,
	TopMiddle,
	TopRight,
	MiddleRight,
	BottomRight,
	BottomMiddle,
	BottomLeft,
	MiddleLeft,
}

func (d *Direction) Values() (int, int) {
	switch *d {
	case TopLeft:
		return -1, -1
	case TopMiddle:
		return 0, -1
	case TopRight:
		return 1, -1
	case MiddleRight:
		return 1, 0
	case BottomRight:
		return 1, 1
	case BottomMiddle:
		return 0, 1
	case BottomLeft:
		return -1, 1
	case MiddleLeft:
		return -1, 0
	default:
		return 0, 0
	}
}

func (d *Direction) String() string {
	switch *d {
	case TopLeft:
		return "TopLeft"
	case TopMiddle:
		return "TopMiddle"
	case TopRight:
		return "TopRight"
	case MiddleRight:
		return "MiddleRight"
	case BottomRight:
		return "BottomRight"
	case BottomMiddle:
		return "BottomMiddle"
	case BottomLeft:
		return "BottomLeft"
	case MiddleLeft:
		return "MiddleLeft"
	default:
		return "unknown"
	}
}
