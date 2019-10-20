package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/rosenhouse/rmkb/models"
)

var _ = Describe("Runs", func() {
	Describe("NewRun", func() {
		It("constructs a new run", func() {
			run, err := NewRun(Tile{Blue, 1}, Tile{Blue, 2}, Tile{Blue, 3})
			Expect(err).NotTo(HaveOccurred())
			Expect(run.Length()).To(Equal(3))
		})
		It("ensures length", func() {
			_, err := NewRun(Tile{Blue, 1}, Tile{Blue, 2})
			Expect(err).To(MatchError(ErrSetTooShort))
		})
		It("ensures all tiles are the same color", func() {
			_, err := NewRun(Tile{Blue, 1}, Tile{Red, 2}, Tile{Blue, 3})
			Expect(err).To(MatchError(ErrRunNotMonochrome))
		})
		It("ensures the numbers are a sequence", func() {
			_, err := NewRun(Tile{Blue, 1}, Tile{Blue, 2}, Tile{Blue, 4})
			Expect(err).To(MatchError(ErrRunNonSequential))
		})
	})
})
