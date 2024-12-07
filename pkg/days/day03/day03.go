package days

import (
	"regexp"
	"strconv"
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
	sum := execMultiply(content)
	return strconv.Itoa(sum), nil
}

func part2(content *[]byte) (string, error) {
	cleanUpRegEx := regexp.MustCompile(`\r?\n`)
	input := cleanUpRegEx.ReplaceAll(*content, []byte{})
	regEx := regexp.MustCompile(`(?:^|do\(\)).*?(?:$|don't\(\))`)
	matches := regEx.FindAll(input, -1)

	sum := 0
	for _, match := range matches {
		sum += execMultiply(&match)
	}

	return strconv.Itoa(sum), nil
}

func execMultiply(bytes *[]byte) int {
	multiplyRegEx := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := multiplyRegEx.FindAllSubmatch(*bytes, -1)

	sum := 0
	for _, match := range matches {
		first, _ := strconv.Atoi(string(match[1]))
		second, _ := strconv.Atoi(string(match[2]))
		sum += first * second
	}

	return sum
}
