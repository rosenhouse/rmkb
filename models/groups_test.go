package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/rosenhouse/rmkb/models"
)

var _ = Describe("Groups", func() {
	Describe("NewGroup", func() {
		It("constructs a new group", func() {
			group, err := NewGroup(Tile{ColorBlack, 1}, Tile{ColorBlue, 1}, Tile{ColorOrange, 1})
			Expect(err).NotTo(HaveOccurred())
			Expect(group.Length()).To(Equal(3))
			Expect(group.Number()).To(Equal(1))
		})
		It("allows groups of four", func() {
			group, err := NewGroup(Tile{ColorBlack, 13}, Tile{ColorBlue, 13}, Tile{ColorOrange, 13}, Tile{ColorRed, 13})
			Expect(err).NotTo(HaveOccurred())
			Expect(group.Length()).To(Equal(4))
			Expect(group.Number()).To(Equal(13))
		})

		It("ensures length", func() {
			_, err := NewGroup(Tile{ColorBlack, 1}, Tile{ColorRed, 1})
			Expect(err).To(MatchError(ErrSetTooShort))
		})

		It("ensures all the tiles have the same number", func() {
			_, err := NewGroup(Tile{ColorBlack, 1}, Tile{ColorBlue, 2}, Tile{ColorOrange, 1})
			Expect(err).To(MatchError(ErrGroupNotUnitary))
		})

		It("ensures no color is repeated", func() {
			_, err := NewGroup(Tile{ColorBlack, 1}, Tile{ColorBlue, 1}, Tile{ColorBlack, 1})
			Expect(err).To(MatchError(ErrGroupDuplicateColors))
		})
	})
})
