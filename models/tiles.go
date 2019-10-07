package models

import "errors"

type Color string

const (
	ColorBlack  Color = "black"
	ColorBlue   Color = "blue"
	ColorOrange Color = "orange"
	ColorRed    Color = "red"

	NumColors     int = 4
	SetMinLength  int = 3
	MinTileNumber int = 1
	MaxTileNumber int = 13
)

var (
	ErrSetTooShort = errors.New("set too short")

	orderedColors = []Color{ColorBlack, ColorBlue, ColorOrange, ColorRed}
)

type Tile struct {
	Color  Color
	Number int
}

func GroupByColor(tiles []Tile) map[Color][]Tile {
	colors := make(map[Color][]Tile)
	for _, t := range tiles {
		colors[t.Color] = append(colors[t.Color], t)
	}
	return colors
}

func GroupByNumber(tiles []Tile) map[int][]Tile {
	numbers := make(map[int][]Tile)
	for _, t := range tiles {
		numbers[t.Number] = append(numbers[t.Number], t)
	}
	return numbers
}
