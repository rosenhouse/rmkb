package models

import (
	"errors"
	"math/rand"
)

var ErrPlayers = errors.New("invalid list of players")

const MinNumPlayers = 2

type Board struct {
	Groups []Group
	Runs   []Run
}

func (b Board) TilesInPlay() []Tile {
	var n int
	for _, g := range b.Groups {
		n += g.Length()
	}
	for _, r := range b.Runs {
		n += r.Length()
	}
	allTiles := make([]Tile, 0, n)
	for _, g := range b.Groups {
		allTiles = append(allTiles, g.Tiles()...)
	}
	for _, r := range b.Runs {
		allTiles = append(allTiles, r.Tiles()...)
	}
	return allTiles
}

type Player struct {
	Name string
	Rack []Tile
}

type GameState struct {
	Board Board

	Pool []Tile

	Players []Player

	NextTurn int
}

func NewGame(seed int64, playerNames ...string) (*GameState, error) {
	if len(playerNames) < MinNumPlayers {
		return nil, ErrPlayers
	}
	gs := &GameState{}
	for _, playerName := range playerNames {
		if playerName == "" {
			return nil, ErrPlayers
		}
		gs.Players = append(gs.Players, Player{Name: playerName})
	}
	tiles := BuildTiles()
	shuffle(seed, tiles)
	gs.Pool = tiles
	return gs, nil
}

func shuffle(seed int64, tiles []Tile) {
	rand.New(rand.NewSource(seed)).Shuffle(len(tiles), func(i, j int) { tiles[i], tiles[j] = tiles[j], tiles[i] })
}

func (gs *GameState) Done() bool {
	return false
}

func BuildTiles() []Tile {
	allTiles := []Tile{}
	for _, color := range orderedColors {
		for n := MinTileNumber; n <= MaxTileNumber; n++ {
			allTiles = append(allTiles, Tile{Color: color, Number: n})
			allTiles = append(allTiles, Tile{Color: color, Number: n}) // two copies of each tile
		}
	}
	return allTiles
}
