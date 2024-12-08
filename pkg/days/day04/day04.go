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

func part2(content *[]byte) (string, error) {
	return "", nil
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
