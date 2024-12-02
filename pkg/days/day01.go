package days

import (
	"aoc2024/pkg/common"
	"log/slog"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Day01(content *[]byte) (string, string, error) {
	part1Result, err := part1(content)
	if err != nil {
		return "", "", err
	}
	part2Result, err := part2(content)
	if err != nil {
		return part1Result, "", err
	}
	return part1Result, part2Result, nil
}

func part1(content *[]byte) (string, error) {
	left, right, err := prepare(content)
	if err != nil {
		return "", err
	}
	sort.Ints(left)
	sort.Ints(right)

	var sum float64 = 0
	// both left and right have the same length
	for i := 0; i < len(left); i++ {
		a := float64(left[i])
		b := float64(right[i])
		dist := math.Abs(a - b)
		sum += dist
	}

	return strconv.FormatFloat(sum, 'f', 0, 64), nil
}

func part2(content *[]byte) (string, error) {
	left, right, err := prepare(content)
	if err != nil {
		return "", err
	}

	result := 0
	for _, a := range left {
		countInRight := 0
		for _, b := range right {
			if a == b {
				countInRight++
			}
		}
		result += a * countInRight
	}

	return strconv.Itoa(result), nil
}

func prepare(content *[]byte) ([]int, []int, error) {
	lines := common.ReadByLine(content)
	size := len(lines)
	left := make([]int, size)
	right := make([]int, size)

	regexMultiSpace := regexp.MustCompile(`\s+`)
	for index, line := range lines {
		line = regexMultiSpace.ReplaceAllString(line, " ")
		split := strings.Split(line, " ")
		value, err := strconv.Atoi(split[0])
		if err != nil {
			slog.Debug("issue with left side of split", "error", err)
			return nil, nil, err
		}
		left[index] = value
		value, err = strconv.Atoi(split[1])
		if err != nil {
			slog.Debug("issue with right side of split", "error", err)
			return nil, nil, err
		}
		right[index] = value
	}
	return left, right, nil
}
