package common

import (
	"bufio"
	"bytes"
	"log"
)

func ReadByLine(content *[]byte) []string {
	scanner := bufio.NewScanner(bytes.NewReader(*content))
	scanner.Split(bufio.ScanLines)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func ReadByWord(content *[]byte) []string {
	scanner := bufio.NewScanner(bytes.NewReader(*content))
	scanner.Split(bufio.ScanWords)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func ReadAsTwoDimensionalArray(content *[]byte) [][]byte {
	scanner := bufio.NewScanner(bytes.NewReader(*content))
	scanner.Split(bufio.ScanLines)
	var result [][]byte
	for scanner.Scan() {
		lineBytes := scanner.Bytes()
		// scanner.Bytes gives back the internal array that it might override again!
		// copy to not have strange issues with lines being overwritten.
		tmp := make([]byte, len(lineBytes))
		copy(tmp, lineBytes)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		result = append(result, tmp)
	}

	return result
}
