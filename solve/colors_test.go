package solve_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/rosenhouse/rmkb/solve"
)

var _ = Describe("CountOfColor", func() {
	It("counts the number of tiles of a given color within a tilestack", func() {
		Expect(CountOfColor(EmptyStack, IdxBlack)).To(Equal(0))
		Expect(CountOfColor(EmptyStack, IdxBlue)).To(Equal(0))
		Expect(CountOfColor(EmptyStack, IdxOrange)).To(Equal(0))
		Expect(CountOfColor(EmptyStack, IdxRed)).To(Equal(0))

		Expect(CountOfColor(OneBlack, IdxBlack)).To(Equal(1))
		Expect(CountOfColor(OneBlack, IdxBlue)).To(Equal(0))
		Expect(CountOfColor(OneBlack, IdxOrange)).To(Equal(0))
		Expect(CountOfColor(OneBlack, IdxRed)).To(Equal(0))

		Expect(CountOfColor(2*OneBlack, IdxBlack)).To(Equal(2))
		Expect(CountOfColor(2*OneBlack, IdxBlue)).To(Equal(0))
		Expect(CountOfColor(2*OneBlack, IdxOrange)).To(Equal(0))
		Expect(CountOfColor(2*OneBlack, IdxRed)).To(Equal(0))

		Expect(CountOfColor(2*OneBlack+OneBlue+2*OneRed, IdxBlack)).To(Equal(2))
		Expect(CountOfColor(2*OneBlack+OneBlue+2*OneRed, IdxBlue)).To(Equal(1))
		Expect(CountOfColor(2*OneBlack+OneBlue+2*OneRed, IdxOrange)).To(Equal(0))
		Expect(CountOfColor(2*OneBlack+OneBlue+2*OneRed, IdxRed)).To(Equal(2))
	})
})

var _ = Describe("SumWithCeiling", func() {
	It("does basic addition when the result is <= 2", func() {
		Expect(SumWithCeiling(0, 0)).To(Equal(TileStack(0)))
		Expect(SumWithCeiling(OneBlack, OneOrange)).To(Equal(OneBlack + OneOrange))
		Expect(SumWithCeiling(OneBlack, OneBlack)).To(Equal(2 * OneBlack))
		Expect(SumWithCeiling(OneRed, 0)).To(Equal(OneRed))
		Expect(SumWithCeiling(2*OneBlack, 2*OneOrange)).To(Equal(2*OneBlack + 2*OneOrange))
	})

	It("limits the result to no more than 2 per color", func() {
		Expect(SumWithCeiling(OneBlack, 2*OneBlack)).To(Equal(2 * OneBlack))
		Expect(SumWithCeiling(2*OneBlack, 2*OneBlack)).To(Equal(2 * OneBlack))
		Expect(SumWithCeiling(2*OneBlue, OneBlue)).To(Equal(2 * OneBlue))
	})
})
