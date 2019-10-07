package models

import "errors"

type Run struct {
	tiles []Tile
}

var (
	ErrRunNotMonochrome = errors.New("run must be monochrome")
	ErrRunNonSequential = errors.New("run must be a sequence")
)

func (r *Run) Color() Color {
	return r.tiles[0].Color
}

func (r *Run) Length() int {
	return len(r.tiles)
}

func NewRun(tiles ...Tile) (*Run, error) {
	if len(tiles) < 3 {
		return nil, ErrSetTooShort
	}
	if !allSameColor(tiles) {
		return nil, ErrRunNotMonochrome
	}
	if !numbersAreSequence(tiles) {
		return nil, ErrRunNonSequential
	}
	return &Run{tiles}, nil
}

func allSameColor(tiles []Tile) bool {
	return len(groupByColor(tiles)) == 1
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
