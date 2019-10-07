package models

import "errors"

type Color string

const (
	ColorBlack  Color = "black"
	ColorRed    Color = "red"
	ColorBlue   Color = "blue"
	ColorOrange Color = "orange"
)

var (
	ErrSetTooShort = errors.New("set too short")
)

const SetMinLength = 3

type Tile struct {
	Color  Color
	Number int
}

func groupByColor(tiles []Tile) map[Color][]Tile {
	colors := make(map[Color][]Tile)
	for _, t := range tiles {
		colors[t.Color] = append(colors[t.Color], t)
	}
	return colors
}
