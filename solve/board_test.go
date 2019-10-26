package solve_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/rosenhouse/rmkb/models"

	"github.com/rosenhouse/rmkb/solve"
)

var _ = Describe("Board", func() {
	It("generates grouping options", func() {
		board := FromGameState(buildTestData())
		boardOptions := solve.AllGroupingOptions(*board)
		Expect(boardOptions).To(HaveLen(13))
	})

	It("solves?", func() {
		board := FromGameState(buildTestData())
		boardOptions := solve.AllGroupingOptions(*board)
		solutions := solve.FindAllValidOptions(boardOptions)
		Expect(solutions).To(HaveLen(3))
	})
})

func FromGameState(gs GameState) *solve.UnstructuredBoard {
	board := &solve.UnstructuredBoard{}
	accumulate(board, gs.Pool)
	fmt.Printf("with pool:   %08b\n", board)
	for _, g := range gs.Board.Groups {
		accumulate(board, g.Tiles())
	}
	fmt.Printf("with groups: %08b\n", board)
	for _, r := range gs.Board.Runs {
		accumulate(board, r.Tiles())
	}
	fmt.Printf("with runs:   %08b\n", board)
	return board
}

func accumulate(board *solve.UnstructuredBoard, tiles []Tile) {
	for _, t := range tiles {
		board[t.Number-1] += ConvertColor(t.Color)
	}
}

func ConvertColor(c Color) solve.TileStack {
	return solve.Colors[ColorIndex[c]]
}

var buildTestData = func() GameState {
	return GameState{
		Pool: []Tile{
			{Color: Red, Number: 11},
			{Color: Black, Number: 5},
			{Color: Blue, Number: 10},
			{Color: Red, Number: 10},
		},
		Board: Board{
			Groups: []Group{
				Gr(5, Blue, Black, Orange),
				Gr(6, Blue, Black, Orange),
				Gr(7, Blue, Black, Orange),

				Gr(8, Black, Orange, Red),
				Gr(8, Blue, Black, Orange),
				Gr(9, Blue, Black, Orange),
				Gr(10, Blue, Orange, Red),
				Gr(11, Blue, Black, Orange, Red),
				Gr(12, Blue, Black, Orange),
				Gr(13, Blue, Black, Orange),

				Gr(9, Orange, Black, Red),

				Gr(4, Black, Blue, Red),
				Gr(1, Orange, Blue, Red),

				Gr(13, Blue, Black, Red),

				Gr(3, Black, Orange, Blue),

				Gr(6, Red, Black, Blue),
				Gr(7, Red, Black, Blue),
			},
			Runs: []Run{
				Rn(Black, 10, 11, 12),
				Rn(Black, 1, 2, 3),
				Rn(Red, 1, 2, 3),
			},
		},
	}
}

func Gr(number int, colors ...Color) Group {
	tiles := []Tile{}
	for _, c := range colors {
		tiles = append(tiles, Tile{Color: c, Number: number})
	}
	group, err := NewGroup(tiles...)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return group
}

func Rn(color Color, numbers ...int) Run {
	tiles := []Tile{}
	for _, n := range numbers {
		tiles = append(tiles, Tile{Color: color, Number: n})
	}
	run, err := NewRun(tiles...)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())
	return run
}
