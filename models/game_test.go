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

		groupedByColor := GroupByColor(allTiles)
		Expect(groupedByColor).To(HaveLen(4))
		for _, tilesOfThisColor := range groupedByColor {
			Expect(tilesOfThisColor).To(HaveLen(26)) // 26 tiles for each color

			singleColorGroupedByNumber := GroupByNumber(tilesOfThisColor)
			Expect(singleColorGroupedByNumber).To(HaveLen(13))
			for _, tilesOfThisColorAndNumber := range singleColorGroupedByNumber {
				Expect(tilesOfThisColorAndNumber).To(HaveLen(2)) // two tiles of each (number,color)
			}
		}
	})
})
