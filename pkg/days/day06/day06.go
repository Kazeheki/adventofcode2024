package days

import (
	"fmt"
	"strconv"
	"strings"

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
	board := newBoard(content)
	guard := board.guardStart
	board.markAsVisited(guard.position)

	for guard.stillWalksOnFloor(board) {
		if guard.canWalk(board) {
			guard.walk()
			board.markAsVisited(guard.position)
		} else {
			guard.turnRight()
		}
	}

	return strconv.Itoa(board.countVisited()), nil
}

func part2(content *[]byte) (string, error) {
	// todo: implement
	return "", nil
}

type direction int

const (
	UP direction = iota
	RIGHT
	DOWN
	LEFT
)

func (d direction) turnRight() direction {
	switch d {
	case UP:
		return RIGHT
	case RIGHT:
		return DOWN
	case DOWN:
		return LEFT
	case LEFT:
		return UP
	default:
		panic("unknown direction")
	}
}

type position struct {
	x int
	y int
}

func (p position) nextPositionInDirection(d direction) position {
	switch d {
	case UP:
		p.y = p.y - 1
	case RIGHT:
		p.x = p.x + 1
	case DOWN:
		p.y = p.y + 1
	case LEFT:
		p.x = p.x - 1
	default:
		panic("unknown direction")
	}
	return p
}

type guard struct {
	position      position
	lookDirection direction
}

func (g *guard) walk() {
	g.position = g.position.nextPositionInDirection(g.lookDirection)
}

func (g *guard) canWalk(b board) bool {
	target := g.position.nextPositionInDirection(g.lookDirection)
	return b.isValidPosition(target) && b.tileAt(target) != OBSTACLE
}

func (g *guard) stillWalksOnFloor(b board) bool {
	target := g.position.nextPositionInDirection(g.lookDirection)
	return b.isValidPosition(target)
}

func (g *guard) turnRight() {
	g.lookDirection = g.lookDirection.turnRight()
}

type tile int

const (
	EMPTY tile = iota
	OBSTACLE
	VISITED
)

type board struct {
	floorPlan  [][]tile
	height     int
	width      int
	guardStart guard
}

func newBoard(content *[]byte) board {
	input := common.ReadAsTwoDimensionalArray(content)

	board := board{
		floorPlan: make([][]tile, len(input)),
		height:    len(input),
		width:     len(input[0]),
	}
	for i := range board.height {
		board.floorPlan[i] = make([]tile, board.width)
	}

	for y, row := range input {
		for x, character := range row {
			var tile tile
			switch character {
			case '.':
				tile = EMPTY
			case '#':
				tile = OBSTACLE
			case '^':
				tile = VISITED
				board.guardStart = guard{position{x, y}, UP}
			default:
				panic(fmt.Sprintf("invalid floor item '%q'", character))
			}
			board.floorPlan[y][x] = tile
		}
	}

	return board
}

func (b *board) isValidPosition(p position) bool {
	validY := 0 <= p.y && p.y < b.height
	validX := 0 <= p.x && p.x < b.width
	return validY && validX
}

func (b *board) tileAt(p position) tile {
	return b.floorPlan[p.y][p.x]
}

func (b *board) markAsVisited(p position) {
	b.floorPlan[p.y][p.x] = VISITED
}

func (b *board) print() {
	var sb strings.Builder
	for _, row := range b.floorPlan {
		for _, tile := range row {
			var character string
			switch tile {
			case EMPTY:
				character = "."
			case OBSTACLE:
				character = "#"
			case VISITED:
				character = "V"
			}
			_, _ = fmt.Fprintf(&sb, "%s", character)
		}
		_, _ = fmt.Fprintf(&sb, "\n")
	}
	fmt.Print(sb.String())
}

func (b *board) countVisited() int {
	sum := 0
	for _, row := range b.floorPlan {
		for _, tile := range row {
			if tile == VISITED {
				sum += 1
			}
		}
	}
	return sum
}
