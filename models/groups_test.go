package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/rosenhouse/rmkb/models"
)

var _ = Describe("Groups", func() {
	Describe("NewGroup", func() {
		It("constructs a new group", func() {
			group, err := NewGroup(Tile{Black, 1}, Tile{Blue, 1}, Tile{Orange, 1})
			Expect(err).NotTo(HaveOccurred())
			Expect(group.Length()).To(Equal(3))
			Expect(group.CommonNumber()).To(Equal(1))
		})
		It("allows groups of four", func() {
			group, err := NewGroup(Tile{Black, 13}, Tile{Blue, 13}, Tile{Orange, 13}, Tile{Red, 13})
			Expect(err).NotTo(HaveOccurred())
			Expect(group.Length()).To(Equal(4))
			Expect(group.CommonNumber()).To(Equal(13))
		})

		It("ensures length", func() {
			_, err := NewGroup(Tile{Black, 1}, Tile{Red, 1})
			Expect(err).To(MatchError(ErrSetTooShort))
		})

		It("ensures all the tiles have the same number", func() {
			_, err := NewGroup(Tile{Black, 1}, Tile{Blue, 2}, Tile{Orange, 1})
			Expect(err).To(MatchError(ErrGroupNotUnitary))
		})

		It("ensures no color is repeated", func() {
			_, err := NewGroup(Tile{Black, 1}, Tile{Blue, 1}, Tile{Black, 1})
			Expect(err).To(MatchError(ErrGroupDuplicateColors))
		})
	})
})
