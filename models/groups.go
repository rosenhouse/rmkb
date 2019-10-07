package models

import "errors"

var (
	ErrGroupNotUnitary      = errors.New("group must be a single number")
	ErrGroupDuplicateColors = errors.New("group must not reuse a color")
)

type Group struct {
	tiles map[Tile]struct{}
}

func (g *Group) Length() int {
	return len(g.tiles)
}

func (g *Group) Number() int {
	for tile, _ := range g.tiles {
		return tile.Number
	}
	panic("empty group")
}

func NewGroup(tiles ...Tile) (*Group, error) {
	if len(tiles) < 3 {
		return nil, ErrSetTooShort
	}
	if !allDifferentColors(tiles) {
		return nil, ErrGroupDuplicateColors
	}
	if !allSameNumber(tiles) {
		return nil, ErrGroupNotUnitary
	}

	asMap := make(map[Tile]struct{})
	for _, t := range tiles {
		asMap[t] = struct{}{}
	}

	return &Group{asMap}, nil
}

func allDifferentColors(tiles []Tile) bool {
	return len(GroupByColor(tiles)) == len(tiles)
}

func allSameNumber(tiles []Tile) bool {
	return len(GroupByNumber(tiles)) == 1
}
