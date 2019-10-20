package models

import "errors"

var (
	ErrGroupNotUnitary      = errors.New("group must be a single number")
	ErrGroupDuplicateColors = errors.New("group must not reuse a color")
)

type Group struct {
	tiles map[Tile]struct{}
}

func (g Group) Length() int {
	return len(g.tiles)
}

func (g Group) CommonNumber() int {
	for tile := range g.tiles {
		return tile.Number
	}
	panic("empty group")
}

func (g Group) Tiles() []Tile {
	tiles := make([]Tile, 0, len(g.tiles))
	for t := range g.tiles {
		tiles = append(tiles, t)
	}
	return tiles
}

func NewGroup(tiles ...Tile) (Group, error) {
	if len(tiles) < 3 {
		return Group{}, ErrSetTooShort
	}
	if !allDifferentColors(tiles) {
		return Group{}, ErrGroupDuplicateColors
	}
	if !allSameNumber(tiles) {
		return Group{}, ErrGroupNotUnitary
	}

	asMap := make(map[Tile]struct{})
	for _, t := range tiles {
		asMap[t] = struct{}{}
	}

	return Group{asMap}, nil
}

func allDifferentColors(tiles []Tile) bool {
	return len(CollectByColor(tiles)) == len(tiles)
}

func allSameNumber(tiles []Tile) bool {
	return len(CollectByNumber(tiles)) == 1
}
