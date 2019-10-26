package models

import (
	"errors"
	"sort"
)

type Color string

const (
	Black  Color = "black"
	Blue   Color = "blue"
	Orange Color = "orange"
	Red    Color = "red"

	NumColors     int = 4
	SetMinLength  int = 3
	MinTileNumber int = 1
	MaxTileNumber int = 13
)

var (
	ErrSetTooShort = errors.New("set too short")

	orderedColors = []Color{Black, Blue, Orange, Red}
	ColorIndex    = map[Color]int{
		Black:  0,
		Blue:   1,
		Orange: 2,
		Red:    3,
	}
)

type Tile struct {
	Color  Color
	Number int
}

type Set interface {
	Tiles() []Tile
}

func Compare(x, y *Tile) int {
	if x.Number < y.Number {
		return 1
	}
	if x.Number > y.Number {
		return -1
	}
	if ColorIndex[x.Color] < ColorIndex[y.Color] {
		return 1
	}
	if ColorIndex[x.Color] > ColorIndex[y.Color] {
		return -1
	}
	return 0
}

func SortTiles(toSort []Tile) {
	sort.Slice(toSort, func(i, j int) bool {
		return Compare(&toSort[i], &toSort[j]) < 0
	})
}

func CollectByColor(tiles []Tile) map[Color][]Tile {
	colors := make(map[Color][]Tile)
	for _, t := range tiles {
		colors[t.Color] = append(colors[t.Color], t)
	}
	return colors
}

func CollectByNumber(tiles []Tile) map[int][]Tile {
	numbers := make(map[int][]Tile)
	for _, t := range tiles {
		numbers[t.Number] = append(numbers[t.Number], t)
	}
	return numbers
}
