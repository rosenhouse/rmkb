package solve_test

import (
	"sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/rosenhouse/rmkb/solve"
)

var _ = Describe("Groups of tiles with the same number", func() {
	Describe("Groups", func() {
		It("is sorted", func() {
			Expect(sort.SliceIsSorted(Groups, func(i, j int) bool {
				return Groups[i] < Groups[j]
			})).To(BeTrue())
		})
	})

	Describe("GroupCombinations", func() {
		It("maps combos to their components", func() {
			for gc, dc := range GroupCombinations {
				Expect(Combine(dc[0], dc[1])).To(Equal(gc))
			}
		})
	})

	Describe("SortedGroupCombos", func() {
		It("Includes the empty group, the group of 4, and two groups of 4", func() {
			Expect(SortedGroupCombos).To(ContainElement(GroupCombo(0)))
			Expect(SortedGroupCombos).To(ContainElement(GroupCombo(OneBlack + OneBlue + OneOrange + OneRed)))
			Expect(SortedGroupCombos).To(ContainElement(GroupCombo(2 * (OneBlack + OneBlue + OneOrange + OneRed))))
		})
		It("contains the groups of 3", func() {
			Expect(SortedGroupCombos).To(ContainElement(GroupCombo(OneBlack + OneBlue + OneOrange))) // no red
			Expect(SortedGroupCombos).To(ContainElement(GroupCombo(OneBlack + OneBlue + OneRed)))    // no orange
			Expect(SortedGroupCombos).To(ContainElement(GroupCombo(OneBlack + OneOrange + OneRed)))  // no blue
			Expect(SortedGroupCombos).To(ContainElement(GroupCombo(OneBlue + OneOrange + OneRed)))   // no black
		})
		It("contains all pairs of groups", func() {
			singleGroups := []GroupCombo{
				GroupCombo(0),
				GroupCombo(OneBlack + OneBlue + OneOrange + OneRed),
				GroupCombo(OneBlack + OneBlue + OneOrange), // no red
				GroupCombo(OneBlack + OneBlue + OneRed),    // no orange
				GroupCombo(OneBlack + OneOrange + OneRed),  // no blue
				GroupCombo(OneBlue + OneOrange + OneRed),   // no black
			}
			for _, g1 := range singleGroups {
				for _, g2 := range singleGroups {
					Expect(SortedGroupCombos).To(ContainElement(g1 + g2))
				}
			}
		})
		It("has the correct number of elements", func() {
			const expectedCombinations = (6 * (6 + 1)) / 2 // formula for nChoose2
			Expect(GroupCombinations).To(HaveLen(expectedCombinations))
		})
		It("is sorted", func() {
			Expect(sort.SliceIsSorted(SortedGroupCombos, func(i, j int) bool {
				return SortedGroupCombos[i] < SortedGroupCombos[j]
			})).To(BeTrue())
		})

		It("contains the same elements as the map to compoennts", func() {
			Expect(SortedGroupCombos).To(HaveLen(len(GroupCombinations)))
			for i := range SortedGroupCombos {
				Expect(GroupCombinations).To(HaveKey(SortedGroupCombos[i]))
			}
		})
	})
})

var _ = Describe("FindGroupings", func() {
	Specify("it returns all valid group-combos and any tiles that remain", func() {
		Expect(FindGroupings(OneBlack + OneBlue + OneOrange)).To(Equal(GroupingOptions{
			GroupCombo_Nothing:               OneBlack + OneBlue + OneOrange,
			Combine(Group_NoRed, Group_None): EmptyStack,
		}))

		Expect(FindGroupings(OneBlack + OneBlue + 2*OneOrange)).To(Equal(GroupingOptions{
			GroupCombo_Nothing:               OneBlack + OneBlue + 2*OneOrange,
			Combine(Group_NoRed, Group_None): OneOrange,
		}))

		Expect(FindGroupings(OneBlack + OneBlue + OneOrange + OneRed)).To(Equal(GroupingOptions{
			GroupCombo_Nothing:                  OneBlack + OneBlue + OneOrange + OneRed,
			Combine(Group_All, Group_None):      EmptyStack,
			Combine(Group_NoBlue, Group_None):   OneBlue,
			Combine(Group_NoBlack, Group_None):  OneBlack,
			Combine(Group_NoOrange, Group_None): OneOrange,
			Combine(Group_NoRed, Group_None):    OneRed,
		}))

		Expect(FindGroupings(OneBlack + 2*OneBlue + 2*OneOrange + OneRed)).To(Equal(GroupingOptions{
			GroupCombo_Nothing:                  OneBlack + 2*OneBlue + 2*OneOrange + OneRed,
			Combine(Group_All, Group_None):      OneBlue + OneOrange,
			Combine(Group_NoBlue, Group_None):   2*OneBlue + OneOrange,
			Combine(Group_NoBlack, Group_None):  OneBlack + OneBlue + OneOrange,
			Combine(Group_NoOrange, Group_None): OneBlue + 2*OneOrange,
			Combine(Group_NoRed, Group_None):    OneBlue + OneOrange + OneRed,
			Combine(Group_NoBlack, Group_NoRed): EmptyStack,
		}))
	})
})

var _ = Describe("Contains", func() {
	Specify("without any tiles, only the empty group-combo fits", func() {
		for gc := range GroupCombinations {
			expectedResult := (gc == GroupCombo_Nothing)
			Expect(Contains(EmptyStack, gc)).To(Equal(expectedResult))
		}
	})
	Specify("for a single tile, only the Nothing combo fits", func() {
		for _, c := range Colors {
			for gc := range GroupCombinations {
				if gc == GroupCombo_Nothing {
					Expect(Contains(c, GroupCombo_Nothing)).To(BeTrue())
				} else {
					Expect(Contains(c, gc)).To(BeFalse())
				}
			}
		}
	})
	Specify("for two tiles, only the Nothing combo fits", func() {
		for _, c1 := range Colors {
			for _, c2 := range Colors {
				c := c1 + c2
				for gc := range GroupCombinations {
					expectedResult := (gc == GroupCombo_Nothing)
					Expect(Contains(c, gc)).To(Equal(expectedResult))
				}
			}
		}
	})
	Specify("with two of every tile, every combination fits", func() {
		all := 2*OneBlack + 2*OneBlue + 2*OneOrange + 2*OneRed

		for gc := range GroupCombinations {
			Expect(Contains(all, gc)).To(BeTrue(), "colorbits: %#b groupcombo %#b", all, gc)
		}
	})

	Specify("with one of every tile, the single-groups fit but doubles do not", func() {
		all := OneBlack + OneBlue + OneOrange + OneRed

		for gc, componentGroups := range GroupCombinations {
			expectedResult := (componentGroups[0] == Group_None) // single-group
			Expect(Contains(all, gc)).To(Equal(expectedResult), "colorbits: %#b groupcombo %#b", all, gc)
		}
	})

	Specify("when missing a color, the fits are correct", func() {
		Expect(Contains(OneBlue+OneBlack+OneOrange, Combine(Group_NoRed, Group_None))).To(BeTrue())
		Expect(Contains(2*OneBlue+OneBlack+OneOrange, Combine(Group_NoRed, Group_None))).To(BeTrue())
		Expect(Contains(OneBlue+2*OneBlack+OneOrange, Combine(Group_NoRed, Group_None))).To(BeTrue())
		Expect(Contains(OneBlue+OneBlack+2*OneOrange, Combine(Group_NoRed, Group_None))).To(BeTrue())
		Expect(Contains(2*(OneBlue+OneBlack+OneOrange), Combine(Group_NoRed, Group_None))).To(BeTrue())
		Expect(Contains(2*(OneBlue+OneBlack+OneOrange), Combine(Group_NoRed, Group_NoRed))).To(BeTrue())

		Expect(Contains(OneBlue+OneBlack+OneOrange, Combine(Group_NoRed, Group_NoRed))).To(BeFalse())
		Expect(Contains(OneBlue+OneBlack+OneOrange, Combine(Group_NoOrange, Group_None))).To(BeFalse())
	})

})
