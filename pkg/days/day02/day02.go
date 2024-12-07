package days

import (
	"aoc2024/pkg/common"
	"math"
	"strconv"
	"strings"
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
	lines := common.ReadByLine(content)
	counter := 0

	for _, line := range lines {
		stringSlice := strings.Split(line, " ")
		intSlice := mapStringSliceToIntSlice(stringSlice)

		if isValid(intSlice) {
			counter += 1
		}
	}

	return strconv.Itoa(counter), nil
}

func part2(content *[]byte) (string, error) {
	lines := common.ReadByLine(content)
	counter := 0

	for _, line := range lines {
		stringSlice := strings.Split(line, " ")
		intSlice := mapStringSliceToIntSlice(stringSlice)

		if isValid(intSlice) {
			counter += 1
		} else {
			for i := range len(intSlice) {
				newSlice := cutOut(intSlice, i)
				if isValid(newSlice) {
					counter += 1
					break
				}
			}
		}
	}

	return strconv.Itoa(counter), nil
}

func cutOut(slice []int, at int) []int {
	internalCopy := make([]int, len(slice))
	copy(internalCopy, slice)
	return append(internalCopy[:at], internalCopy[at+1:]...)
}

func mapStringSliceToIntSlice(input []string) []int {
	var intSlice []int = make([]int, len(input))
	for index, value := range input {
		str, err := strconv.Atoi(value)
		if err != nil {
			panic("something went very wrong")
		}
		intSlice[index] = str
	}
	return intSlice
}

func isValid(arr []int) bool {
	if len(arr) <= 2 {
		return true
	}

	direction := getDirection(arr[0], arr[1])

	valid := true
	for i := 0; i < len(arr)-1 && valid; i++ {
		a := arr[i]
		b := arr[i+1]
		if getDirection(a, b) != direction {
			valid = false
		}
		abs := math.Abs(float64(a - b))
		if abs < 1 || abs > 3 {
			valid = false
		}
	}
	return valid
}

func getDirection(from int, to int) int {
	if from-to > 0 {
		return 1
	}
	return -1
}
