package common

import (
	"bufio"
	"bytes"
)

func ReadByLine(content *[]byte) []string {
	scanner := bufio.NewScanner(bytes.NewReader(*content))
	scanner.Split(bufio.ScanLines)
	result := []string{}
	index := 0
	for scanner.Scan() {
		result = append(result, scanner.Text())
		index += 1
	}

	return result
}

func ReadByWord(content *[]byte) []string {
	scanner := bufio.NewScanner(bytes.NewReader(*content))
	scanner.Split(bufio.ScanWords)
	result := []string{}
	index := 0
	for scanner.Scan() {
		result = append(result, scanner.Text())
		index += 1
	}

	return result
}
