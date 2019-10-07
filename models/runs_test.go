package models_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/rosenhouse/rmkb/models"
)

var _ = Describe("Runs", func() {
	Describe("NewRun", func() {
		It("constructs a new run", func() {
			run, err := NewRun(Tile{ColorBlue, 1}, Tile{ColorBlue, 2}, Tile{ColorBlue, 3})
			Expect(err).NotTo(HaveOccurred())
			Expect(run.Length()).To(Equal(3))
		})
		It("ensures length", func() {
			_, err := NewRun(Tile{ColorBlue, 1}, Tile{ColorBlue, 2})
			Expect(err).To(MatchError(ErrSetTooShort))
		})
		It("ensures all tiles are the same color", func() {
			_, err := NewRun(Tile{ColorBlue, 1}, Tile{ColorRed, 2}, Tile{ColorBlue, 3})
			Expect(err).To(MatchError(ErrRunNotMonochrome))
		})
		It("ensures the numbers are a sequence", func() {
			_, err := NewRun(Tile{ColorBlue, 1}, Tile{ColorBlue, 2}, Tile{ColorBlue, 4})
			Expect(err).To(MatchError(ErrRunNonSequential))
		})
	})
})
