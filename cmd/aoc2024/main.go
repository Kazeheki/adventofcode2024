package main

import (
	"aoc2024/pkg/days"
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		slog.Error("Expected one argument (day to run)")
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])
	handleErrorWithExit(err)

	fmt.Printf(`You chose day %02d%v`, day, "\n")

	content, err := readInputForDay(day)
	handleErrorWithExit(err)

	var resultPart1 string
	var resultPart2 string
	switch day {
	case 1:
		resultPart1, resultPart2, err = days.Day01(&content)
	case 2:
		resultPart1, resultPart2, err = days.Day02(&content)
	case 3:
		resultPart1, resultPart2, err = days.Day03(&content)
	case 4:
		resultPart1, resultPart2, err = days.Day04(&content)
	case 5:
		resultPart1, resultPart2, err = days.Day05(&content)
	case 6:
		resultPart1, resultPart2, err = days.Day06(&content)
	case 7:
		resultPart1, resultPart2, err = days.Day07(&content)
	case 8:
		resultPart1, resultPart2, err = days.Day08(&content)
	case 9:
		resultPart1, resultPart2, err = days.Day09(&content)
	case 10:
		resultPart1, resultPart2, err = days.Day10(&content)
	case 11:
		resultPart1, resultPart2, err = days.Day11(&content)
	case 12:
		resultPart1, resultPart2, err = days.Day12(&content)
	case 13:
		resultPart1, resultPart2, err = days.Day13(&content)
	case 14:
		resultPart1, resultPart2, err = days.Day14(&content)
	case 15:
		resultPart1, resultPart2, err = days.Day15(&content)
	case 16:
		resultPart1, resultPart2, err = days.Day16(&content)
	case 17:
		resultPart1, resultPart2, err = days.Day17(&content)
	case 18:
		resultPart1, resultPart2, err = days.Day18(&content)
	case 19:
		resultPart1, resultPart2, err = days.Day19(&content)
	case 20:
		resultPart1, resultPart2, err = days.Day20(&content)
	case 21:
		resultPart1, resultPart2, err = days.Day21(&content)
	case 22:
		resultPart1, resultPart2, err = days.Day22(&content)
	case 23:
		resultPart1, resultPart2, err = days.Day23(&content)
	case 24:
		resultPart1, resultPart2, err = days.Day24(&content)
	}
	handleErrorWithExit(err)
	handleResult(resultPart1, resultPart2)
}

func handleResult(resultPart1 string, resultPart2 string) {
	fmt.Printf(`Part1: %[1]s%[3]vPart2: %[2]s%[3]v`, resultPart1, resultPart2, "\n")
}

func readInputForDay(day int) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf(`input/day%02d`, day))
}

func handleErrorWithExit(err error) {
	if err != nil {
		slog.Error("An error occurred, please check", "error", err)
		os.Exit(1)
	}
}
