package solve

import (
	"fmt"
	"sort"
)

const (
	IdxBlack = iota
	IdxBlue
	IdxOrange
	IdxRed
)

type ColorBits byte

const MaskCount byte = 0b00000011
const BitsPerColor = 2
const NumColors = 4

// TileStack represents a collection of tiles sharing a common number but potentially different colors
// The collection may contain zero, one or two of each color.
//
// The count of each color is represented with 2-bits
type TileStack byte

func CountOfColor(tilestack TileStack, colorIndex int) int {
	return int(tilestack>>(BitsPerColor*colorIndex)) & int(MaskCount)
}

const EmptyStack TileStack = 0

const (
	OneBlack TileStack = 0b00000001 << (2 * iota)
	OneBlue
	OneOrange
	OneRed
)

var Colors = []TileStack{OneBlack, OneBlue, OneOrange, OneRed}

type Group byte

const (
	Group_None     = Group(0)
	Group_All      = Group(OneBlack + OneBlue + OneOrange + OneRed)
	Group_NoBlack  = Group(OneBlue + OneOrange + OneRed)
	Group_NoBlue   = Group(OneBlack + OneOrange + OneRed)
	Group_NoOrange = Group(OneBlack + OneBlue + OneRed)
	Group_NoRed    = Group(OneBlack + OneBlue + OneOrange)
)

var groupNames = map[Group]string{
	Group_None:     "None",
	Group_NoRed:    "NoRed",
	Group_NoOrange: "NoOrange",
	Group_NoBlue:   "NoBlue",
	Group_NoBlack:  "NoBlack",
	Group_All:      "All",
}

func (g Group) String() string {
	return groupNames[g]
}

var Groups = []Group{
	Group_None,
	Group_NoRed,
	Group_NoOrange,
	Group_NoBlue,
	Group_NoBlack,
	Group_All,
}

// GroupCombo represents how a set of tiles with a common number may be formed into Groups
type GroupCombo byte

const GroupCombo_Nothing GroupCombo = 0

func Combine(g1, g2 Group) GroupCombo {
	return GroupCombo(g1 + g2)
}

// GroupCombinations are valid pairs of Groups
// There are 21 valid combinations
var GroupCombinations = (func() map[GroupCombo][2]Group {
	combos := map[GroupCombo][2]Group{}
	for i := range Groups {
		for j := i; j < len(Groups); j++ {
			gc := Combine(Groups[i], Groups[j])
			combos[gc] = [2]Group{Groups[i], Groups[j]}
		}
	}
	return combos
})()

var SortedGroupCombos = (func() []GroupCombo {
	all := make([]GroupCombo, 0, len(GroupCombinations))
	for gc := range GroupCombinations {
		all = append(all, gc)
	}
	sort.Slice(all, func(i, j int) bool {
		return all[i] < all[j]
	})
	return all
})()

var comboNames = (func() map[GroupCombo]string {
	names := map[GroupCombo]string{}
	for gc, dc := range GroupCombinations {
		names[gc] = fmt.Sprintf("%s+%s", dc[0], dc[1])
	}
	return names
})()

func (gc GroupCombo) String() string {
	return comboNames[gc]
}

// GroupingOptions represents ways to form Groups
// A single map key represents one possible way of forming Groups
// and the corresponding map value represents the remaining tiles after that grouping is applied
// Remaining tiles would need to be covered by a Run
type GroupingOptions map[GroupCombo]TileStack

// FindGroupings returns all possible options for grouping the tiles in the given tilestack
// without accounting for the remainders (which would need to be covered by Runs)
func FindGroupings(tilestack TileStack) GroupingOptions {
	groupings := GroupingOptions{}
	for gc := range GroupCombinations {
		if Contains(tilestack, gc) {
			groupings[gc] = TileStack(byte(tilestack) - byte(gc))
		}
	}
	return groupings
}

// Contains indicates if the given tilestack contain the given group-combo
func Contains(tiles TileStack, gc GroupCombo) bool {
	t := byte(tiles)
	g := byte(gc)
	m := byte(MaskCount)

	ok := (t & m) >= (g & m)

	t = t >> BitsPerColor
	g = g >> BitsPerColor
	ok = ok && ((t & m) >= (g & m))

	t = t >> BitsPerColor
	g = g >> BitsPerColor
	ok = ok && ((t & m) >= (g & m))

	t = t >> BitsPerColor
	g = g >> BitsPerColor
	ok = ok && ((t & m) >= (g & m))

	return ok
}
