package main

import (
	day01 "aoc2024/pkg/days/day01"
	day02 "aoc2024/pkg/days/day02"
	day03 "aoc2024/pkg/days/day03"
	day04 "aoc2024/pkg/days/day04"
	day05 "aoc2024/pkg/days/day05"
	day06 "aoc2024/pkg/days/day06"
	day07 "aoc2024/pkg/days/day07"
	day08 "aoc2024/pkg/days/day08"
	day09 "aoc2024/pkg/days/day09"
	day10 "aoc2024/pkg/days/day10"
	day11 "aoc2024/pkg/days/day11"
	day12 "aoc2024/pkg/days/day12"
	day13 "aoc2024/pkg/days/day13"
	day14 "aoc2024/pkg/days/day14"
	day15 "aoc2024/pkg/days/day15"
	day16 "aoc2024/pkg/days/day16"
	day17 "aoc2024/pkg/days/day17"
	day18 "aoc2024/pkg/days/day18"
	day19 "aoc2024/pkg/days/day19"
	day20 "aoc2024/pkg/days/day20"
	day21 "aoc2024/pkg/days/day21"
	day22 "aoc2024/pkg/days/day22"
	day23 "aoc2024/pkg/days/day23"
	day24 "aoc2024/pkg/days/day24"
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
		resultPart1, resultPart2, err = day01.Process(&content)
	case 2:
		resultPart1, resultPart2, err = day02.Process(&content)
	case 3:
		resultPart1, resultPart2, err = day03.Process(&content)
	case 4:
		resultPart1, resultPart2, err = day04.Process(&content)
	case 5:
		resultPart1, resultPart2, err = day05.Process(&content)
	case 6:
		resultPart1, resultPart2, err = day06.Process(&content)
	case 7:
		resultPart1, resultPart2, err = day07.Process(&content)
	case 8:
		resultPart1, resultPart2, err = day08.Process(&content)
	case 9:
		resultPart1, resultPart2, err = day09.Process(&content)
	case 10:
		resultPart1, resultPart2, err = day10.Process(&content)
	case 11:
		resultPart1, resultPart2, err = day11.Process(&content)
	case 12:
		resultPart1, resultPart2, err = day12.Process(&content)
	case 13:
		resultPart1, resultPart2, err = day13.Process(&content)
	case 14:
		resultPart1, resultPart2, err = day14.Process(&content)
	case 15:
		resultPart1, resultPart2, err = day15.Process(&content)
	case 16:
		resultPart1, resultPart2, err = day16.Process(&content)
	case 17:
		resultPart1, resultPart2, err = day17.Process(&content)
	case 18:
		resultPart1, resultPart2, err = day18.Process(&content)
	case 19:
		resultPart1, resultPart2, err = day19.Process(&content)
	case 20:
		resultPart1, resultPart2, err = day20.Process(&content)
	case 21:
		resultPart1, resultPart2, err = day21.Process(&content)
	case 22:
		resultPart1, resultPart2, err = day22.Process(&content)
	case 23:
		resultPart1, resultPart2, err = day23.Process(&content)
	case 24:
		resultPart1, resultPart2, err = day24.Process(&content)
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
