package models

import "errors"

type Run struct {
	tiles []Tile
}

var (
	ErrRunNotMonochrome = errors.New("run must be monochrome")
	ErrRunNonSequential = errors.New("run must be a sequence")
)

func (r Run) CommonColor() Color {
	return r.tiles[0].Color
}

func (r Run) Length() int {
	return len(r.tiles)
}

func (r Run) Tiles() []Tile {
	tiles := make([]Tile, len(r.tiles))
	copy(tiles, r.tiles)
	return tiles
}

func NewRun(tiles ...Tile) (Run, error) {
	if len(tiles) < 3 {
		return Run{}, ErrSetTooShort
	}
	if !allSameColor(tiles) {
		return Run{}, ErrRunNotMonochrome
	}
	if !numbersAreSequence(tiles) {
		return Run{}, ErrRunNonSequential
	}
	return Run{tiles}, nil
}

func allSameColor(tiles []Tile) bool {
	return len(CollectByColor(tiles)) == 1
}

func numbersAreSequence(tiles []Tile) bool {
	first := tiles[0].Number
	for i := 0; i < len(tiles); i++ {
		if tiles[i].Number != first+i {
			return false
		}
	}
	return true
}
