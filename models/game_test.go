package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/rosenhouse/rmkb/models"
)

var _ = Describe("NewGame", func() {
	It("returns a GameState", func() {
		players := []string{"player 1", "player 2", "player 3"}
		gs, err := NewGame(1, players...)
		Expect(err).NotTo(HaveOccurred())
		Expect(gs.Done()).To(BeFalse())

		Expect(gs.Players).To(HaveLen(3))

		for i, p := range gs.Players {
			Expect(p.Name).To(Equal(players[i]))
		}
	})

	It("validates players", func() {
		_, err := NewGame(1, "player 1", "player 2", "")
		Expect(err).To(MatchError(ErrPlayers))

		_, err = NewGame(1, "player 1")
		Expect(err).To(MatchError(ErrPlayers))
	})

	It("populates the pool of tiles", func() {
		gs, err := NewGame(1, "player 1", "player 2")
		Expect(err).NotTo(HaveOccurred())
		Expect(gs.Pool).To(HaveLen(104))
		Expect(gs.Pool).To(ConsistOf(BuildTiles()))
	})

	It("randomizes the order of tiles in the pool", func() {
		gs1, err := NewGame(1, "player 1", "player 2")
		Expect(err).NotTo(HaveOccurred())

		gs2, err := NewGame(2, "player 1", "player 2")
		Expect(err).NotTo(HaveOccurred())

		gs3, err := NewGame(3, "player 1", "player 2")
		Expect(err).NotTo(HaveOccurred())

		Expect(gs1.Pool).NotTo(Equal(gs2.Pool))
		Expect(gs2.Pool).NotTo(Equal(gs3.Pool))
		Expect(gs3.Pool).NotTo(Equal(gs1.Pool))
	})
})

var _ = Describe("BuildTiles", func() {
	It("returns the full set of tiles", func() {
		allTiles := BuildTiles()
		Expect(allTiles).To(HaveLen(104))

		groupedByColor := CollectByColor(allTiles)
		Expect(groupedByColor).To(HaveLen(4))
		for _, tilesOfThisColor := range groupedByColor {
			Expect(tilesOfThisColor).To(HaveLen(26)) // 26 tiles for each color

			singleColorGroupedByNumber := CollectByNumber(tilesOfThisColor)
			Expect(singleColorGroupedByNumber).To(HaveLen(13))
			for _, tilesOfThisColorAndNumber := range singleColorGroupedByNumber {
				Expect(tilesOfThisColorAndNumber).To(HaveLen(2)) // two tiles of each (number,color)
			}
		}
	})
})

var _ = Describe("Board basics", func() {
	It("validates a board", func() {
		before := GameState{
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
		after := GameState{
			Pool: []Tile{
				{Color: Orange, Number: 9},
				{Color: Black, Number: 9},
				{Color: Orange, Number: 9},
			},
			Board: Board{
				Groups: []Group{
					Gr(5, Blue, Black, Orange),
					Gr(6, Blue, Black, Red),
					Gr(7, Blue, Black, Red),
					Gr(8, Red, Black, Orange),

					//					Gr(9, Orange, Black, Orange), // invalid

					Gr(10, Blue, Red, Orange),
					Gr(11, Blue, Black, Orange, Red),
					Gr(12, Blue, Black, Orange),
					Gr(13, Blue, Black, Orange),

					Gr(4, Blue, Red, Black),
					Gr(3, Blue, Orange, Black),
					Gr(1, Blue, Orange, Red),

					Gr(13, Black, Red, Blue),
				},

				Runs: []Run{
					Rn(Black, 1, 2, 3),
					Rn(Red, 1, 2, 3),

					Rn(Black, 5, 6, 7, 8, 9, 10, 11, 12),
					Rn(Red, 9, 10, 11),

					Rn(Blue, 6, 7, 8, 9, 10),
					Rn(Orange, 6, 7, 8),
				},
			},
		}
		beforeAllTiles := append(before.Board.TilesInPlay(), before.Pool...)
		afterAllTiles := append(after.Board.TilesInPlay(), after.Pool...)

		SortTiles(beforeAllTiles)
		SortTiles(afterAllTiles)

		Expect(beforeAllTiles).To(Equal(afterAllTiles))

	})
})

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
